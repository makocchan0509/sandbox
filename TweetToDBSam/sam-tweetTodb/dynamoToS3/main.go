package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/guregu/dynamo"
	"github.com/leekchan/timeutil"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("info: start handler()")

	var result uploadRes

	getItems, err := getItemDynamo()

	if err != nil {
		fmt.Println("error: error occured", err)

		result.Result = "NG"
		result.UploadCount = 0
		b, err := json.Marshal(result)

		return events.APIGatewayProxyResponse{
			Body:       string(b),
			StatusCode: 500,
		}, err
	}

	if len(getItems) == 0 {

		result.Result = "OK"
		result.UploadCount = 0
		b, _ := json.Marshal(result)
		return events.APIGatewayProxyResponse{
			Body:       string(b),
			StatusCode: 200,
		}, nil
	}

	err = uploadToS3(getItems)

	if err != nil {
		fmt.Println("error: error occured", err)

		result.Result = "NG"
		result.UploadCount = 0
		b, err := json.Marshal(result)

		return events.APIGatewayProxyResponse{
			Body:       string(b),
			StatusCode: 500,
		}, err
	}

	result.Result = "OK"
	result.UploadCount = len(getItems)
	b, _ := json.Marshal(result)

	fmt.Println("info: completed handler()")

	return events.APIGatewayProxyResponse{
		Body:       string(b),
		StatusCode: 200,
	}, nil
}

func uploadToS3(getItems []dynamoData) error {

	fmt.Println("info: start uploadToS3()")

	//get time.
	t := time.Now()

	var sess = session.Must(session.NewSession(&aws.Config{
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(os.Getenv("REGION")),
		Endpoint:         aws.String(os.Getenv("S3_ENDPOINT")),
	}))

	uploader := s3manager.NewUploader(sess)

	var keyName string
	keyName = "tweetListFromDynamo_" + timeutil.Strftime(&t, "%Y%m%d%H%M%S") + ".json"

	b, _ := json.Marshal(getItems)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(keyName),
		Body:   bytes.NewReader(b),
	})
	if err != nil {
		fmt.Println("error: error occured", err)
		return err
	}

	fmt.Println("info : upload output", result)
	fmt.Println("info: completed uploadToS3()")
	return nil

}

func getItemDynamo() ([]dynamoData, error) {

	fmt.Println("info: start getItemDynamo()")
	var getItems []dynamoData

	db := dynamo.New(session.New())
	table := db.Table("tweet")
	err := table.Scan().All(&getItems)

	if err != nil {

		fmt.Println("error: error occured", err)
		return nil, err
	}

	fmt.Println("info: get table count", len(getItems))

	fmt.Println("info: completed getItemDynamo()")

	return getItems, nil
}

func main() {
	lambda.Start(handler)
}

type uploadRes struct {
	Result      string `json:"result"`
	UploadCount int    `json:"uploadcount"`
}

type dynamoData struct {
	ID    int64       `json:"id" dynamo:"id"`
	Query string      `json:"query" dynamo:"query"`
	Tweet interface{} `json:"tweet" dynamo:"tweet"`
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
