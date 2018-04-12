package main

import (
	"fmt"
	"reflect"

	"arf/v1/cert" // TODO

)

func main() {
	caPriv, caCer, err := cert.GenCACert("self.signed", "Self Signed", "ID", 1, "ca.key", "ca.cer")
	panicIf(err)
	_, svrCer, err := cert.GenCert("localhost", "local.host", "Local Host", "ID", 10, caPriv, caCer, "server.key", "server.cer")
	panicIf(err)
	panicIf(svrCer.CheckSignatureFrom(caCer))

	go server()
	client(caCer)
}

func panicIf(err error) {
	if err != nil {
		fmt.Println(reflect.TypeOf(err))
		panic(err)
	}
}
