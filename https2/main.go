package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	//	genCer()

	go svr()
	time.Sleep(time.Second)
	cli()
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Miaaaw ...\n")
}

func svr() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":10000", "host.cer", "host.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func cli() {
	t := new(tls.Config)
	pool := x509.NewCertPool()
	b, err := ioutil.ReadFile("ca.cer")
	if err != nil {
		log.Fatalf("Can not read file ca.cer: %v", err)
	}
	if !pool.AppendCertsFromPEM(b) {
		log.Fatalf("Can not load ca.cer into pool")
	}
	t.RootCAs = pool

	//

	req, e := http.NewRequest("POST", "https://localhost:10000/hello", bytes.NewReader([]byte{1, 2, 3}))
	panicIf(e)

	cli := &http.Client{Transport: &http.Transport{
		TLSClientConfig:    t,
		DisableCompression: true,
	}}
	res, e := cli.Do(req)
	panicIf(e)
	out, e := ioutil.ReadAll(res.Body)
	panicIf(e)
	log.Println(string(out))
}

func panicIf(e error) {
	if e != nil {
		panic(e)
	}
}
