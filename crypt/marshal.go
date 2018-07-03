package crypt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
	"devx/iferr"
)

func MarshalPrivateKey(pri *rsa.PrivateKey) []byte {
	return x509.MarshalPKCS1PrivateKey(pri)
}

func UnmarshalPrivateKey(b []byte) (*rsa.PrivateKey, error) {
	return x509.ParsePKCS1PrivateKey(b)
}

func MarshalPublicKey(pub *rsa.PublicKey) ([]byte, error) {
	return x509.MarshalPKIXPublicKey(pub)
}

func UnmarshalPublicKey(b []byte) (pub *rsa.PublicKey, err error) {
	p, e := x509.ParsePKIXPublicKey(b)
	if e != nil {
		err = e
		return
	}
	return p.(*rsa.PublicKey), nil
}

func WritePrivateKeyToFile(f *os.File, pri *rsa.PrivateKey) error {
	return pem.Encode(f, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: MarshalPrivateKey(pri)})
}

func SavePrivateKey(fileName string, pri *rsa.PrivateKey) (err error) {
	f, e := os.Create(fileName)
	if e != nil {
		err = e
		return
	}
	defer f.Close()
	return WritePrivateKeyToFile(f, pri)
}

func ReadPrivateKeyFromFile(f *os.File) (pri *rsa.PrivateKey, err error) {
	b, e := ioutil.ReadAll(f)
	iferr.Panic(e)
	block, _ := pem.Decode(b)
	return UnmarshalPrivateKey(block.Bytes)
}

func LoadPrivateKey(fileName string) (pri *rsa.PrivateKey, err error) {
	f, e := os.Open(fileName)
	if e != nil {
		err = e
		return
	}
	defer f.Close()
	return ReadPrivateKeyFromFile(f)
}

func WritePublicKeyToFile(f *os.File, pub *rsa.PublicKey) (err error) {
	b, e := MarshalPublicKey(pub)
	if e != nil {
		err = e
		return
	}
	return pem.Encode(f, &pem.Block{Type: "PUBLIC KEY", Bytes: b})
}

func SavePublicKey(fileName string, pub *rsa.PublicKey) (err error) {
	f, e := os.Create(fileName)
	if e != nil {
		err = e
		return
	}
	defer f.Close()
	return WritePublicKeyToFile(f, pub)
}

func ReadPublicKeyFromFile(f *os.File) (pub *rsa.PublicKey, err error) {
	b, e := ioutil.ReadAll(f)
	iferr.Panic(e)
	block, _ := pem.Decode(b)
	return UnmarshalPublicKey(block.Bytes)
}

func LoadPublicKey(fileName string) (pub *rsa.PublicKey, err error) {
	f, e := os.Open(fileName)
	if e != nil {
		err = e
		return
	}
	defer f.Close()
	return ReadPublicKeyFromFile(f)
}
