package main

import (
	"fmt"
	"log"
	"net/http"
	"vos/onerror"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, e := upgrader.Upgrade(w, r, nil)
	onerror.Panic(e)
	defer conn.Close()

	for {
		msgType, p, e := conn.ReadMessage()
		if e != nil {
			log.Println(e)
			return
		}

		if e = conn.WriteMessage(msgType, p); e != nil {
			log.Println(e)
			return
		}
	}
}

func runServer(port int, certFile, keyFile string) {
	http.HandleFunc("/", wsHandler)
	addr := fmt.Sprintf(":%d", port)
	onerror.Panic(http.ListenAndServe(addr, nil))
	//onerror.Panic(http.ListenAndServeTLS(addr, certFile, keyFile, nil))
}
