package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"projects/Services/common/data"
	"projects/Services/common/net"

	"github.com/comail/colog"
	"github.com/dghubble/oauth1"
	"github.com/gin-gonic/gin"
)

func main() {

	//setting log
	router := gin.Default()

	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()

	router.OPTIONS("/:uri", optionsReceive)

	router.POST("/premiumsearch", premiumSearch)
	router.Run(":8095")

	// //search_tweet(httpClient)

	// var request data.PremiumSearchReq

	// request.Query = "#タピオカ"
	// request.MaxResults = "10"

	// var limitCount = 10

	// PremiumSearch(request, httpClient, "0", limitCount)

}

//OPTIONSメソッドの処理
func optionsReceive(ctx *gin.Context) {

	log.Println("info: Called OPTONS method.")
	net.CORSForGin(ctx)
	ctx.Status(200)
}

// /*
//  * Standard search API
//  */
// func search_tweet(client *http.Client) {

// 	values := url.Values{}
// 	values.Add("q", "#渋谷")
// 	values.Add("count", "20")

// 	var result data.StandardSearchRes

// 	request, err := http.NewRequest("GET", "https://api.twitter.com/1.1/search/tweets.json"+"?"+values.Encode(), nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		//return result, err
// 	}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		fmt.Println(err)
// 		//return result, err
// 	}

// 	b, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		//return result, err
// 	}

// 	json.Unmarshal(b, &result)
// 	//fmt.Println(string(b))
// 	response.Body.Close()

// 	statuses := result.Statuses

// 	for _, tweet := range statuses { //デフォルトでは15個のTweetを拾ってくる
// 		fmt.Println("-------tweet start-------")
// 		fmt.Println("text:" + tweet.Text) //テキストの表示
// 		fmt.Println("-------tweet end-------")
// 	}
// }

func premiumSearch(ctx *gin.Context) {

	var req data.TweetServiceReq

	ctx.BindJSON(&req)

	log.Println("info: Received parameter", req)

	var apiQuery data.PremiumSearchReq

	apiQuery.Query = req.Query
	apiQuery.MaxResults = req.MaxResults

	//OAuth
	var consumerKey string
	var consumerSecret string
	var accessToken string
	var accessTokenSecret string

	var apiUrl string

	consumerKey = "T0z98JPbNM2dUZcpjJ49ospqm"
	consumerSecret = "GRWGxH72dbtSwqPOA2agLD9EFstHZszIS0dwJrB87DiOqsIsZU"
	accessToken = "2441537754-WfJieUTiRsSb098u6RY2EMC5laxB6bzGyVI48Nq"
	accessTokenSecret = "GXrKXibeCs0XHh4HMq5u3x4sxbqprxVSxx0vbOZGKtmAT"

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	//tweetを溜め込むリスト
	var resultTweets []data.TweetList

	var nextExists = true
	var encoded []byte
	var getCount = 0

	//2回目以降のJsonリクエストフォーマット
	var requestNext data.PremiumSearchNextReq
	requestNext.Query = req.Query
	requestNext.MaxResults = req.MaxResults

	//取得データ上限数
	orderCount := req.OrderCount

	//SearchModeからURL判定
	if req.SearchMode == "0" {
		apiUrl = "https://api.twitter.com/1.1/tweets/search/fullarchive/searchTweetsFull.json"
	} else {
		apiUrl = "https://api.twitter.com/1.1/tweets/search/30day/searchTweets30.json"
	}

	//Twiter APIへのリクエストパース
	encoded, err := json.Marshal(apiQuery)

	if err != nil {
		fmt.Println(err)
		nextExists = false
	}

	for nextExists {

		//POSTリクエスト生成
		req, err := http.NewRequest(
			"POST",
			apiUrl,
			bytes.NewBuffer(encoded),
		)
		if err != nil {
			fmt.Println(err)
			nextExists = false
		}
		req.Header.Set("Content-Type", "application/json")

		//送信＆レスポンス取得
		response, err := httpClient.Do(req)
		if err != nil {
			log.Println("error: err:", err)
			nextExists = false
		}
		//レスポンス解析
		b, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println("error: err:", err)
			nextExists = false
		}

		//レスポンス解析
		var result data.PremiumSearchRes
		json.Unmarshal(b, &result)
		response.Body.Close()

		for _, tweet := range result.Results {
			resultTweets = append(resultTweets, tweet)
		}

		//APIからNextが存在するかのチェック
		var next = ""
		next = result.Next
		if next != "" {
			//Nextをリクエストに設定
			requestNext.Next = next
			encodedNew, err := json.Marshal(requestNext)

			if err != nil {
				log.Println("error: err:", err)
				nextExists = false
			}
			//次のリクエストデータの上書き
			encoded = encodedNew
		} else {
			nextExists = false
		}

		//取得上限数超過チェック
		if orderCount != 0 {
			getCount = getCount + result.RequestParameters.MaxResults
			if getCount >= orderCount {
				nextExists = false
			}
		}

	}
	// for _, tweet := range resultTweets {
	// 	log.Println("info: tweet start")
	// 	log.Println("info: ", tweet.Text)
	// }

	ctx.JSON(http.StatusOK, resultTweets)

}
