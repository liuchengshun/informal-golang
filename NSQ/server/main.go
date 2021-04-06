package main

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

type ConsumerT struct{}

func main() {
	InitConsumer("test", "test-channel", "127.0.0.1:4161")
	for {
		time.Sleep(time.Second * 10)
	}
}

// 处理消息
func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive:", msg.NSQDAddress, " message:", string(msg.Body))
	return nil
}

func InitConsumer(topic string, channel string, address string) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second          // 设置重连时间
	c, err := nsq.NewConsumer(topic, channel, cfg) // 新建一个消费者
	if err != nil {
		panic(err)
	}
	c.SetLogger(nil, 0)
	c.AddHandler(&ConsumerT{})

	// 建立SQLookupd连接
	if err := c.ConnectToNSQLookupd(address); err != nil {
		panic(err)
	}
}
