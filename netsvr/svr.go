// net
package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
	"devx/iferr"
)

func main() {
	l, e := net.Listen("tcp", ":8001")
	iferr.Panic(e)
	defer l.Close()

	for {
		c, e := l.Accept()
		if e != nil {
			fmt.Println("ERROR listener:", e)
			return
		}
		go w(c)
		go r(c)
	}
}

func w(c net.Conn) {
	defer c.Close()

	// Write data
	i := 0
	for {
		_, e := c.Write([]byte(fmt.Sprintf("From svr %d\n", i)))
		if e != nil {
			fmt.Println("ERROR write:", e)
			return
		}
		time.Sleep(time.Second)
		i++
	}
}

func r(c net.Conn) {
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

// func sockServer() {
// 	defer bg.Wg.Done()

// 	l, e := net.Listen("tcp", ":8001")
// 	iferr.Panic(e)
// 	defer l.Close()

// 	// Routine to close listener
// 	bg.Wg.Add(1)
// 	go func() {
// 		defer bg.Wg.Done()
// 		select {
// 		case <-bg.ChStop:
// 			l.Close()
// 		}
// 	}()

// 	// Accept looping
// 	for !bg.Stopped {
// 		c, e := l.Accept()
// 		if e != nil {
// 			fmt.Println("Error accept:", e)
// 			time.Sleep(time.Second)
// 			continue
// 		}

// 		bg.Wg.Add(1)
// 		go sockHandler(c)
// 	}
// }

// func sockHandler(c net.Conn) {
// 	defer bg.Wg.Done()
// 	defer c.Close()

// 	// Routine to close connection routine
// 	bg.Wg.Add(1)
// 	go func() {
// 		defer bg.Wg.Done()
// 		select {
// 		case <-bg.ChStop:
// 			c.Close()
// 		}
// 	}()

// 	// Write data
// 	bg.Wg.Add(1)
// 	go func() {
// 		defer bg.Wg.Done()
// 		for {
// 			time.Sleep(time.Second * 5)
// 			_, e := c.Write([]byte("TEST"))
// 			iferr.Panic(e)
// 		}
// 	}()

// 	// Read data
// 	for {
// 		s, e := bufio.NewReader(c).ReadString('\n')
// 		iferr.Panic(e)

// 		x := strings.Split(s, "\t")
// 		if len(x) != 2 {
// 			fmt.Println("ERROR:", x)
// 			continue
// 		}

// 		switch x[0] {
// 		case "REG":
// 			fmt.Println(x)
// 		case "TRG":
// 			fmt.Println(x)
// 		default:
// 			fmt.Println("ERROR:", x)
// 			continue
// 		}
// 	}
// }

// func sockClient(name string) {
// 	defer wg.Done()
// 	c, e := net.Dial("tcp", "localhost:8001")
// 	iferr.Panic(e)
// 	defer c.Close()

// 	for {
// 		b, e := json.Marshal(&msg)
// 		iferr.Panic(e)
// 		b = append(b, '\n')
// 		_, e = c.Write(b)
// 		iferr.Panic(e)
// 		time.Sleep(time.Second)
// 		msg.Age++
// 	}
// }
