package data

/**
* CreateQueueService Request Struct.
**/
type CreateQueueReq struct {
	QueueName          string `json:"queueName"`
	DeadLaterQueueName string `json:"deadLaterQueueName"`
	MaxReceiveCount    string `json:"maxReceiveCount"`
	VisibilityTimeout  string `json:"visibilityTimeout"`
}

/**
* CreateQueueService Response Struct.
**/
type CreateQueueRes struct {
	Result            string `json:"result"`
	MainQueueUrl      string `json:"mainQueueUrl"`
	DeadLaterQueueUrl string `json:"deadLateQueueUrl"`
}

/**
* sendMessageService Request Struct.
**/
type SendMessageReq struct {
	QueueURL     string `json:"queueURL"`
	MessageBody  string `json:"messageBody"`
	DelaySeconds int64  `json:"delaySeconds"`
}

/**
* sendMessageService Response Struct.
**/
type SendMessageRes struct {
	Result      string `json:"result"`
	MessageId   string `json:"messageId"`
	QueueURL    string `json:"queueURL"`
	MessageBody string `json:"messageBody"`
}
