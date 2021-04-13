package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// 建立连接
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 获取通道
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	// 声明队列
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

	// 声明每次从队列获取任务的数量
	// 1 是预取任务数量  0 领取大小  false 全局设置
	err = ch.Qos(1, 0, false)
	if err != nil {
		log.Fatal(err)
	}

	// 消费者接收消息
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 打印消息队列
	for v := range msgs {
		fmt.Printf("body = %v\n", string(v.Body))
	}
}
