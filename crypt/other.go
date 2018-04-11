package crypt

import (
	"crypto/rand"
	"vos/onerror"
)

func GenerateKey(numBytes int) (key []byte) {
	key = make([]byte, numBytes)
	_, e := rand.Read(key)
	onerror.Panic(e)
	return
}
