package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	serverAddr = "localhost:8080"
	upgrader   = websocket.Upgrader{}
)

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	panicIf(err)
	defer c.Close()

	// panicIf(c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")))

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			switch err.(type) {
			case *websocket.CloseError:
				fmt.Println("Receive close message")
				return
			}
			panic(err)
		}
		panicIf(err)
		fmt.Println("Server read:", string(message))

		// Will error if write message after sending close frame, see above remarked close frame
		panicIf(c.WriteMessage(mt, message))
		//_ = mt
	}
}

func server() {
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServeTLS(serverAddr, "server.cer", "server.key", nil))
}
