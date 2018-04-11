package crypt

import (
	"crypto/hmac"
	"crypto/sha256"
)

func HMACHash(key, msg []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(msg)
	return h.Sum(nil)
}

func HMACEqual(mac1, mac2 []byte) bool {
	return hmac.Equal(mac1, mac2)
}
