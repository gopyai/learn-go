package main

import (
	"flag"
	"fmt"
)

var (
	start = flag.Bool("start", false, "Start")
	stop  = flag.Bool("stop", false, "Stop")
)

func main() {
	flag.Parse()
	if *start {
		fmt.Println("Start")
	}
	if *stop {
		fmt.Println("Stop")
	}
}
