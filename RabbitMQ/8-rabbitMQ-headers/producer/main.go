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
		"logs_headers",
		"headers",
		false,
		false,
		false,
		false,
		nil,
	)
	checkErr(err)

	for n := 0; n < 100; n++ {
		body := fmt.Sprintf("%v", rand.Int())

		err := ch.Publish(
			"logs_headers",
			"",
			false,
			false,
			amqp.Publishing{
				Headers: amqp.Table{"name": "ZhangSan"},
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
		checkErr(err)

		err = ch.Publish(
			"logs_headers",
			"",
			false,
			false,
			amqp.Publishing{
				Headers: amqp.Table{"address": "hubei"},
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
		checkErr(err)


		log.Printf("%s", body)
		time.Sleep(time.Second)
	}

}
