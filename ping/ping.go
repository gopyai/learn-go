package main

import (
	"fmt"
	"os"
	"time"

	"encoding/binary"
	"net"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const (
	ProtocolICMP = 1
)

func main() {
	fmt.Println("Requires super user priviledge")
	ipAddrs := []*net.IPAddr{
		{IP: net.IP{10, 10, 10, 10}},
		{IP: net.IP{172, 217, 24, 4}},
		{IP: net.IP{127, 0, 0, 1}},
	}
	doPing(len(ipAddrs),
		func(seq int) net.Addr {
			return ipAddrs[seq]
		},
		func(seq int) {
			fmt.Printf("Reply from %v\n", ipAddrs[seq])
		})
}

func doPing(numOfIPAddrs int, getIPAddr func(int) net.Addr, replyFrom func(int)) {
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	panicIf(err)
	defer conn.Close()
	panicIf(conn.SetReadDeadline(time.Now().Add(3 * time.Second)))

	// Send echo
loopEcho:
	for seq := 0; seq < numOfIPAddrs; seq++ {
		ipAddr := getIPAddr(seq)

		msg := icmp.Message{
			Type: ipv4.ICMPTypeEcho, Code: 0,
			Body: &icmp.Echo{
				ID:   os.Getpid() & 0xffff, Seq: seq,
				Data: []byte("HELLO-R-U-THERE"),
			},
		}

		b, err := msg.Marshal(nil)
		panicIf(err)
		n, err := conn.WriteTo(b, ipAddr)
		if err != nil {
			switch err.(type) {
			case *net.OpError:
				continue loopEcho
			default:
				panic(err)
			}
		}
		if n != len(b) {
			panic(fmt.Errorf("got %v; want %v", n, len(b)))
		}
		fmt.Println("Ping:", ipAddr, "sequence:", seq)
	}

	// Receive reply
	rb := make([]byte, 1500)
loopReply:
	for loop := 0; loop < numOfIPAddrs; loop++ {
		n, peer, err := conn.ReadFrom(rb)
		if err != nil {
			switch err.(type) {
			case *net.OpError:
				continue loopReply
			default:
				panic(err)
			}
		}
		rm, err := icmp.ParseMessage(ProtocolICMP, rb[:n])
		panicIf(err)
		switch rm.Type {
		case ipv4.ICMPTypeEchoReply:
		default:
			panic(fmt.Errorf("got %+v from %v; want echo reply", rm, peer))
		}
		b, err := rm.Body.Marshal(ProtocolICMP)
		panicIf(err)
		seq := int(binary.BigEndian.Uint16(b[2:4]))
		replyFrom(seq)
	}
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
