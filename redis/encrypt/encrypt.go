package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"log"
	"io"
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
		log.Fatal("generate random iv error:", err)
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
	stream.XORKeyStream(plainBytes, plainBytes)
	return string(plainBytes)
}
