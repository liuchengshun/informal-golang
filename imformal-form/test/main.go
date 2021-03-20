package main

import (
	"crypto/md5"
	"fmt"
	"time"
	"strconv"
)

func main() {
	timestamp := strconv.Itoa(time.Now().Nanosecond())
	fmt.Println("time:", timestamp)
	testTime := time.Now().Unix()
	fmt.Println("testTime:", strconv.FormatInt(testTime, 10))
	testMD5 := md5.New()
	fmt.Println("testMD5:", fmt.Sprintf("%x", testMD5.Sum(nil)))
}