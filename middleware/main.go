package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/justinas/alice"
)

const (
	port = 8080
)

func m1(next http.Handler) http.Handler {
	log.Println("m1")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func m2(next http.Handler) http.Handler {
	log.Println("m2")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	fiH := http.HandlerFunc(final)

	mux := http.NewServeMux()
	mux.Handle("/", fiH)

	c := alice.New(m1, m2).Then(mux)
	http.ListenAndServe(fmt.Sprintf(":%d", port), c)
}
