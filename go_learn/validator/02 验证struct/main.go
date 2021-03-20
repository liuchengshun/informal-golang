package main

import (
	"fmt"
    "github.com/go-playground/validator/v10"
)

type Person struct {
    Name    string  `validate:"check"`
}

// 自定义验证器
func CustomValidator(f validator.FieldLevel) bool {
	fmt.Println(f.Field().String())
    return f.Field().String() == "liu"
}

func main() {
    person := &Person {
        Name : "liucheng",
    }
	
    // new一个实例
    v := validator.New()

    // 注册自定义tag
    v.RegisterValidation("check", func(f validator.FieldLevel) bool {
		fmt.Println(f.Field().String())
		return f.Field().String() == "liu"
	})

    // 检测结构体
    err := v.Struct(person)
	if err != nil {
		fmt.Println("struct filed value error:", err)
		return
	}
}