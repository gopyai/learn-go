// main
package main

import (
	"fmt"
	"net"
)

func main() {
	intfs, _ := net.Interfaces()
	for _, i := range intfs {
		addrs, _ := i.Addrs()
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					fmt.Println(i.Name, ":", ipnet.IP.String())
				}
			}
		}
	}
}
