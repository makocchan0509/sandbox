package rabbitMQ

import (
	"log"
	"projects/DogOrCat/web/config"

	"github.com/streadway/amqp"
)

var rabbitUrl = config.Env.RabbitUrl
var rabbitUser = config.Env.RabbitUser
var rabbitPass = config.Env.RabbitPass
var rabbitQName = config.Env.RabbitQName

func PublishMessage(body string) (err error) {
	conn, err := amqp.Dial("amqp://" + rabbitUser + ":" + rabbitPass + "@" + rabbitUrl)
	if err != nil {
		log.Fatalf("Failed connect to RabbitMQ: %s", err.Error())
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed open to Channel: %s", err.Error())
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rabbitQName,
		false,
		false,
		false,
		false,
		nil,
	)

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed publish message: %s", err.Error())
		return err
	}
	log.Printf("Success publish message: %s ,%s", q.Name, body)
	return err
}
