package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"

	"devx/web"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"learn-go/gohtml/internal/gohtml"
)

const port = 8000

var tpl = gohtml.Template()

func main() {
	runServer()
}

func runServer() {
	mux := http.NewServeMux()

	publishVueJsApp(mux)
	publishApi(mux)

	// c:=cors.New(cors.Options{
	// 	AllowedOrigins: []string{},
	// 	AllowedHeaders: []string{},
	// 	AllowedMethods: []string{},
	// })
	// c.Handler(apiMux)
	http.ListenAndServe(fmt.Sprintf(":%d", port), cors.Default().Handler(mux))
}

func publishVueJsApp(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join(os.Getenv("GOPATH"), "src/devx/webapp", "index.html"))
	})
	mux.Handle("/dist/",
		http.StripPrefix("/dist/",
			http.FileServer(
				http.Dir(
					path.Join(os.Getenv("GOPATH"), "src/devx/webapp/dist")))))
}

func publishApi(mux *http.ServeMux) {
	apiMux := httprouter.New()
	defer mux.Handle("/api/", apiMux)
	apiMux.GET("/api/xxx", xxx)
	apiMux.POST("/api/test", web.APIServer(apiTest))
	apiMux.POST("/api/okr/objective", web.APIServer(apiCreateObjective))
}

func xxx(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("Hello ...")
}

// Accepted signatures:
//
// type (
// 	handler1 func() (data, ret interface{})
// 	handler2 func(in *struct{}) (data, ret interface{})
// 	handler3 func(in *struct{}, ctx *APIContext) (data, ret interface{})
// )
func apiCreateObjective(in *struct {
	Value string
	Id    uint64
}) (data, ret interface{}) {
	// Processing ...

	// Response
	return struct {
		Value string
		Id    uint64
	}{"Return:" + in.Value, in.Id + 100}, nil
}

// Accepted signatures:
//
// type (
// 	handler1 func() (data, ret interface{})
// 	handler2 func(in *struct{}) (data, ret interface{})
// 	handler3 func(in *struct{}, ctx *APIContext) (data, ret interface{})
// )
func apiTest(in *struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}) (data, ret interface{}) {
	// Processing ...
	fmt.Printf("Name: '%s', Value: '%s'\n", in.Name, in.Value)
	status := rand.Float64() > 0.5

	time.Sleep(time.Second * 2)

	// Response
	return struct {
		Status bool `json:"status"`
	}{status}, nil
}

