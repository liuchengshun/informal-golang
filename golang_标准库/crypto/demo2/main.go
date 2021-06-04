package main

import (
	"crypto/md5"
	"fmt"
)

func Md5Encrypt(value string) string {
	has := md5.Sum([]byte(value))
	return fmt.Sprintf("%x", has)
}

func main() {
	value := "_n091!@~"
	resutlt := Md5Encrypt(value)
	fmt.Println("result:", resutlt)
}
