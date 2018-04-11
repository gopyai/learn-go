// Stomp2 adalah untuk testing stomp dengan PHP

package main

import (
	"fmt"
	"time"
)

func main() {
	go Producer("localhost", "/queue/calculator", func(opName string, in []byte) []byte {
		return []byte(string(in) + " world")
	})

	consumer := NewConsumer("localhost", "/queue/my/response")
	out := consumer.Request("/queue/calculator", "add", []byte("Hello wow"))
	fmt.Println("Hore:", string(out))

	time.Sleep(time.Second * 10)

}
