package gohtml

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path"

	"devx/iferr"
	"learn-go/gohtml/internal/web"
)

const msgErrTemplateParse = "Template: Error parsing '*.gohtml' files"

var id uint

func Template() *template.Template {

	// Initialize templates
	tplWidget, err := template.
		New("").Funcs(template.FuncMap{
		"goId": goId,
	}).ParseGlob(path.Join(os.Getenv("GOPATH"), "src/learn-go/gohtml/template/widgets/*.gohtml"))
	iferr.Exit(err, msgErrTemplateParse)
	tplInput, err := template.
		New("").
		ParseGlob(path.Join(os.Getenv("GOPATH"), "src/learn-go/gohtml/template/inputs/*.gohtml"))
	iferr.Exit(err, msgErrTemplateParse)

	tpl, err := template.
		New("").Funcs(template.FuncMap{
		"parseWidget": parseWidget(tplWidget),
		"parseInput":  parseInput(tplInput),
		"goId":        goId,
	}).ParseGlob(path.Join(os.Getenv("GOPATH"), "src/learn-go/gohtml/template/pages/*.gohtml"))
	iferr.Exit(err, msgErrTemplateParse)

	return tpl
}

func goId() string {
	id++
	return fmt.Sprintf("go%d", id)
}

func parseWidget(tpl *template.Template) func(w web.Widget) template.HTML {
	return func(w web.Widget) template.HTML {
		var buf bytes.Buffer
		if err := tpl.ExecuteTemplate(&buf, string(w.Type), w.Arg); err != nil {
			return template.HTML(err.Error())
		}
		return template.HTML(buf.String())
	}
}

func parseInput(tpl *template.Template) func(i web.Input) template.HTML {
	return func(i web.Input) template.HTML {
		var buf bytes.Buffer

		var tplName string
		var inputArg interface{}
		switch i.Type {

		case web.InputTypeText:
			tplName = "input"
			arg := i.Arg.(web.InputArgText)
			inputArg = map[string]interface{}{
				"Type":        "text",
				"Label":       arg.Label,
				"Name":        arg.Name,
				"Value":       arg.Value,
				"Placeholder": arg.Placeholder,
			}

		case web.InputTypeEmail:
			tplName = "input"
			arg := i.Arg.(web.InputArgEmail)
			inputArg = map[string]interface{}{
				"Type":        "email",
				"Label":       arg.Label,
				"Name":        arg.Name,
				"Value":       arg.Value,
				"Placeholder": arg.Placeholder,
			}

		default:
			tplName = string(i.Type)
			inputArg = i.Arg
		}

		if err := tpl.ExecuteTemplate(&buf, tplName, inputArg); err != nil {
			return template.HTML(err.Error())
		}
		return template.HTML(buf.String())
	}
}
