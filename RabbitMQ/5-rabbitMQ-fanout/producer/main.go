package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s, error: %v", msg, err)
	}
}

// func bodyFrom(args ...[]byte) string {
// 	if len(args) < 2 || os.Args[1] == "" {
// 		return hello,
// 	}
// }

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	checkErr(err, "failed to dial rabbitmq")
	defer conn.Close()

	ch, err := conn.Channel()
	checkErr(err, "failed to create channel")

	err = ch.ExchangeDeclare(
		"logs",
		"fanout",
		false,
		false,
		false,
		false,
		nil,
	)
	checkErr(err, "failed to declare exchange")

	for i := 0; i < 100; i ++ {
		body := fmt.Sprintf("%v", rand.Uint32())

		err = ch.Publish(
			"logs",
			"",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
		checkErr(err, "failed to publish messages")

		log.Println(body)
		time.Sleep(time.Second)  
	}
}
