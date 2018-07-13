package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"learn-go/gohtml/internal/web"
)

const (
	p1 = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce non leo eu est porttitor venenatis. Curabitur efficitur, ipsum sed hendrerit scelerisque, nibh dui dictum enim, a dapibus nisi urna ut augue. Vivamus et tellus mauris."
	p2 = "Pellentesque quis pellentesque lectus, convallis varius ex. Sed sit amet neque dui. Duis sit amet arcu at est ultrices malesuada in porta ex. Duis euismod finibus justo, eu suscipit orci ullamcorper in."
)

func handlerSignUp(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if err := tpl.ExecuteTemplate(w, "page_signup", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func htmlApp(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data := web.Page{
		Title: "App",
		Menu: []web.MenuItem{
			{Selected: false, Display: "TODO", CSS: "fa fa-check", URL: "/html/todo"},
			{Selected: false, Display: "Dashboard", CSS: "fa fa-tachometer-alt", URL: "/html/dashboard"},
			{Selected: false, Display: "OKR", CSS: "fa fa fa-dot-circle", URL: "/html/okr/objective"},
		},
		Main: web.Main{
			Before: []web.Widget{
				//
				// Ajax one
				//

				{web.WidgetTypeAjaxText, web.WidgetArgAjaxText{
					// TODO Experimental
					Label:       "Objective",
					Placeholder: "What is your objective for this year?",
					URLCreate:   "/test",
					Saved:       "Satu",
					Id:          1,
				}},
				{web.WidgetTypeAjaxText, web.WidgetArgAjaxText{
					// TODO Experimental
					IsTextArea:  true,
					Label:       "Multiline statements",
					Placeholder: "You may enter multiline here",
					URLCreate:   "/test",
					Saved:       "This is multiline.\nSecond line.\nThird line.\nFourth line.\nBlah blah ...",
					Id:          3,
				}},
				{web.WidgetTypeHorizontalLine, web.WidgetArgHorizontalLine{}},

				//
				// Ajax multi
				//

				{web.WidgetTypeHeader, web.WidgetArgHeader("Objectives")},
				{web.WidgetTypeAjaxMultiText, web.WidgetArgAjaxMultiText{
					URLCreate:   "/test",
					Placeholder: "What is your objective for this year?",
					Items: []web.AjaxItem{
						{Saved: "Satu", Id: 1},
						{Saved: "Dua", Id: 2},
						{Saved: ""},
					},
				}},

				{web.WidgetTypeSpacer, web.WidgetArgSpacer{}},

				{web.WidgetTypeHeader, web.WidgetArgHeader("Multi multiline statements")},
				{web.WidgetTypeAjaxMultiText, web.WidgetArgAjaxMultiText{
					IsTextArea:  true,
					URLCreate:   "/test",
					Placeholder: "You may enter multiline here",
					Items: []web.AjaxItem{
						{Saved: "Satu", Id: 1},
						{Saved: "Dua", Id: 2},
						{Saved: ""},
					},
				}},
				{web.WidgetTypeHorizontalLine, web.WidgetArgHorizontalLine{}},

				//
				// Others
				//

				{web.WidgetTypeParagraph, web.WidgetArgParagraph(p1)},
				{web.WidgetTypeParagraph, web.WidgetArgParagraph(p2)},
				{web.WidgetTypeButton, web.WidgetArgButton{
					Display: "My Button",
					URL:     "/button",
				}},
				{web.WidgetTypeTodoList, []web.TodoInfo{
					{
						Id:    1,
						Title: "This is todo item form based",
						Due:   3,
						Color: "green",
					},
					{
						Id:    2,
						Title: "This is todo item checkbox based",
						Due:   1,
						Color: "yellow",
					},
				}},
			},

			After: []web.Widget{
				{web.WidgetTypeParagraph, web.WidgetArgParagraph(p1)},
				{web.WidgetTypeParagraph, web.WidgetArgParagraph(p2)},
			},

			Form: web.Form{
				IsForm: true,
				Id:     3,
				Title:  "This is todo item form based and checkbox based",
				Due:    0,
				Color:  "red",
				Inputs: []web.Input{
					{
						Type: web.InputTypeText,
						Arg: web.InputArgText{
							Label:       "Name",
							Name:        "name",
							Value:       "",
							Placeholder: "eg: Arief",
						},
					},
					{
						Type: web.InputTypeEmail,
						Arg: web.InputArgEmail{
							Label:       "Email Address",
							Name:        "email",
							Value:       "",
							Placeholder: "eg: username@domain",
						},
					},
					{
						Type: web.InputTypeCheckbox,
						Arg: web.InputArgCheckbox{
							Name:        "todo",
							Value:       "okr",
							Description: "Define your objectives.",
							URL:         "/html/okr/objective",
						},
					},
					{
						Type: web.InputTypeCheckbox,
						Arg: web.InputArgCheckbox{
							Name:        "todo",
							Value:       "lorem",
							Description: p1,
							URL:         "/other",
						},
					},
					{
						Type: web.InputTypeCheckbox,
						Arg: web.InputArgCheckbox{
							Name:        "todo",
							Value:       "nolink",
							Description: "It has no URL/link.",
						},
					},
				},
				Submit: "Create and Done",
			},
		},
	}
	if err := tpl.ExecuteTemplate(w, "page_main", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
