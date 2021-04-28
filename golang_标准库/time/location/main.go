package main

import (
	"fmt"
	"time"
)

func main() {
	var w time.Weekday
	fmt.Println("weekey:", w.String())

	_, month, day := time.Now().Date()
	fmt.Println("month:", month)
	fmt.Println("day:", day)
	fmt.Println("Now:", time.Now())

	timeNow := time.Now()
	fmt.Println("location:", timeNow.Local())
	fmt.Println("UTC time:", timeNow.UTC())
}
