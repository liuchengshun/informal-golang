package main

import (
	"fmt"
	"reflect"
)

type dog struct {
	name string
	age  int
}

func main() {
	var d = []byte("key")

	var str string = "key"

	equal := reflect.DeepEqual(d, []byte(str))
	fmt.Println("equal:", equal)

	var jinmao dog = dog{
		name: "wangcai",
		age:  3,
	}

	var hashiqi dog = dog{
		name: "chaijia",
		age:  3,
	}

	var tudog dog = dog{
		name: "wangcai",
		age:  3,
	}

	equal = reflect.DeepEqual(jinmao, hashiqi)
	fmt.Println("equal for dog: ", equal)

	equal = reflect.DeepEqual(jinmao, jinmao)
	fmt.Println("equal for dog: ", equal)

	equal = reflect.DeepEqual(jinmao, tudog)
	fmt.Println("equal for dog: ", equal)

}
