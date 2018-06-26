package main

import (
	"html/template"
	"devx/iferr"
	"os"
)

func main() {
	tpl, err := template.New("").Parse(`<div onclick="loadMain(this, {{.URL}} )"></div>`)
	iferr.Panic(err)
	tpl.Execute(os.Stdout, map[string]interface{}{
		"URL": template.URL("/todo/hai"),
	})
}
