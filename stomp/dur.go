// pubsub
package main

import (
	"fmt"
	"time"

	"github.com/go-stomp/stomp"
)

func durable() {
	go subscribeDur("Subs#1")
	go subscribeDur("Subs#2")
	time.Sleep(time.Millisecond * 100)
	go publishDur()
}

func publishDur() error {
	conn, err := stomp.Dial("tcp", "localhost:61613")
	if err != nil {
		return err
	}

	for i := 0; i < 2; i++ {
		err = conn.Send(
			"/topic/Durable",                 // destination
			"application/json",               // content-type
			[]byte(fmt.Sprintf("Msg#%d", i)), // body
		)
		if err != nil {
			return err
		}
	}
	fmt.Println("Publish to /topic/Durable")

	return conn.Disconnect()
}

func subscribeDur(clientId string) error {
	conn, err := stomp.Dial("tcp", "localhost:61613", stomp.ConnOpt.Header("client-id", clientId))
	if err != nil {
		fmt.Println(err)
		return err
	}

	sub, err := conn.Subscribe(
		"/topic/Durable",
		stomp.AckClientIndividual,
		stomp.SubscribeOpt.Header("activemq.subscriptionName", clientId),
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for {
		msg, err := sub.Read()
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("Client ID:", clientId, "Message:", string(msg.Body))
		err = conn.Ack(msg)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return conn.Disconnect()
}
