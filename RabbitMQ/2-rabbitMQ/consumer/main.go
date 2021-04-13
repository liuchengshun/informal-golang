package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 定义通道
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	// 定义交换机
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

	// 定义临时通道
	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	// 将交换机和临时通道关联
	err = ch.QueueBind(
		q.Name,
		"",
		"ex-change",
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	// 监听消息(在消费者看来是"消费"消息)
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 打印消息
	for v := range msgs {
		log.Printf("body = %s\n", v.Body)
	}
}
