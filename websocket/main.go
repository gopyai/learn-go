package main

import "vos/lib"

func main() {
	go runServer(10000, "cert.pem", "key.pem")
	go runClient(10000)

	lib.WaitKeyEnter()
}
