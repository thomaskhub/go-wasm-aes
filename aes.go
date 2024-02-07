package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/scrypt"
)

func GetKey(password []byte, salt []byte) ([]byte, error) {
	key, err := scrypt.Key(password, salt, 65536, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func Encrypt(password string, salt string, data string) string {
	dataBuffer := []byte(data)
	key32, err := GetKey([]byte(password), []byte(salt)) //return 32byte 256 bit key
	if err != nil {
		return ""
	}

	c, err := aes.NewCipher(key32)
	if err != nil {
		return ""
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return ""
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return ""
	}

	encData := gcm.Seal(nonce, nonce, dataBuffer, nil)

	return base64.StdEncoding.EncodeToString(encData)
}

func Decrypt(password string, salt string, data string) string {
	dataBuffer, _ := base64.StdEncoding.DecodeString(data)
	key32, err := GetKey([]byte(password), []byte(salt)) //return 32byte 256 bit key
	if err != nil {
		return ""
	}

	c, err := aes.NewCipher(key32)
	if err != nil {
		return ""
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return ""
	}

	nonce, cipherText := dataBuffer[:gcm.NonceSize()], dataBuffer[gcm.NonceSize():]
	decData, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return ""
	}

	return string(decData)
}
