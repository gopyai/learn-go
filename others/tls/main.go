package main

import "vos/lib"

func main() {
	go svr()
	go cli()
	lib.WaitKeyEnter()
}
