package main

import (
	"html/template"
	"os"
	"fmt"
)

type (
	Page struct {
		Title   string
		Menu    []string
		Content Content
	}

	Content struct {
		Before []Widget
		After  []Widget
		Form   Form
	}

	Form struct {
		Render bool
		Inputs []Input
		Submit string
	}

	Input struct {
	}

	Widget struct {
		Type string
		Arg  interface{}
	}
)

var (
	tpl = template.Must(template.
		New("").
		Funcs(template.FuncMap{"parseWidget": parseWidget}).
		ParseGlob("templates/*.gohtml"))

	widgetParsers = map[string]func(arg interface{}) string{
		"card": widgetCard,
		"age":  widgetAge,
	}
)

func main() {
	data := Page{
		Title: "Judul asik",
		Menu:  []string{"Satu", "Dua"},
		Content: Content{
			After: []Widget{
				{Type: "card", Arg: "queen heart"},
				{Type: "age", Arg: 17},
				{Type: "no", Arg: "processor"},
			},
			Form: Form{Render: true, Submit: "Submit"},
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "Page", data)
	if err != nil {
		panic(err)
	}
}

func parseWidget(w Widget) string {
	if parser, ok := widgetParsers[w.Type]; ok {
		return parser(w.Arg)
	} else {
		return fmt.Sprintf("No processor: %v", w)
	}
}

func widgetCard(arg interface{}) string {
	return fmt.Sprintf("Card name: %s", arg.(string))
}

func widgetAge(arg interface{}) string {
	return fmt.Sprintf("Age: %d years old", arg.(int))
}
