package service

import (
	"log"
	"projects/Services/amazon/sqs/data"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

/**
 * send msessage service function.
 */
func SendMessageService(req data.SendMessageReq) (res data.SendMessageRes) {

	log.Println("info: called SendMessageService()")

	messageID, err := sendMessage(req.QueueURL, req.MessageBody, req.DelaySeconds)

	if err != nil {
		log.Println("error: send Message error", err)
		res.Result = "NG"
		return res
	}

	res.Result = "OK"
	res.QueueURL = req.QueueURL
	res.MessageBody = req.MessageBody
	res.MessageId = messageID

	return res
}

/**
 * send msessage common function.
 */
func sendMessage(queueURL string, message string, delaySecond int64) (messageID string, err error) {

	log.Println("info: called sendMessage()")

	region := "ap-northeast-1"

	session, err := session.NewSession(&aws.Config{Region: aws.String(region)})

	if err != nil {
		log.Println("error: create aws session", err)
		return "", err
	}
	svc := sqs.New(session)

	params := &sqs.SendMessageInput{
		MessageBody:  aws.String(message),
		QueueUrl:     aws.String(queueURL),
		DelaySeconds: aws.Int64(delaySecond),
	}

	sqsRes, err := svc.SendMessage(params)
	if err != nil {
		return "", err
	}

	messageID = *sqsRes.MessageId

	return messageID, nil
}
