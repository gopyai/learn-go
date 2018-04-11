package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"v/err"

	"github.com/kr/pretty"
)

func main() {
	runWsServer()
}

func runWsServer() {
	m := http.NewServeMux()
	m.Handle("/", http.HandlerFunc(handler))
	err.Panic(http.ListenAndServe(":80", m))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("From:", r.RemoteAddr)
	fmt.Println("URI:", r.RequestURI)
	pretty.Println("Header:", r.Header)

	b, e := ioutil.ReadAll(r.Body)
	if e != nil {
		fmt.Println("ERROR:", e)
		return
	}
	fmt.Println("Input:", string(b))
	w.Write([]byte("OK"))
}
