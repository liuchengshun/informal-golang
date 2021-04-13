package main

import (
	"log"

	"github.com/streadway/amqp"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}

func main() {
	// 创建一个连接
	// 改链接抽象了socket连接，并为我们处理协议版本协商和身份验证等
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	checkErr(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 创建一个管道，其工作是完成大多数API所在的位置。
	ch, err := conn.Channel()
	checkErr(err, "failed to open a channel")
	defer ch.Close()

	// 定义一个队列，我们之后会向其发送数据
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durbale
		false,   // delete when unused
		false,   //exclusive
		false,   //no-wait
		nil,     // arguments
	)
	checkErr(err, "Failed to declare a queue")

	body := "hello world"

	err = ch.Publish(
		"",     //exchange name (默认交换机)
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	checkErr(err, "Failed to publish message")
}
