package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dghubble/oauth1"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("info: start handler")

	var req getTweetReq

	reqBody := request.Body
	jsonBytes := ([]byte)(reqBody)

	json.Unmarshal(jsonBytes, &req)

	resultTweets := getTweet(req)

	var result getTweetRes

	if len(resultTweets) == 0 {

		result.Result = "OK"
		result.TweetCount = len(resultTweets)
		result.PutCount = 0

		b, err := json.Marshal(result)

		if err != nil {
			fmt.Println(err)
			return events.APIGatewayProxyResponse{
				Body:       "",
				StatusCode: 500,
			}, nil
		}
		return events.APIGatewayProxyResponse{
			Body:       string(b),
			StatusCode: 200,
		}, nil
	}

	count, err := putDynamo(resultTweets, req)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 500,
		}, nil
	}

	result.Result = "OK"
	result.TweetCount = len(resultTweets)
	result.PutCount = count

	b, err := json.Marshal(result)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 500,
		}, nil
	}

	fmt.Println("info: completed handler")

	return events.APIGatewayProxyResponse{
		Body:       string(b),
		StatusCode: 200,
	}, nil
}

func getTweet(req getTweetReq) (resultTweets []tweetList) {

	fmt.Println("info: start getTweet")

	var apiQuery premiumSearchReq

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
	//var resultTweets []tweetList

	var nextExists = true
	var encoded []byte
	var getCount = 0

	//2回目以降のJsonリクエストフォーマット
	var requestNext premiumSearchNextReq
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
			fmt.Println("error: err:", err)
			nextExists = false
		}
		req.Header.Set("Content-Type", "application/json")

		//送信＆レスポンス取得
		response, err := httpClient.Do(req)
		if err != nil {
			fmt.Println("error: err:", err)
			nextExists = false
		}
		//レスポンス解析
		b, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("error: err:", err)
			nextExists = false
		}

		//レスポンス解析
		var result premiumSearchRes
		json.Unmarshal(b, &result)
		response.Body.Close()

		for _, tweet := range result.Results {
			resultTweets = append(resultTweets, tweet)
		}

		//APIからNextが存在するかのチェック
		var next = ""
		next = result.Next
		if next != "" {
			//fmt.Print("Nextあり--------------")
			//Nextをリクエストに設定
			requestNext.Next = next
			encodedNew, err := json.Marshal(requestNext)

			if err != nil {
				fmt.Println("error: err:", err)
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
	fmt.Println("info: get tweet count", getCount)
	fmt.Println("info: completed getTweet")

	return resultTweets

}

func putDynamo(resultTweets []tweetList, req getTweetReq) (count int, err error) {

	fmt.Println("info: start putDynamo")
	count = 0
	sess, err := session.NewSession()

	if err != nil {
		fmt.Println("error: ", err)
		return count, err
	}
	svc := dynamodb.New(sess)

	var putData dynamoData

	for _, tweet := range resultTweets {

		putData.ID = tweet.ID
		putData.Query = req.Query
		putData.Tweet = tweet

		item, err := dynamodbattribute.MarshalMap(putData)

		if err != nil {
			fmt.Println("error: ", err)
			return count, err
		}

		putParams := &dynamodb.PutItemInput{
			TableName: aws.String("tweet"),
			Item:      item,
		}

		_, putErr := svc.PutItem(putParams)
		if putErr != nil {
			fmt.Println("error: ", putErr)
			return count, putErr
		}

		count = count + 1
	}

	fmt.Println("info: put item count", count)
	fmt.Println("info: completed putDynamo")

	return count, nil
}

func main() {
	lambda.Start(handler)
}

type getTweetReq struct {
	Query      string `json:"query"`
	OrderCount int    `json:"ordercount"`
	MaxResults string `json:"maxResults"`
	SearchMode string `json:"searchmode"`
}

type getTweetRes struct {
	Result     string `json:"result"`
	TweetCount int    `json:"gettweetcount"`
	PutCount   int    `json:"putcount"`
}

type premiumSearchReq struct {
	Query      string `json:"query"`
	MaxResults string `json:"maxResults"`
}

type premiumSearchNextReq struct {
	Query      string `json:"query"`
	MaxResults string `json:"maxResults"`
	Next       string `json:"next"`
}

type premiumSearchRes struct {
	Results           []tweetList `json:"results"`
	Next              string      `json:"next"`
	RequestParameters struct {
		MaxResults int    `json:"maxResults"`
		FromDate   string `json:"fromDate"`
		ToDate     string `json:"toDate"`
	} `json:"requestParameters"`
}

type dynamoData struct {
	ID    int64     `json:"id" dynamo:"id"`
	Query string    `json:"query" dynamo:"query"`
	Tweet tweetList `json:"tweet" dynamo:"tweet"`
}

type tweetList struct {
	CreatedAt            string      `json:"created_at"`
	ID                   int64       `json:"id"`
	IDStr                string      `json:"id_str"`
	Text                 string      `json:"text"`
	Source               string      `json:"source"`
	Truncated            bool        `json:"truncated"`
	InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
	InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
	InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
	User                 struct {
		ID                             int64       `json:"id"`
		IDStr                          string      `json:"id_str"`
		Name                           string      `json:"name"`
		ScreenName                     string      `json:"screen_name"`
		Location                       string      `json:"location"`
		URL                            interface{} `json:"url"`
		Description                    interface{} `json:"description"`
		TranslatorType                 string      `json:"translator_type"`
		Protected                      bool        `json:"protected"`
		Verified                       bool        `json:"verified"`
		FollowersCount                 int         `json:"followers_count"`
		FriendsCount                   int         `json:"friends_count"`
		ListedCount                    int         `json:"listed_count"`
		FavouritesCount                int         `json:"favourites_count"`
		StatusesCount                  int         `json:"statuses_count"`
		CreatedAt                      string      `json:"created_at"`
		UtcOffset                      interface{} `json:"utc_offset"`
		TimeZone                       interface{} `json:"time_zone"`
		GeoEnabled                     bool        `json:"geo_enabled"`
		Lang                           interface{} `json:"lang"`
		ContributorsEnabled            bool        `json:"contributors_enabled"`
		IsTranslator                   bool        `json:"is_translator"`
		ProfileBackgroundColor         string      `json:"profile_background_color"`
		ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
		ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
		ProfileBackgroundTile          bool        `json:"profile_background_tile"`
		ProfileLinkColor               string      `json:"profile_link_color"`
		ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
		ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
		ProfileTextColor               string      `json:"profile_text_color"`
		ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
		ProfileImageURL                string      `json:"profile_image_url"`
		ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
		ProfileBannerURL               string      `json:"profile_banner_url"`
		DefaultProfile                 bool        `json:"default_profile"`
		DefaultProfileImage            bool        `json:"default_profile_image"`
		Following                      interface{} `json:"following"`
		FollowRequestSent              interface{} `json:"follow_request_sent"`
		Notifications                  interface{} `json:"notifications"`
	} `json:"user"`
	Geo             interface{} `json:"geo"`
	Coordinates     interface{} `json:"coordinates"`
	Place           interface{} `json:"place"`
	Contributors    interface{} `json:"contributors"`
	RetweetedStatus struct {
		CreatedAt            string      `json:"created_at"`
		ID                   int64       `json:"id"`
		IDStr                string      `json:"id_str"`
		Text                 string      `json:"text"`
		DisplayTextRange     []int       `json:"display_text_range"`
		Source               string      `json:"source"`
		Truncated            bool        `json:"truncated"`
		InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
		InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
		InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
		InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
		InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
		User                 struct {
			ID                             int64       `json:"id"`
			IDStr                          string      `json:"id_str"`
			Name                           string      `json:"name"`
			ScreenName                     string      `json:"screen_name"`
			Location                       interface{} `json:"location"`
			URL                            string      `json:"url"`
			Description                    string      `json:"description"`
			TranslatorType                 string      `json:"translator_type"`
			Protected                      bool        `json:"protected"`
			Verified                       bool        `json:"verified"`
			FollowersCount                 int         `json:"followers_count"`
			FriendsCount                   int         `json:"friends_count"`
			ListedCount                    int         `json:"listed_count"`
			FavouritesCount                int         `json:"favourites_count"`
			StatusesCount                  int         `json:"statuses_count"`
			CreatedAt                      string      `json:"created_at"`
			UtcOffset                      interface{} `json:"utc_offset"`
			TimeZone                       interface{} `json:"time_zone"`
			GeoEnabled                     bool        `json:"geo_enabled"`
			Lang                           interface{} `json:"lang"`
			ContributorsEnabled            bool        `json:"contributors_enabled"`
			IsTranslator                   bool        `json:"is_translator"`
			ProfileBackgroundColor         string      `json:"profile_background_color"`
			ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
			ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
			ProfileBackgroundTile          bool        `json:"profile_background_tile"`
			ProfileLinkColor               string      `json:"profile_link_color"`
			ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
			ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
			ProfileTextColor               string      `json:"profile_text_color"`
			ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
			ProfileImageURL                string      `json:"profile_image_url"`
			ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
			ProfileBannerURL               string      `json:"profile_banner_url"`
			DefaultProfile                 bool        `json:"default_profile"`
			DefaultProfileImage            bool        `json:"default_profile_image"`
			Following                      interface{} `json:"following"`
			FollowRequestSent              interface{} `json:"follow_request_sent"`
			Notifications                  interface{} `json:"notifications"`
		} `json:"user"`
		Geo           interface{} `json:"geo"`
		Coordinates   interface{} `json:"coordinates"`
		Place         interface{} `json:"place"`
		Contributors  interface{} `json:"contributors"`
		IsQuoteStatus bool        `json:"is_quote_status"`
		QuoteCount    int         `json:"quote_count"`
		ReplyCount    int         `json:"reply_count"`
		RetweetCount  int         `json:"retweet_count"`
		FavoriteCount int         `json:"favorite_count"`
		Entities      struct {
			Hashtags []struct {
				Text    string `json:"text"`
				Indices []int  `json:"indices"`
			} `json:"hashtags"`
			Urls         []interface{} `json:"urls"`
			UserMentions []interface{} `json:"user_mentions"`
			Symbols      []interface{} `json:"symbols"`
			Media        []struct {
				ID                  int64  `json:"id"`
				IDStr               string `json:"id_str"`
				Indices             []int  `json:"indices"`
				AdditionalMediaInfo struct {
					Monetizable bool `json:"monetizable"`
				} `json:"additional_media_info"`
				MediaURL      string `json:"media_url"`
				MediaURLHTTPS string `json:"media_url_https"`
				URL           string `json:"url"`
				DisplayURL    string `json:"display_url"`
				ExpandedURL   string `json:"expanded_url"`
				Type          string `json:"type"`
				Sizes         struct {
					Thumb struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"thumb"`
					Medium struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"medium"`
					Small struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"small"`
					Large struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"large"`
				} `json:"sizes"`
			} `json:"media"`
		} `json:"entities"`
		ExtendedEntities struct {
			Media []struct {
				ID                  int64  `json:"id"`
				IDStr               string `json:"id_str"`
				Indices             []int  `json:"indices"`
				AdditionalMediaInfo struct {
					Monetizable bool `json:"monetizable"`
				} `json:"additional_media_info"`
				MediaURL      string `json:"media_url"`
				MediaURLHTTPS string `json:"media_url_https"`
				URL           string `json:"url"`
				DisplayURL    string `json:"display_url"`
				ExpandedURL   string `json:"expanded_url"`
				Type          string `json:"type"`
				VideoInfo     struct {
					AspectRatio    []int `json:"aspect_ratio"`
					DurationMillis int   `json:"duration_millis"`
					Variants       []struct {
						ContentType string `json:"content_type"`
						URL         string `json:"url"`
						Bitrate     int    `json:"bitrate,omitempty"`
					} `json:"variants"`
				} `json:"video_info"`
				Sizes struct {
					Thumb struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"thumb"`
					Medium struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"medium"`
					Small struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"small"`
					Large struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"large"`
				} `json:"sizes"`
			} `json:"media"`
		} `json:"extended_entities"`
		Favorited         bool   `json:"favorited"`
		Retweeted         bool   `json:"retweeted"`
		PossiblySensitive bool   `json:"possibly_sensitive"`
		FilterLevel       string `json:"filter_level"`
		Lang              string `json:"lang"`
	} `json:"retweeted_status"`
	IsQuoteStatus bool `json:"is_quote_status"`
	QuoteCount    int  `json:"quote_count"`
	ReplyCount    int  `json:"reply_count"`
	RetweetCount  int  `json:"retweet_count"`
	FavoriteCount int  `json:"favorite_count"`
	Entities      struct {
		Hashtags []struct {
			Text    string `json:"text"`
			Indices []int  `json:"indices"`
		} `json:"hashtags"`
		Urls         []interface{} `json:"urls"`
		UserMentions []struct {
			ScreenName string `json:"screen_name"`
			Name       string `json:"name"`
			ID         int64  `json:"id"`
			IDStr      string `json:"id_str"`
			Indices    []int  `json:"indices"`
		} `json:"user_mentions"`
		Symbols []interface{} `json:"symbols"`
		Media   []struct {
			ID                  int64  `json:"id"`
			IDStr               string `json:"id_str"`
			Indices             []int  `json:"indices"`
			AdditionalMediaInfo struct {
				Monetizable bool `json:"monetizable"`
			} `json:"additional_media_info"`
			MediaURL      string `json:"media_url"`
			MediaURLHTTPS string `json:"media_url_https"`
			URL           string `json:"url"`
			DisplayURL    string `json:"display_url"`
			ExpandedURL   string `json:"expanded_url"`
			Type          string `json:"type"`
			Sizes         struct {
				Thumb struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"thumb"`
				Medium struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"medium"`
				Small struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"small"`
				Large struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"large"`
			} `json:"sizes"`
			SourceStatusID    int64  `json:"source_status_id"`
			SourceStatusIDStr string `json:"source_status_id_str"`
			SourceUserID      int64  `json:"source_user_id"`
			SourceUserIDStr   string `json:"source_user_id_str"`
		} `json:"media"`
	} `json:"entities"`
	ExtendedEntities struct {
		Media []struct {
			ID                  int64  `json:"id"`
			IDStr               string `json:"id_str"`
			Indices             []int  `json:"indices"`
			AdditionalMediaInfo struct {
				Monetizable bool `json:"monetizable"`
			} `json:"additional_media_info"`
			MediaURL      string `json:"media_url"`
			MediaURLHTTPS string `json:"media_url_https"`
			URL           string `json:"url"`
			DisplayURL    string `json:"display_url"`
			ExpandedURL   string `json:"expanded_url"`
			Type          string `json:"type"`
			VideoInfo     struct {
				AspectRatio    []int `json:"aspect_ratio"`
				DurationMillis int   `json:"duration_millis"`
				Variants       []struct {
					ContentType string `json:"content_type"`
					URL         string `json:"url"`
					Bitrate     int    `json:"bitrate,omitempty"`
				} `json:"variants"`
			} `json:"video_info"`
			Sizes struct {
				Thumb struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"thumb"`
				Medium struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"medium"`
				Small struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"small"`
				Large struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"large"`
			} `json:"sizes"`
			SourceStatusID    int64  `json:"source_status_id"`
			SourceStatusIDStr string `json:"source_status_id_str"`
			SourceUserID      int64  `json:"source_user_id"`
			SourceUserIDStr   string `json:"source_user_id_str"`
		} `json:"media"`
	} `json:"extended_entities"`
	Favorited         bool   `json:"favorited"`
	Retweeted         bool   `json:"retweeted"`
	PossiblySensitive bool   `json:"possibly_sensitive"`
	FilterLevel       string `json:"filter_level"`
	Lang              string `json:"lang"`
	MatchingRules     []struct {
		Tag interface{} `json:"tag"`
	} `json:"matching_rules"`
}
