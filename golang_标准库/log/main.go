package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "|||", 10)
	logger.Printf("wuhu qifei %v\n", "headfasd")
}
