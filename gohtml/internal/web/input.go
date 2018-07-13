package web

import "html/template"

const (
	InputTypeText     InputType = "text"
	InputTypeEmail    InputType = "email"
	InputTypeCheckbox InputType = "checkbox"
)

type (
	InputArgText struct {
		Label       string
		Name        string
		Value       string
		Placeholder string
	}

	InputArgEmail struct {
		Label       string
		Name        string
		Value       string
		Placeholder string
	}

	InputArgTextArea struct {
		Label string
		Name  string
		Value string
	}

	InputArgCheckbox struct {
		Name        string
		Value       string
		Description string
		URL         template.URL
	}
)
