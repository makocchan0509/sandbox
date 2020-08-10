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
	subscribeMessage()
}

func subscribeMessage() {

	conn, err := amqp.Dial("amqp://admin:password@localhost:5672/")
	failOnError(err, "Failed connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a Channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		"dogOrCatQ",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("Wating for message. To exit press CTRL + C")
	<-forever
}
