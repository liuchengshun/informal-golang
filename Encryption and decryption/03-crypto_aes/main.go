package main

import (
	"log"
	"io"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
)

var key = []byte("0123456789ABCDEF")

func AesEncryptCFB(plainText string) (cipherStr string) {
	plainBytes := []byte(plainText)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal("create instance of encryption error:", err)
		return
	}
	cipherBytes := make([]byte, aes.BlockSize+len(plainBytes))
	iv := cipherBytes[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatal("generate random number error:", err)
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherBytes[aes.BlockSize:], plainBytes)
	return hex.EncodeToString(cipherBytes)
}

func AesDecryptCFB(cipherStr string) (plainText string) {
	cipherBytes, _ := hex.DecodeString(cipherStr)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal("create instance of encryption error:", err)
		return
	}
	iv := cipherBytes[:aes.BlockSize]
	plainBytes := cipherBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plainBytes, cipherBytes[aes.BlockSize:])
	return string(plainBytes)
}

func main() {
	plainText := "Hello World"
	cipherStr := AesEncryptCFB(plainText)
	log.Println("cipherStr:", cipherStr)

	plain := AesDecryptCFB(cipherStr)
	log.Println("解密结果：", plain)
}