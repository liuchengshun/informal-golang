package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name    string    `validate:"required,max=10"`
	Age     int       `validate:"required,min=18,max=80"`
	Phone   string    `validate:"len=11"`
	Hobby   []string  `validate:"required,gt=0,lte=3"`
	Company string    `validate:"oneof=tongfang youyun"`
}

func main() {
	user := &User {
		Name: "ZhangYaJun",
		Age:  20,
		Phone: "123456789",
		Hobby: []string{
			"write code",
			"write code",
			"write code",
			"write code",
		},
		Company: "hubeidaxue",
	}
	fmt.Println("Start")

	// 创建一个实例
	validate :=  validator.New()

	// 检测结构体
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err)
		// error_1  User.Phone   failed on the 'len' tag
		// error_2  User.Hobby   failed on the 'lte' tag
		// error_3  User.Company failed on the 'oneof' tag
	}

	fmt.Println("End")
}