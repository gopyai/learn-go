package main

import (
	"html/template"
	"os"
)

func main() {
	t := template.Must(template.ParseGlob("templates/*.gohtml"))
	t.ExecuteTemplate(os.Stdout, "main.gohtml", struct {
		Title string
		Name  string
		Age   int
	}{"Template Structure", "Arief", 17})
}
