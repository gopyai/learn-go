package main

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
)

func svr() {
	cerCA_b, _ := ioutil.ReadFile("ca.cer")
	cerCA, _ := x509.ParseCertificate(cerCA_b)
	keyCA, _ := ioutil.ReadFile("ca.key")
	privCA, _ := x509.ParsePKCS1PrivateKey(keyCA)

	poolCAs := x509.NewCertPool()
	poolCAs.AddCert(cerCA)

	cerTLS := tls.Certificate{
		Certificate: [][]byte{cerCA_b},
		PrivateKey:  privCA,
	}

	config := tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cerTLS},
		ClientCAs:    poolCAs,
	}
	config.Rand = rand.Reader
	service := "0.0.0.0:443"
	listener, err := tls.Listen("tcp", service, &config)
	panicIf(err)
	log.Print("server: listening")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept: %s", err)
			break
		}
		defer conn.Close()
		log.Printf("server: accepted from %s", conn.RemoteAddr())
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 512)
	for {
		log.Print("server: conn: waiting")
		n, err := conn.Read(buf)
		if err != nil {
			if err != nil {
				log.Printf("server: conn: read: %s", err)
			}
			break
		}

		tlscon, ok := conn.(*tls.Conn)
		if ok {
			state := tlscon.ConnectionState()
			sub := state.PeerCertificates[0].Subject
			log.Println(sub)
		}

		log.Printf("server: conn: echo %q\n", string(buf[:n]))
		n, err = conn.Write(buf[:n])

		n, err = conn.Write(buf[:n])
		log.Printf("server: conn: wrote %d bytes", n)

		if err != nil {
			log.Printf("server: write: %s", err)
			break
		}
	}
	log.Println("server: conn: closed")
}
