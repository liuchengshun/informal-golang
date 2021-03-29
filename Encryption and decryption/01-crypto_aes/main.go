package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func main() {
	//需要去加密的字符串
	plaintext := []byte("liuchengshun")

	//aes的加密字符串
	key_text := "astaxie12798akljzmknm.ahkjkljl;k"
	fmt.Println(len(key_text))

	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}

	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	res := fmt.Sprintf("%x", ciphertext)
	fmt.Printf("type is %T, value is %s\n", res, res)
	fmt.Println("res []byte is :", []byte(res))
	fmt.Println("ciphertext:", ciphertext)
	fmt.Printf("%s=>%v\n", plaintext, ciphertext)

	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Println("plaintextCopy", plaintextCopy)
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)
}