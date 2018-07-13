package web

import "html/template"

const (
	WidgetTypeHeader         WidgetType = "header"
	WidgetTypeParagraph      WidgetType = "paragraph"
	WidgetTypeHorizontalLine WidgetType = "hr"
	WidgetTypeSpacer         WidgetType = "spacer"
	WidgetTypeTodoList       WidgetType = "todo_list"
	WidgetTypeButton         WidgetType = "button"
	WidgetTypeAjaxButton     WidgetType = "ajax_button"
	WidgetTypeAjaxText       WidgetType = "ajax_text"
	WidgetTypeAjaxMultiText  WidgetType = "ajax_multi_text"
	WidgetTypeAjaxSelect     WidgetType = "ajax_select"
)

type (
	WidgetArgHeader string
	WidgetArgParagraph string
	WidgetArgHorizontalLine struct{}
	WidgetArgSpacer struct{}

	WidgetArgTodoList []TodoInfo
	TodoInfo struct {
		Id    uint64
		Title string
		Due   int
		Color string
	}

	WidgetArgButton struct {
		Display string
		URL     template.URL
	}

	WidgetArgAjaxButton struct {
		Display    string
		Method     string
		URL        template.URL
		URLGetHTML template.URL
	}

	WidgetArgAjaxText struct {
		IsTextArea  bool
		IsGetHTML   bool
		Label       string
		Placeholder string
		URLCreate   template.URL
		URLUpdate   template.URL
		URLGetHTML  template.URL
		Saved       string // Saved value
		Id          uint64 // If Saved is not ""
	}

	WidgetArgAjaxMultiText struct {
		IsTextArea  bool
		IsGetHTML   bool
		URLCreate   template.URL
		URLUpdate   template.URL
		URLGetHTML  template.URL
		Placeholder string
		Items       []AjaxItem
	}
	AjaxItem struct {
		Saved string
		Id    uint64
	}

	WidgetArgAjaxSelect struct {
		Label      string
		URLLoad    template.URL
		URLUpdate  template.URL
		JSOptId    template.JS
		JSDataHTML template.JS
	}
)
