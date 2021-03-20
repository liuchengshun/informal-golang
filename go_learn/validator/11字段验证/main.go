package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type contact struct {
	Phone  string
	Email  string
}

type User struct { 
    Name       string  `validate:"lte=20"` 
    Age        int     `validate:"min=20"`
	Phone      string  `validate:"necsfield=Contact.Phone"`
	Email      string  `validate:"eqcsfield=Contact.Email"`
	Contact    contact
    Password   string  `validate:"min=10,nefield=Name"`
    Password2  string  `validate:"eqfield=Password"`
}

func main() {
	user := &User {
		Name: "ZhangYaJun",
		Age: 20,
		Phone: "12244448888",
		Email: "132@fox.com",
		Contact: contact{
			Phone: "12244448888",
			Email: "132@fox.com",
		},
		Password: "111222333444",
		Password2: "111444777888",
	}
	fmt.Println("Start")

	// 创建一个实例
	validate :=  validator.New()

	// 检测结构体
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err)
		// error_1  User.Phone  'Phone' failed on the 'necsfield' tag
		// error_2  User.Password2   'Password2' failed on the 'eqfield' tag
	}

	fmt.Println("End")
}