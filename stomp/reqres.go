// reqhandler
package main

import (
	"fmt"
	"v/csrf"

	"github.com/go-stomp/stomp"
)

func requestResponse() {
	go proA()
	go sendToProA()
}

func sendToProA() error {
	conn, err := stomp.Dial("tcp", "localhost:61613")
	if err != nil {
		return err
	}

	for i := 0; i < 2; i++ {
		corId, err := csrf.GenerateRandomString(32)
		if err != nil {
			return err
		}
		err = conn.Send(
			"/pro/A",               // destination
			"application/json",     // content-type
			[]byte("Test message"), // body
			stomp.SendOpt.Header("reply-to", "/con/me"),
			stomp.SendOpt.Header("correlation-id", corId),
		)
		if err != nil {
			return err
		}
		fmt.Println("Send to /pro/A:", corId)
	}

	return conn.Disconnect()
}

func proA() error {
	conn, err := stomp.Dial("tcp", "localhost:61613")
	if err != nil {
		return err
	}

	sub, err := conn.Subscribe(
		"/pro/A",
		stomp.AckClientIndividual,
		stomp.SubscribeOpt.Id("proA"),
		//		stomp.SubscribeOpt.Header("persistent", "true"),
	)
	if err != nil {
		return err
	}

	for {
		msg, err := sub.Read()
		if err != nil {
			return err
		}
		repTo := msg.Header.Get("reply-to")
		corId := msg.Header.Get("correlation-id")
		fmt.Println(repTo, corId, string(msg.Body))
		err = conn.Ack(msg)
		if err != nil {
			return err
		}
	}

	return conn.Disconnect()
}
