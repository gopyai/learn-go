package crypt

import (
	"crypto/sha256"
)

func SHA256Hash(msg []byte) []byte {
	h := sha256.New()
	h.Write(msg)
	return h.Sum(nil)
}
