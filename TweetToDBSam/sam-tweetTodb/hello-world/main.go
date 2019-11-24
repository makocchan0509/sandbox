package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// colog.SetDefaultLevel(colog.LDebug)
	// colog.SetMinLevel(colog.LTrace)
	// colog.SetFormatter(&colog.StdFormatter{
	// 	Colors: true,
	// 	Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	// })
	// colog.Register()

	fmt.Println("info: called handler")

	sess, err := session.NewSession()
	if err != nil {
		log.Println("error: ", err)
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintln(err),
			StatusCode: 200,
		}, err
	}

	fmt.Println("info: connect to dynamoDB")

	svc := dynamodb.New(sess)
	putParams := &dynamodb.PutItemInput{
		TableName: aws.String("user"),
		Item: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String("00001"),
			},
			"userName": {
				S: aws.String("testuser"),
			},
		},
	}

	fmt.Println("info: update dynamoDB")

	putItem, putErr := svc.PutItem(putParams)
	if putErr != nil {
		log.Println("error: ", putErr)
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintln(putErr),
			StatusCode: 200,
		}, putErr
	}

	fmt.Println(putItem)

	fmt.Println("info: completed handler")

	return events.APIGatewayProxyResponse{
		Body:       "{result,OK}",
		StatusCode: 200,
	}, nil

	// resp, err := http.Get(DefaultHTTPGetAddress)
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{}, err
	// }

	// if resp.StatusCode != 200 {
	// 	return events.APIGatewayProxyResponse{}, ErrNon200Response
	// }

	// ip, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{}, err
	// }

	// if len(ip) == 0 {
	// 	return events.APIGatewayProxyResponse{}, ErrNoIP
	// }

	// return events.APIGatewayProxyResponse{
	// 	Body:       fmt.Sprintf("Hello, %v", string(ip)),
	// 	StatusCode: 200,
	// }, nil
}

func main() {
	lambda.Start(handler)
}
