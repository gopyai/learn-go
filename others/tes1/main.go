package main

import (
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"vos/onerror"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Miaaaw ...\n")
}

func main() {
	go svr()
	time.Sleep(time.Second)
	cli()
}

func svr() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":10000", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func cli() {
	req, e := http.NewRequest("POST", "https://localhost:10000/hello", bytes.NewReader([]byte{1, 2, 3}))
	onerror.Panic(e)

	cli := &http.Client{Transport: &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}}
	res, e := cli.Do(req)
	onerror.Panic(e)
	out, e := ioutil.ReadAll(res.Body)
	onerror.Panic(e)
	log.Println(string(out))
}
