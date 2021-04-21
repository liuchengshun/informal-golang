package main

import (
	"log"

	"github.com/streadway/amqp"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s, error: %v", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	checkErr(err, "failed to dial rabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	checkErr(err, "failed to create channel")

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	checkErr(err, "failed to declare queue")

	err = ch.QueueBind(
		q.Name,
		"",
		"logs",
		false,
		nil,
	)
	checkErr(err, "bind queue error")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	checkErr(err, "failed to consumer messages")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("%s", d.Body)
		}
	}()

	log.Println("Waiting for logs, To exit press CTRL+C")
	<-forever
}
