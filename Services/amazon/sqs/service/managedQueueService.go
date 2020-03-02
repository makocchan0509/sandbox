package service

import (
	"fmt"
	"log"
	"projects/Services/amazon/sqs/data"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	AWS_REGION = "ap-northeast-1"
)

var svc *sqs.SQS

/**
 * Management queue service function.
 */
func CreateQueueService(req data.CreateQueueReq) (res data.CreateQueueRes) {

	log.Println("info: called CreateQueueService()")

	session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_REGION)})

	if err != nil {
		log.Println("error: create aws session", err)
		res.Result = "NG"
		return res
	}

	svc = sqs.New(session)

	var queueArn string
	var deadQueueUrl string

	queueArn = ""

	if req.DeadLaterQueueName != "" {

		deadQueueUrl, err = createDeadLaterQueue(req.DeadLaterQueueName)

		if err != err {
			log.Println("error: create dead later queue error", err)
			res.Result = "NG"
			return res
		}
		res.DeadLaterQueueUrl = deadQueueUrl
		attributes, err := getQueueArn(deadQueueUrl)

		if err != err {
			log.Println("error: get queue arn error", err)
			res.Result = "NG"
			return res
		}
		queueArn = *attributes["QueueArn"]
	}

	mainQueueUrl, err := createMainQueue(req.QueueName, queueArn, req.MaxReceiveCount, req.VisibilityTimeout)

	if err != err {
		log.Println("error: create main queue error", err)
		res.Result = "NG"
		return res
	}

	res.Result = "OK"
	res.MainQueueUrl = mainQueueUrl

	return res
}

/**
 * Create dead later queue function.
 */
func createDeadLaterQueue(queueName string) (url string, err error) {

	log.Println("info: called createDeadLaterQueue()")

	params := &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
	}

	url, err = createQueue(params)

	if err != nil {
		return "", err
	}

	return url, err
}

/**
 * Create queue common function.
 */
func createQueue(params *sqs.CreateQueueInput) (url string, err error) {

	log.Println("info: called createQueue()")

	resp, err := svc.CreateQueue(params)

	if err != nil {
		return "", err
	}
	return *resp.QueueUrl, nil
}

/**
 * Get queue ARN function
 */

func getQueueArn(url string) (map[string]*string, error) {

	log.Println("info: called getQueueArn()")

	params := &sqs.GetQueueAttributesInput{
		QueueUrl: aws.String(url),
		AttributeNames: []*string{
			aws.String("All"),
		},
	}
	resp, err := svc.GetQueueAttributes(params)

	if err != nil {
		return nil, err
	}

	return resp.Attributes, nil
}

/**
 * Create main queue function.
 */
func createMainQueue(queuename string, deadQueueArn string, maxreceivecount string, visibilityTimeout string) (queueurl string, err error) {

	log.Println("info: called createMainQueue()")

	var redrivePolicy string

	redrivePolicy = ""

	if deadQueueArn != "" {

		redrivePolicy = fmt.Sprintf(
			"{\"deadLetterTargetArn\":\"%s\",\"maxReceiveCount\":%s}",
			deadQueueArn,
			maxreceivecount,
		)
	}

	// キュー名を指定してキュー作成。
	// 設定内容はAttributesに足していく。
	params := &sqs.CreateQueueInput{
		QueueName: aws.String(queuename),
		Attributes: map[string]*string{
			// VisibilityTimeout:取得したメッセージは指定した秒数の間、他から見えなくする。
			"VisibilityTimeout": aws.String(visibilityTimeout),

			// ReceiveMessageWaitTimeSeconds: ロングポーリングの秒数。
			// キューが空だった場合、どれだけ待つか。
			// ここで指定しなくても、メッセージ取得時に指定は可能。
			//"ReceiveMessageWaitTimeSeconds": aws.String("20"),

			// RedrivePolicy: デッドキュー用ポリシー。先に作っておいた値を設定。
			"RedrivePolicy": aws.String(redrivePolicy),
		},
	}

	queueurl, err = createQueue(params)

	if err != nil {
		return "", err
	}

	return queueurl, nil
}
