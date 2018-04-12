package main

import (
	"fmt"

	"github.com/go-stomp/stomp"
)

func Producer(host, qRequest string, f func(string, []byte) []byte) {
	mq, e := stomp.Dial("tcp", fmt.Sprintf("%s:61613", host))
	panicIf(e)
	sub, e := mq.Subscribe(qRequest, stomp.AckAuto)
	panicIf(e)
	for {
		msg, e := sub.Read()
		panicIf(e)

		corrId := msg.Header.Get("correlation-id")
		qReply := msg.Header.Get("reply-to")
		opName := msg.Header.Get("operation-name")

		panicIf(mq.Send(
			qReply,
			"application/json",
			f(opName, msg.Body), // execute handler
			stomp.SendOpt.Header("correlation-id", corrId),
		))

		// ???

	}
}
