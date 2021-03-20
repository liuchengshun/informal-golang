package main

import (
	"fmt"
	"crypto/rand"
	"encoding/base64"

)

func main() {
	b := make([]byte, 32)
	intNum, err := rand.Read(b)
	fmt.Println("intNum: ", intNum,"err: ", err)
	base := base64.URLEncoding.EncodeToString(b)
	fmt.Println("base64:", base)
}