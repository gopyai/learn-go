package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func client(caCer *x509.Certificate) {
	u := url.URL{Scheme: "wss", Host: serverAddr, Path: "/echo"}
	fmt.Println("Client connect to:", u.String())

	// Create certificate pool contains self signed CA certificate
	root := x509.NewCertPool()
	root.AddCert(caCer)

	// Create dialer using self signed pool
	selfSignedDialer := websocket.Dialer{TLSClientConfig: &tls.Config{RootCAs: root}}

	// Dial, write message, and close
	c, _, err := selfSignedDialer.Dial(u.String(), nil)
	panicIf(err)
	defer c.Close()
	panicIf(c.SetReadDeadline(time.Now().Add(time.Second * 3)))

	panicIf(c.WriteMessage(websocket.TextMessage, []byte("Hello there ...")))

	msgType, msg, err := c.ReadMessage()
	if err != nil {
		switch err.(type) {
		case *websocket.CloseError:
			c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			return
		}
		panic(err)
	}
	fmt.Printf("Client: msgType:%v, msg:%s\n", msgType, msg)

	panicIf(c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Asik")))

	time.Sleep(time.Second)

	// Will error when write message after websocket is closed
	if err = c.WriteMessage(websocket.TextMessage, []byte("Hello there again ...")); err != nil {
		fmt.Println("Expected error:", err)
	}
}
