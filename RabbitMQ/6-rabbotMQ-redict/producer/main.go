package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

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

	err = ch.ExchangeDeclare(
		"logs_direct",
		"direct",
		false,
		false,
		false,
		false,
		nil,
	)
	checkErr(err)

	for n := 0; n < 100; n++ {
		msg := fmt.Sprintf("%v", rand.Int63())

		err = ch.Publish(
			"logs_direct",
			"wuhu",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
			},
		)
		checkErr(err)

		err = ch.Publish(
			"logs_direct",
			"wuhu123",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
			},
		)
		checkErr(err)

		log.Printf("%v", msg)
		time.Sleep(time.Second)
	}
}
