package web

import "html/template"

type (
	Page struct {
		Title string
		Menu  []MenuItem
		Main  Main
	}

	Main struct {
		Form   Form
		Before []Widget
		After  []Widget
	}

	Form struct {
		IsForm bool
		Id     uint64
		Title  string
		Due    int
		Color  string
		Inputs []Input
		Submit string
	}

	MenuItem struct {
		Selected bool
		Display  string
		CSS      string
		URL      template.URL
	}

	Select []Option
	Option struct {
		Value    uint64
		Selected bool
		Display  string
	}
)

type (
	Input struct {
		Type InputType
		Arg  interface{}
	}

	InputType string
)

type (
	Widget struct {
		Type WidgetType
		Arg  interface{}
	}

	WidgetType string
)
