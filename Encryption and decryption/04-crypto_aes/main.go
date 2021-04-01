package main

import (
	"log"
	"io"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
)

func AesEncryptCFB(plainText string) (cipherStr string) {
	plainBytes := []byte(plainText)
	cipherBytes := make([]byte, aes.BlockSize*2+len(plainBytes))
	key := cipherBytes[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		log.Fatal("generate random key error:", err)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal("create instance of encryption error:", err)
		return
	}

	iv := cipherBytes[aes.BlockSize:aes.BlockSize*2]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatal("generate random iv error:", err)
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherBytes[aes.BlockSize*2:], plainBytes)
	return hex.EncodeToString(cipherBytes)
}

func AesDecryptCFB(cipherStr string) (plainText string) {
	cipherBytes, _ := hex.DecodeString(cipherStr)
	key := cipherBytes[:aes.BlockSize]

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal("create instance of encryption error:", err)
		return
	}
	iv := cipherBytes[aes.BlockSize : aes.BlockSize*2]
	plainBytes := cipherBytes[aes.BlockSize*2:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plainBytes, plainBytes)
	return string(plainBytes)
}

func main() {
	plainText := "Hello World"
	cipherStr := AesEncryptCFB(plainText)
	log.Println("cipherStr:", cipherStr)

	plain := AesDecryptCFB(cipherStr)
	log.Println("解密结果：", plain)
}