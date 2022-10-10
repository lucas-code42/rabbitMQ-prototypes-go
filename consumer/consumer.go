package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "fail to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Fail to open a channel")
	defer ch.Close()

	msg, err := ch.Consume(
		"test",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "fail to consume msg")

	var test chan struct{}

	go func() {
		for d := range msg {
			log.Printf("msg received: %s", d.Body)
		}
	}()

	log.Printf(" [*] waiting for msgs")
	<-test

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", err, msg)
	}
}
