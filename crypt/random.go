package crypt

import (
	"crypto/rand"
	"devx/iferr"
	"encoding/base64"
)

func GenerateKey(numBytes int) (key []byte) {
	return GenerateRandomBytes(numBytes)
}

func GenerateRandomBytes(numBytes int) (key []byte) {
	key = make([]byte, numBytes)
	_, err := rand.Read(key)
	iferr.Panic(err)
	return
}

func GenerateRandomString(numBytes int) string {
	return base64.StdEncoding.EncodeToString(GenerateRandomBytes(numBytes))
}

func GenerateRandomStringURLSafe(numBytes int) string {
	return base64.URLEncoding.EncodeToString(GenerateRandomBytes(numBytes))
}
