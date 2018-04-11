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
	isErr(err)
	defer c.Close()
	isErr(c.WriteMessage(websocket.TextMessage, []byte("Hello there ...")))
	msgType, msg, err := c.ReadMessage()
	isErr(err)
	fmt.Printf("Client: msgType:%v, msg:%s\n", msgType, msg)
	isErr(c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Asik")))

	if err = c.WriteMessage(websocket.TextMessage, []byte("Hello there again ...")); err != nil {
		fmt.Println("Expected error:", err)
	}
}
