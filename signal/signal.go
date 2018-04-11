// signal
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

var (
	terminated = false
)

func init() {
	// Disable interrupt signal
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		for {
			fmt.Println(<-c)
			terminated = true
		}
	}()
}

func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("Program is running")
	}
}
