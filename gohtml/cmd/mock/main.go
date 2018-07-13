package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"learn-go/gohtml/internal/gohtml"
)

const port = 8000

var tpl = gohtml.Template()

func main() {
	runServer()
}

func runServer() {
	mux := httprouter.New()
	mux.ServeFiles("/res/*filepath", http.Dir(path.Join(os.Getenv("GOPATH"), "src/learn-go/gohtml/res/")))

	mux.GET("/signup", handlerSignUp)
	mux.GET("/", htmlApp)

	mux.GET("/html/okr/objective", htmlAllObjectives)
	mux.GET("/html/okr/objective/:id", htmlObjective)
	mux.GET("/html/okr/objective/:id/parent", htmlObjectiveParent)
	mux.GET("/html/okr/keyresult/:id", htmlKeyResult)

	mux.POST("/api/okr/objective", apiCreateObjective)
	mux.PUT("/api/okr/objective/:id", apiUpdateObjective)
	mux.DELETE("/api/okr/objective/:id", apiDeleteObjective)
	mux.PUT("/api/okr/objective/:id/parent/:pid", apiUpdateObjectiveParent)
	mux.DELETE("/api/okr/keyresult/:id", apiDeleteKeyResult)

	// mux.GET("/todos", handlerTodoList)
	// mux.GET("/todos/1", handlerTodoItemForm)
	// mux.GET("/todos/2", handlerTodoItemCheckBox)

	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func paramId(p httprouter.Params, name string) (uint64, error) {
	id, err := strconv.ParseInt(p.ByName(name), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

func genHtmlObjParent(sel string) string {
	type option struct {
		val      string
		disabled bool
		text     string
	}
	m := []option{
		{"0", false, "No alignment"},
		{"", true, "---"},
		{"101", false, "Opt 1"},
		{"102", false, "Opt 2"},
		{"1", false, "Opt 3"},
		{"2", false, "Opt 4"},
	}

	html := ""
	for _, o := range m {
		s := ""
		if o.disabled {
			s = " disabled"
		} else if o.val == sel {
			s = " selected"
		}
		html += fmt.Sprintf("<option value=\"%s\"%s>%s</option>\n", o.val, s, o.text)
	}
	return html
}
