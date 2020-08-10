package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}

func main() {
	publishMessage()

}

func publishMessage() {
	conn, err := amqp.Dial("amqp://admin:password@localhost:5672/")
	failOnError(err, "Failed connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a Channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"dogOrCatQ",
		false,
		false,
		false,
		false,
		nil,
	)

	body := "HelloWorld"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a channel")

}
