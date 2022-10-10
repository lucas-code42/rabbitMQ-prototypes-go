package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "fail to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "fail to open a channel")
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "hello world!"

	err = ch.PublishWithContext(
		ctx,
		"",
		"test",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "fail to publish msg")
	log.Printf("[*] sent %s\n", body)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s --> %s", err, msg)
	}
}
