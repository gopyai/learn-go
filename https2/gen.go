package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

const (
	rsaBits      = 2048
	organization = "OrgName"
	validFor     = 1 // minutes
)

func genCer() {
	//
	// CA
	//

	caPriv, err := rsa.GenerateKey(rand.Reader, rsaBits)
	if err != nil {
		log.Fatalf("failed to generate private key: %s", err)
	}

	notBefore := time.Now()
	notAfter := notBefore.Add(time.Minute * validFor)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("failed to generate serial number: %s", err)
	}

	caCer := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{organization},
		},

		NotBefore: notBefore,
		NotAfter:  notAfter,

		BasicConstraintsValid: true,
		IsCA: true,

		KeyUsage: x509.KeyUsageDigitalSignature |
			x509.KeyUsageCertSign,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
			x509.ExtKeyUsageServerAuth},

		//SubjectKeyId:          []byte{1, 2, 3, 4, 5},
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, caCer, caCer, &caPriv.PublicKey, caPriv)
	if err != nil {
		log.Fatalf("Failed to create certificate: %s", err)
	}

	certOut, err := os.Create("ca.cer")
	if err != nil {
		log.Fatalf("failed to open ca.cer for writing: %s", err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: caBytes})
	certOut.Close()
	log.Print("written ca.cer\n")

	keyOut, err := os.OpenFile("ca.key", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Print("failed to open ca.key for writing:", err)
		return
	}
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(caPriv)})
	keyOut.Close()
	log.Print("written ca.key\n")

	//
	// Server
	//

	hosts := []string{"localhost"}

	//

	priv, err := rsa.GenerateKey(rand.Reader, rsaBits)
	if err != nil {
		log.Fatalf("failed to generate private key: %s", err)
	}

	notBefore = time.Now()
	notAfter = notBefore.AddDate(0, 0, validFor)

	serialNumberLimit = new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err = rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("failed to generate serial number: %s", err)
	}

	hostCer := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{organization},
		},

		NotBefore:             notBefore,
		NotAfter:              notAfter,
		BasicConstraintsValid: false,

		KeyUsage: x509.KeyUsageDigitalSignature |
			x509.KeyUsageCertSign,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
			x509.ExtKeyUsageServerAuth},

		//SubjectKeyId:          []byte{1, 2, 3, 4, 5},
	}

	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			hostCer.IPAddresses = append(hostCer.IPAddresses, ip)
		} else {
			hostCer.DNSNames = append(hostCer.DNSNames, h)
		}
	}

	hostBytes, err := x509.CreateCertificate(rand.Reader, hostCer, caCer, &priv.PublicKey, caPriv)
	if err != nil {
		log.Fatalf("Failed to create certificate: %v", err)
	}

	certOut, err = os.Create("host.cer")
	if err != nil {
		log.Fatalf("failed to open host.cer for writing: %v", err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: hostBytes})
	certOut.Close()
	log.Print("written host.cer\n")

	keyOut, err = os.OpenFile("host.key", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Print("failed to open host.key for writing:", err)
		return
	}
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	keyOut.Close()
	log.Print("written host.key\n")

	//
	// Check signature
	//

	if caCer, err = x509.ParseCertificate(caBytes); err != nil {
		log.Fatalf("Error parse certificate CA: %v", err)
	}
	if hostCer, err = x509.ParseCertificate(hostBytes); err != nil {
		log.Fatalf("Error parse certificate Host: %v", err)
	}

	if err = hostCer.CheckSignatureFrom(caCer); err != nil {
		log.Fatalf("Invalid signature")
	}
}
