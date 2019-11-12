// Sample language-quickstart uses the Google Cloud Natural API to analyze the
// sentiment of "Hello, world!".
package main

import (
	"log"
	"projects/Services/common/net"

	"github.com/comail/colog"
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

	router.POST("/premiumsearch", anlzeSentimentle)
	router.Run(":8105")
}

//OPTIONSメソッドの処理
func optionsReceive(ctx *gin.Context) {

	log.Println("info: Called OPTONS method.")
	net.CORSForGin(ctx)
	ctx.Status(200)
}

func anlzeSentimentle(ctx *gin.Context) {

}

// ctx := context.Background()

// // Creates a client.
// client, err := language.NewClient(ctx)
// if err != nil {
// 	log.Fatalf("Failed to create client: %v", err)
// }

// // Sets the text to analyze.
// text := "Hello, world!"

// // Detects the sentiment of the text.
// sentiment, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
// 	Document: &languagepb.Document{
// 		Source: &languagepb.Document_Content{
// 			Content: text,
// 		},
// 		Type: languagepb.Document_PLAIN_TEXT,
// 	},
// 	EncodingType: languagepb.EncodingType_UTF8,
// })
// if err != nil {
// 	log.Fatalf("Failed to analyze text: %v", err)
// }

// fmt.Printf("Text: %v\n", text)
// fmt.Println("Score", sentiment.DocumentSentiment.Score)
// if sentiment.DocumentSentiment.Score >= 0 {
// 	fmt.Println("Sentiment: positive")
// } else {
// 	fmt.Println("Sentiment: negative")
// }
// }
