package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	// 客户端连接消息队列
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 获取通道，所有操作基本都是通道控制的
	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	// 队列声明
	q, err := ch.QueueDeclare(
		"name",
		false,
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
		//发送消息
		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("Jung Ken"), //消息的内容
			},
		)
		if err != nil {
			log.Fatal(err)
			break loop
		}
		time.Sleep(time.Second)
		log.Println("producer is sending messages")
		var i int
		i++
		if i == 100 {
			break loop
		}
	}
}
