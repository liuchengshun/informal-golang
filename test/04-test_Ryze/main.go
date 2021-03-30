package main

import (
	"fmt"
	"net/url"
)

func main() {
	str := "redis.url"
	URL, err := url.Parse(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("URL.Host:", URL.Host)
	fmt.Println("")
}