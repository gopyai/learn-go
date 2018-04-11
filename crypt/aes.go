package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"vos/onerror"
)

func AESGenerateKey() (key []byte) {
	return GenerateKey(32)
}

func AESEncrypt(key, msg []byte) (ciphermsg []byte) {
	block, e := aes.NewCipher(key)
	onerror.Panic(e)
	ciphermsg = make([]byte, aes.BlockSize+len(msg))
	iv := ciphermsg[:aes.BlockSize]
	_, e = rand.Read(iv)
	onerror.Panic(e)
	cipher.NewCTR(block, iv).XORKeyStream(ciphermsg[aes.BlockSize:], msg)
	return
}

func AESDecrypt(key, ciphermsg []byte) (msg []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	iv := ciphermsg[:aes.BlockSize]
	msg = make([]byte, len(ciphermsg)-aes.BlockSize)
	cipher.NewCTR(block, iv).XORKeyStream(msg, ciphermsg[aes.BlockSize:])
	return
}
