package main

import (
	"fmt"
	"html/template"
	"net/http"

	"devx/iferr"

	"github.com/julienschmidt/httprouter"
	"learn-go/gohtml/internal/web"
)

func htmlAllObjectives(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("ContentType", "text/html")
	data := web.Main{
		Before: []web.Widget{
			{web.WidgetTypeHeader, web.WidgetArgHeader("Your objectives:")},
			{web.WidgetTypeAjaxMultiText, web.WidgetArgAjaxMultiText{
				IsGetHTML:   true,
				URLCreate:   "/api/okr/objective",
				URLGetHTML:  "/html/okr/objective",
				Placeholder: "What is your objective for this year?",
				Items: []web.AjaxItem{
					{Saved: "Objective 1", Id: 1},
					{Saved: "Objective 2", Id: 2},
				},
			}},
		},
	}
	if err := tpl.ExecuteTemplate(w, "main", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func htmlObjective(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := paramId(p, "id")
	iferr.Panic(err)

	w.Header().Set("ContentType", "text/html")
	data := web.Main{
		Before: []web.Widget{
			{web.WidgetTypeAjaxText, web.WidgetArgAjaxText{
				Label:       "Objective:",
				Placeholder: "What is your objective for this year?",
				URLUpdate:   "/api/okr/objective",
				Saved:       "Dummy objective",
				Id:          id,
			}},
			{web.WidgetTypeAjaxSelect, web.WidgetArgAjaxSelect{
				Label:      "Alignment with parent key result:",
				URLLoad:    template.URL(fmt.Sprintf("/html/okr/objective/%d/parent", id)),
				URLUpdate:  template.URL(fmt.Sprintf("/api/okr/objective/%d/parent", id)),
				JSOptId:    template.JS(fmt.Sprintf("{Objective: %d, Parent: optId}", id)),
				JSDataHTML: "data.HTML",
			}},
			{web.WidgetTypeHorizontalLine, web.WidgetArgHorizontalLine{}},

			{web.WidgetTypeHeader, web.WidgetArgHeader("Key Result:")},
			{web.WidgetTypeAjaxMultiText, web.WidgetArgAjaxMultiText{
				IsGetHTML:   true,
				URLCreate:   "/api/okr/keyresult",
				URLGetHTML:  "/html/okr/keyresult",
				Placeholder: "Key result to achive the objective",
				Items: []web.AjaxItem{
					{Saved: "Key result 1", Id: 1},
				},
			}},
			{web.WidgetTypeHorizontalLine, web.WidgetArgHorizontalLine{}},

			{web.WidgetTypeAjaxButton, web.WidgetArgAjaxButton{
				Display:    "Delete This Objective",
				Method:     "DELETE",
				URL:        template.URL(fmt.Sprintf("/api/okr/objective/%d", id)),
				URLGetHTML: "/html/okr/objective",
			}},
		},
	}
	if err := tpl.ExecuteTemplate(w, "main", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func htmlObjectiveParent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// id, err := paramId(p, "id")
	// iferr.Panic(err)
	w.Header().Set("ContentType", "text/html")
	w.Write([]byte(genHtmlObjParent("0")))
}

func htmlKeyResult(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := paramId(p, "id")
	iferr.Panic(err)
	oid := 101

	w.Header().Set("ContentType", "text/html")
	data := web.Main{
		Before: []web.Widget{
			{web.WidgetTypeAjaxButton, web.WidgetArgAjaxButton{
				Display:    "Delete This Key Result",
				Method:     "DELETE",
				URL:        template.URL(fmt.Sprintf("/api/okr/keyresult/%d", id)),
				URLGetHTML: template.URL(fmt.Sprintf("/html/okr/objective/%d", oid)),
			}},
		},
	}
	if err := tpl.ExecuteTemplate(w, "main", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
