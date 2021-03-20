package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Client struct {
	Name    string    `validate:"required,contains=HuBei,excludes=DaXue"`
	Linkman string    `validate:"startswith=Mr.|startswith=Ms."`
	Address string    `validate:"endswith=(China)"`
}

func main() {
	client := &Client {
		Name: "HuBeiDaXue",
		Linkman: "Mr.Zhang",
		Address: "HuBeiChina",
	}
	fmt.Println("Start")

	// 创建一个实例
	validate :=  validator.New()

	// 检测结构体
	err := validate.Struct(client)
	if err != nil {
		fmt.Println(err)
		// error_1  Client.Name   failed on the 'excludes' tag
		// error_2  Client.Address   failed on the 'endswith' tag
	}

	fmt.Println("End")
}