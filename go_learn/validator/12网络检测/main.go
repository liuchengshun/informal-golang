package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type contact struct {
	Phone  string
	Email  string
}

type NetWork struct { 
   IP   string    `validate:"ip"` 
   IPV4 string    `validate:"ipv4"`
   IPV6 string    `validate:"ipv6"`
   URL  string    `validate:"url"`
   URI  string    `validate:"uri"`
}

func main() {
	net := &NetWork {
		IP:   "127.0.0.0",
		IPV4: "127.255.255.254",
		IPV6: "2001:0D12:0000:0000:02AA:0987:FE29:9871",
		URL:  "https://github.com/go-playground/validator",
		URI:  "http://127.0.0.1:8080/cmd_helloworld/?name=guowuxin",
	}
	fmt.Println("Start")

	// 创建一个实例
	validate :=  validator.New()

	// 检测结构体
	err := validate.Struct(net)
	if err != nil {
		fmt.Println(err)
		// success
	}

	fmt.Println("End")
}