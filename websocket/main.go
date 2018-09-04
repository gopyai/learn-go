package main

import (
	"github.com/gopyai/go-keyb"
)

func main() {
	go runServer(10000, "cert.pem", "key.pem")
	go runClient(10000)

	keyb.WaitKeyEnter()
}
