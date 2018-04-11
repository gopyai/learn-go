package main

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"fmt"
)

var (
	serverAddr = "localhost:8080"
	upgrader   = websocket.Upgrader{}
)

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	isErr(err)
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		isErr(err)
		fmt.Println("Server read:", string(message))
		isErr(c.WriteMessage(mt, message))
	}
}

func server() {
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
