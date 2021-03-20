package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name string `form:"name" json:"name" validate:"required,CustomerValidation"` //注意：required和CustomerValidation之间不能有空格，否则panic。CustomerValidation：自定义tag-函数标签
	Age  uint8  ` form:"age" json:"age" validate:"gte=0,lte=80"`                 //注意：gte=0和lte=80之间不能有空格，否则panic
}

var validate *validator.Validate

func main() {
	validate = validator.New()
	validate.RegisterValidation("CustomerValidation", CustomerValidationFunc) //注册自定义函数，前一个参数是struct里tag自定义，后一个参数是自定义的函数

	user := &User{
		Name: "jimmy",
		Age:  86,
	}

	fmt.Println("first value: ", user)
	err := validate.Struct(user)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}

	user.Name = "tom"
	user.Age = 29
	fmt.Println("second value: ", user)
	err = validate.Struct(user)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
}

// 自定义函数
func CustomerValidationFunc(f1 validator.FieldLevel) bool {
    // f1 包含了字段相关信息
    // f1.Field() 获取当前字段信息
    // f1.Param() 获取tag对应的参数
    // f1.FieldName() 获取字段名称
    
	return f1.Field().String() == "jimmy"
}