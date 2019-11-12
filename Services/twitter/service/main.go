package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"projects/Services/common/data"
	"projects/Services/common/net"
	"strconv"
	"time"

	"github.com/comail/colog"
	"github.com/gin-gonic/gin"
	"github.com/leekchan/timeutil"
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

	router.GET("/tweetpremium", tweetPreSearchService)
	router.Run(":8085")

}

//OPTIONSメソッドの処理
func optionsReceive(ctx *gin.Context) {

	log.Println("info: Called OPTONS method.")
	net.CORSForGin(ctx)
	ctx.Status(200)
}

//Twitterサービス
func tweetPreSearchService(ctx *gin.Context) {

	log.Println("info: Called tweetPreSearchService.")

	var tweetServiceReq data.TweetServiceReq

	tweetServiceReq.Query = ctx.Query("query")
	tweetServiceReq.OrderCount, _ = strconv.Atoi(ctx.Query("ordercount"))
	tweetServiceReq.MaxResults = ctx.Query("maxresults")
	tweetServiceReq.SearchMode = ctx.Query("searchmode")

	ctx.BindJSON(&tweetServiceReq)

	input, err := json.Marshal(tweetServiceReq)

	if err != nil {
		log.Println("error: ", err.Error())
		ctx.Status(404)
	}

	url := "http://localhost:8095/premiumsearch"

	log.Println("info: url:", url)
	log.Println("info: request parameter:", tweetServiceReq)

	//Request http POST
	res, err := net.JsonPostRequestSender(url, input)

	if err != nil {
		log.Println("error: ", err.Error())
		ctx.Status(404)
	}

	var tweetList []data.TweetList

	if err := json.Unmarshal(res, &tweetList); err != nil {
		log.Println("error: ", err.Error())
	}

	//log.Println("info: response parameter:", tweetServiceRes)

	var tweetServiceRes data.TweetServiceRes

	if len(tweetList) != 0 {
		tweetServiceRes.Result = "0"
		tweetServiceRes.GetCount = len(tweetList)
	} else {
		tweetServiceRes.Result = "1"
		tweetServiceRes.GetCount = 0
		ctx.JSON(http.StatusOK, tweetServiceRes)
	}

	//get time.
	t := time.Now()

	//open file.
	tweetFile, err := os.OpenFile("/Users/makotomase/var/tmp/tweetList_"+timeutil.Strftime(&t, "%Y%m%d%H%M%S")+".json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//エラー処理
		log.Println(err)
	}

	//sentimentFile, err := os.OpenFile("/Users/makotomase/var/tmp/sentimentList_"+timeutil.Strftime(&t, "%Y%m%d%H%M%S")+".json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	// if err != nil {
	// 	//エラー処理
	// 	log.Println(err)
	// }

	defer tweetFile.Close()
	// defer sentimentFile.Close()

	for _, tweet := range tweetList {
		b, err := json.Marshal(tweet)
		if err != nil {
			//エラー処理
			log.Println(err)
			break
		}
		fmt.Fprintln(tweetFile, string(b))
	}
	ctx.JSON(http.StatusOK, tweetServiceRes)

}
