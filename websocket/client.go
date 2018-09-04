package main

import (
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"bitbucket.org/stefarf/iferr"
)

func runClient(port int) {
	// d := websocket.Dialer{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	d := websocket.DefaultDialer
	s := "ws"
	u := url.URL{Scheme: s, Host: fmt.Sprintf("localhost:%d", port), Path: "/"}

	c, _, e := d.Dial(u.String(), nil)
	iferr.Panic(e)
	defer c.Close()

	// c.SetReadLimit(1000)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			iferr.Exit(c.WriteMessage(websocket.BinaryMessage, []byte{0, 1, 2, 3, 4, 5}), "Fatal")
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_, p, e := c.ReadMessage()
			iferr.Exit(e, "Fatal")
			log.Printf("Received: %v", p)
		}
	}()

	wg.Wait()
}
