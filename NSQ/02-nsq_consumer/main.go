package main

import (
	"fmt"
	"log"

	"github.com/nsqio/go-nsq"
)

type myMessageHandler struct{}

func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}

	fmt.Println("Message:", string(m.Body))
	return nil
}

func main() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("topic", "channel", config)
	if err != nil {
		log.Fatal(err)
	}
	consumer.AddHandler(&myMessageHandler{})
	err = consumer.ConnectToNSQLookupd("localhost:4161")
	if err != nil {
		log.Fatal(err)
	}

}
