package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest:admin@127.0.0.1:5672")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	// 定义一个交换机，使用交换机来代替管道发布消息
	err = ch.ExchangeDeclare(
		"ex-change",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

loop:
	for {
		// 发布消息
		err = ch.Publish(
			"ex-change",
			"",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("liuchengshun"),
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
		n := 0
		n++
		if n == 100 {
			break loop
		}
	}

}
