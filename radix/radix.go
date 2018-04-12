package main

import (
	"fmt"

	"github.com/mediocregopher/radix.v2/redis"
)

func main() {
	fmt.Println("Hello World!")

	cli, e := redis.Dial("tcp", "localhost:6379")
	panicIf(e)
	defer func() {
		panicIf(cli.Close())
	}()

	r := cli.Cmd("KEYS", "*")
	ss, e := r.List()
	panicIf(e)
	fmt.Println(ss)
}

func panicIf(e error) {
	if e != nil {
		panic(e)
	}
}
