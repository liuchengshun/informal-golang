package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	sliceOne := []string{"zhangsan", "lisi", "wangwu", "chenliu"}

	validate := validator.New()
	err := validate.Var(sliceOne, "max=5,dive,min=6")
	if err != nil {
		fmt.Println("err:", err) // error cliceOne[1]="lisi"的length小于6
	}

	sliceTwo := []string{}
	err = validate.Var(sliceTwo, "min=4,dive,required")
	if err != nil {
		fmt.Println(err)
		// error
		// sliceTwo[] 条目数为0，小于4 即error
		// sliceTwo[] 的value 为 "" 不满足required 即error
	}
}