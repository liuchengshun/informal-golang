package main

import (
	"log"
	"fmt"

	"crypto/aes"
	"crypto/cipher"
)

type CrypticBody struct {
	password      string
	adminToken    string
	userToken     string
	passwordEpt   []byte
	adminTokenEpt []byte
	userTokenEpt  []byte
	iv            []byte
}

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

var key = []byte("0123456789ABCDEF")

func Encryption(iv []byte, key []byte, plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cryptic: create instance of Cipher %v", err)
		return nil, err
	}
	cipherText := make([]byte, len(plainText))

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText, plainText)
	return cipherText, nil
}

func Decryption(iv []byte, key []byte, cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cryptic: create instance of Cipher %v", err)
		return nil, err
	}
	plainText := make([]byte, len(cipherText))

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plainText, cipherText)
	return plainText, nil
}

func (c *CrypticBody) Encrypt() {
	password := []byte(c.password)
	adminToken := []byte(c.adminToken)
	userToken := []byte(c.userToken)

	if len(password) != 0 {
		cipherPwd, err := Encryption(c.iv, key, password)
		if err != nil {
			log.Fatalf("encoding password error: %v", err)
			return
		}
		c.passwordEpt = cipherPwd
	}
	if len(adminToken) != 0 {
		cipherAdminToken, err := Encryption(c.iv, key, adminToken)
		if err != nil {
			log.Fatalf("encoding adminToken error: %v", err)
			return
		}
		c.adminTokenEpt = cipherAdminToken
	}
	if len(userToken) != 0 {
		cipherUserToken, err := Encryption(c.iv, key, userToken)
		if err != nil {
			log.Fatalf("encoding userToken error: %v", err)
			return
		}
		c.userTokenEpt = cipherUserToken
	}
}

func (c *CrypticBody) Decrypt() {
	if len(c.passwordEpt) != 0 {
		plainPwd, err := Decryption(c.iv, key, c.passwordEpt)
		if err != nil {
			log.Fatalf("decoding password error: %v", err)
			return
		}
		c.password = fmt.Sprintf("%s", plainPwd)
	}
	if len(c.adminTokenEpt) != 0 {
		plainAdminToken, err := Decryption(c.iv, key, c.adminTokenEpt)
		if err != nil {
			log.Fatalf("decoding adminToken error: %v", err)
			return
		}
		c.password = fmt.Sprintf("%s", plainAdminToken)
	}
	if len(c.userTokenEpt) != 0 {
		plainUserToken, err := Decryption(c.iv, key, c.userTokenEpt)
		if err != nil {
			log.Fatalf("decoding userToken error: %v", err)
			return
		}
		c.password = fmt.Sprintf("%s", plainUserToken)
	}
}

func main() {
	plainText := []byte("liuchengshun")
	cipherText, err := Encryption(commonIV, key, plainText)
	if err != nil {
		log.Fatalf("running main error: %v", err)
	}
	fmt.Printf("liuchengshun=>%x\n", cipherText)

	NewPlainText, err := Decryption(commonIV, key, cipherText)
	if err != nil {
		log.Fatalf("running main error: %v", err)
	}
	fmt.Printf("%x=>%s\n", cipherText, NewPlainText)
}