package main

import (
	"fmt"
	"strings"
)

func breakSensitivity(mode string, value string) string {
	if mode == "sms" {
		return value[:3] + "****" + value[7:]
	} else {
		v := strings.Split(value, "@")
		return "****@" + v[1]
	}
}

func main() {
	phone := "15171572765"
	result := breakSensitivity("sms", phone)
	fmt.Println("result:", result)

	email := "113608@qq.com"
	result2 := breakSensitivity("email", email)
	fmt.Println("result2:", result2)
}
