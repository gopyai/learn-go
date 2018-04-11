package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"math/big"
	"time"
)

func genCA() (priv *rsa.PrivateKey, cer *x509.Certificate, key, cer_b []byte) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	isErr(err)

	cer = &x509.Certificate{
		SerialNumber: serialNumber,

		//		Subject: pkix.Name{
		//			Country:            []string{"Indonesia"},
		//			Organization:       []string{"Vostra"},
		//			OrganizationalUnit: []string{"?"},
		//		},

		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 7),

		//		SubjectKeyId:          []byte{1, 2, 3, 4, 5},

		BasicConstraintsValid: true,
		IsCA: true,

		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	}

	priv, _ = rsa.GenerateKey(rand.Reader, 2048)
	pub := &priv.PublicKey

	key = x509.MarshalPKCS1PrivateKey(priv)
	cer_b, err = x509.CreateCertificate(rand.Reader, cer, cer, pub, priv)
	isErr(err)

	return priv, cer, key, cer_b
}

func genSvr(privCA *rsa.PrivateKey, cerCA *x509.Certificate) (
	priv *rsa.PrivateKey, cer *x509.Certificate, key, cer_b []byte,
) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	isErr(err)

	cer = &x509.Certificate{
		SerialNumber: serialNumber,

		//		Subject: pkix.Name{
		//			Country:            []string{"Indonesia"},
		//			Organization:       []string{"Vostra"},
		//			OrganizationalUnit: []string{"?"},
		//		},

		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 7),

		//		SubjectKeyId:          []byte{1, 2, 3, 4, 5},

		BasicConstraintsValid: true,

		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	}

	priv, _ = rsa.GenerateKey(rand.Reader, 1024)
	pub := &priv.PublicKey

	key = x509.MarshalPKCS1PrivateKey(priv)
	cer_b, err = x509.CreateCertificate(rand.Reader, cer, cerCA, pub, privCA)
	isErr(err)

	return priv, cer, key, cer_b
}
