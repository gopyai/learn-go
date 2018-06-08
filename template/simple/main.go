package main

import (
	"os"
	"text/template"
)

const (
	x = `Hello {{.Name}}, your age is {{.Age}}.`
)

var (
	tx = template.Must(template.New("x").Parse(x))
)

func main() {
	d := struct {
		Name string
		Age  int
	}{"Arief", 17}
	panicIf(tx.Execute(os.Stdout, d))
}

func panicIf(e error) {
	if e != nil {
		panic(e)
	}
}

