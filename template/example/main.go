package main

import (
	"html/template"
	"os"
	"path"
)

func main() {
	tpl := template.Must(template.ParseFiles(path.Join(os.Getenv("GOPATH"), "src/learn-go/template/example/template.gohtml")))
	tpl.Execute(os.Stdout, map[string]interface{}{
		"Name":  "Arief",
		"Age":   17,
		"Quote": "\"Hi! How are you?\", I said.",
	})
}
