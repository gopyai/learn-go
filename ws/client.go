package main

import (
	"net/url"
	"github.com/gorilla/websocket"
	"fmt"
)

func client() {
	u := url.URL{Scheme: "ws", Host: serverAddr, Path: "/echo"}
	fmt.Println("Client connect to:", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	panicIf(err)
	defer c.Close()
	panicIf(c.WriteMessage(websocket.TextMessage, []byte("Hello there ...")))
	msgType, msg, err := c.ReadMessage()
	panicIf(err)
	fmt.Printf("Client: msgType:%v, msg:%s\n", msgType, msg)
	panicIf(c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Asik")))

	if err = c.WriteMessage(websocket.TextMessage, []byte("Hello there again ...")); err != nil {
		fmt.Println("Expected error:", err)
	}
}
