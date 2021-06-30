package main

import (
	"log"
	"time"
)

func main() {
	go func() {
		log.Println("hello1")
		go func() {
			log.Println("hello2")
		}()
	}()
	time.Sleep(3 * time.Second)
}
