package main

import (
	"bytes"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	checkErr(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	checkErr(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	checkErr(err, "Failed to declare queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	checkErr(err, "Failed to declare consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Receive a message:%s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			log.Println("dotCount:", dotCount)
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
