// ana
package main

import (
	"net/http"
	"vos/lib"
	"vos/websvc"
)

func main() {
	svr := New()
	go websvc.RunServer(10000, func(mux *http.ServeMux) {
		svr.RegWSAna(mux)
	})
	lib.WaitKeyEnter()
}

type (
	jsInp struct {
	}
	jsOut struct {
		Msg   string
		Error string
	}
	Server struct{}
)

func New() *Server {
	return new(Server)
}

func (my *Server) Ana() (msg string, err error) {
	return "Hello Ana!", nil
}

func (my *Server) RegWSAna(mux *http.ServeMux) {
	mux.HandleFunc(func() (string, http.HandlerFunc) {
		var (
			inp jsInp
			out jsOut
		)
		return "/Ana", websvc.JSONHandler(
			func(r *http.Request) (interface{}, interface{}) {
				return &inp, &out
			},
			func() {
				out.Msg, _ = my.Ana()
			})
	}())
}
