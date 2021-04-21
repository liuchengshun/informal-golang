package main

import (
	"log"

	"github.com/streadway/amqp"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	checkErr(err)
	defer conn.Close()

	ch, err := conn.Channel()
	checkErr(err)

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	checkErr(err)

	err = ch.QueueBind(
		q.Name,
		"",
		"logs",
		false,
		nil,
	)
	checkErr(err)

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		true,
		false,
		false,
		nil,
	)
	checkErr(err)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("%s", d.Body)
		}
	}()

	<-forever
}
