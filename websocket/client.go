package main

import (
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"
	"vos/onerror"

	"github.com/gorilla/websocket"
)

func runClient(port int) {
	//d := websocket.Dialer{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	d := websocket.DefaultDialer
	s := "ws"
	u := url.URL{Scheme: s, Host: fmt.Sprintf("localhost:%d", port), Path: "/"}

	c, _, e := d.Dial(u.String(), nil)
	onerror.Panic(e)
	defer c.Close()

	c.SetReadLimit(5)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			onerror.Fatal(c.WriteMessage(websocket.BinaryMessage, []byte{0, 1, 2, 3, 4, 5}))
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_, p, e := c.ReadMessage()
			onerror.Fatal(e)
			log.Printf("Received: %v", p)
		}
	}()

	wg.Wait()
}
