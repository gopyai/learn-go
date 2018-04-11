// pubsub
package main

import (
	"fmt"
	"time"

	"github.com/go-stomp/stomp"
)

func publishSubscribe() {
	go subscribe("Subs#1")
	go subscribe("Subs#2")
	time.Sleep(time.Millisecond * 100)
	go publish()
}

func publish() error {
	conn, err := stomp.Dial("tcp", "localhost:61613")
	if err != nil {
		return err
	}

	for i := 0; i < 2; i++ {
		err = conn.Send(
			"/topic/News",                    // destination
			"application/json",               // content-type
			[]byte(fmt.Sprintf("Msg#%d", i)), // body
		)
		if err != nil {
			return err
		}
	}
	fmt.Println("Publish to /topic/News")

	return conn.Disconnect()
}

func subscribe(id string) error {
	conn, err := stomp.Dial("tcp", "localhost:61613")
	if err != nil {
		fmt.Println(err)
		return err
	}

	sub, err := conn.Subscribe(
		"/topic/News",
		stomp.AckClientIndividual,
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
		fmt.Println("ID:", id, "Message:", string(msg.Body))
		err = conn.Ack(msg)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return conn.Disconnect()
}
