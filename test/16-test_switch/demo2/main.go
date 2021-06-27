package main

import (
	"fmt"
	"time"
)

func main() {
	way := "sms"

	// switch {
	// case way == "sms":
	// 	go func() {
	// 		fmt.Println("sms")
	// 	}()
	// case way == "email":
	// 	go func() {
	// 		fmt.Println("email")
	// 	}()
	// }

	switch way {
	case "sms":
		go func() {
			fmt.Println("sms")
		}()
	case "email":
		go func() {
			fmt.Println("email")
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("end")
}
