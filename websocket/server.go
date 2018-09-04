package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"bitbucket.org/stefarf/iferr"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, e := upgrader.Upgrade(w, r, nil)
	iferr.Panic(e)
	defer conn.Close()

	// conn.SetReadLimit(1000)

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
	addr := fmt.Sprintf(":%d", port)
	iferr.Panic(http.ListenAndServe(addr, http.HandlerFunc(wsHandler)))
	// iferr.Panic(http.ListenAndServeTLS(addr, certFile, keyFile, nil))
}
