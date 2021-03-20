package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	var mapone map[string]string

	mapone = map[string]string{"one": "zhangsan", "two": "lisi", "three": ""}

	validate := validator.New()
	err := validate.Var(mapone, "gte=3,dive,keys,eq=1|eq=2,endkeys,required")
	if err != nil {
		fmt.Println(err)
		// error
		// "one" 不满足 eq=1|eq=2
		// "two" 不满足 eq=1|eq=2
		// "three" 不满足 eq=1|eq=2
		// "three" 不满足 required
	}
}