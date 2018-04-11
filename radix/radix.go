// radix
package main

import (
	"fmt"

	"github.com/mediocregopher/radix.v2/redis"
)

func main() {
	fmt.Println("Hello World!")

	cli, e := redis.Dial("tcp", "localhost:6379")
	ifErr(e)
	defer func() {
		ifErr(cli.Close())
	}()

	r := cli.Cmd("KEYS", "*")
	ss, e := r.List()
	ifErr(e)
	fmt.Println(ss)
}

func ifErr(e error) {
	if e != nil {
		panic(e)
	}
}
