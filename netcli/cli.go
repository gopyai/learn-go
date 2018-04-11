// cli
package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
	"v/bg"
	"v/err"
)

func main() {
	c, e := net.Dial("tcp", "localhost:8001")
	err.Panic(e)

	bg.Wg.Add(2)
	go w(c)
	go r(c)
	bg.Wg.Wait()
}

func w(c net.Conn) {
	defer bg.Wg.Done()
	defer c.Close()

	// Write data
	i := 0
	for {
		_, e := c.Write([]byte(fmt.Sprintf("From cli %d\n", i)))
		if e != nil {
			fmt.Println("ERROR write:", e)
			return
		}
		time.Sleep(time.Second)
		i++
	}
}

func r(c net.Conn) {
	defer bg.Wg.Done()
	defer c.Close()

	// Read data
	for {
		s, e := bufio.NewReader(c).ReadString('\n')
		if e != nil {
			fmt.Println("ERROR read:", e)
			return
		}
		fmt.Print(s)
	}
}
