package main

//
// func handlerTodoList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	data := webref.Page{
// 		Title: "Todo List",
// 		Menu: []webref.MenuItem{
// 			{true, "/todos", "fa fa-check", "TODO"},
// 			{false, "/dashboard", "fa fa-tachometer-alt", "Dashboard"},
// 			{false, "/okr", "fa fa fa-dot-circle", "OKR"},
// 		},
// 		Main: webref.MainHome{
// 			After: []webref.Widget{
// 				{webref.WidgetTypeTodoList, []webref.TodoInfo{
// 					{
// 						Id:    1,
// 						Title: "This is todo item form based",
// 						Due:   3,
// 						Color: "green",
// 					},
// 					{
// 						Id:    2,
// 						Title: "This is todo item checkbox based",
// 						Due:   1,
// 						Color: "yellow",
// 					},
// 				}},
// 			},
// 		},
// 	}
// 	if err := tpl.ExecuteTemplate(w, "page_main", data); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }
//
// func handlerTodoItemForm(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	data := webref.Page{
// 		Title: "Todo Item Form",
// 		Menu: []webref.MenuItem{
// 			{true, "/todos", "fa fa-check", "TODO"},
// 			{false, "/dashboard", "fa fa-tachometer-alt", "Dashboard"},
// 			{false, "/okr", "fa fa fa-dot-circle", "OKR"},
// 		},
// 		Main: webref.MainHome{
// 			Form: webref.Form{
// 				IsForm: true,
// 				Id:     1,
// 				Title:  "This is todo item form based",
// 				Due:    3,
// 				Color:  "green",
// 				Inputs: []webref.Input{
// 					{
// 						Type: webref.InputTypeText,
// 						Arg: webref.InputArgText{
// 							Label:       "Name",
// 							Name:        "name",
// 							Value:       "",
// 							Placeholder: "eg: Arief",
// 						},
// 					},
// 					{
// 						Type: webref.InputTypeEmail,
// 						Arg: webref.InputArgEmail{
// 							Label:       "Email Address",
// 							Name:        "email",
// 							Value:       "",
// 							Placeholder: "eg: username@domain",
// 						},
// 					},
// 				},
// 				Submit: "Create",
// 			},
// 		},
// 	}
// 	if err := tpl.ExecuteTemplate(w, "page_main", data); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }
//
// func handlerTodoItemCheckBox(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	data := webref.Page{
// 		Title: "Todo Item Checkbox",
// 		Menu: []webref.MenuItem{
// 			{true, "/todos", "fa fa-check", "TODO"},
// 			{false, "/dashboard", "fa fa-tachometer-alt", "Dashboard"},
// 			{false, "/okr", "fa fa fa-dot-circle", "OKR"},
// 		},
// 		Main: webref.MainHome{
// 			Form: webref.Form{
// 				IsForm: true,
// 				Id:     2,
// 				Title:  "This is todo item checkbox based",
// 				Due:    1,
// 				Color:  "yellow",
// 				Inputs: []webref.Input{
// 					{
// 						Type: webref.InputTypeCheckbox,
// 						Arg: webref.InputArgCheckbox{
// 							Name:        "todo",
// 							Value:       "okr",
// 							Description: "Define your objectives.",
// 							URL:         "/okr",
// 						},
// 					},
// 					{
// 						Type: webref.InputTypeCheckbox,
// 						Arg: webref.InputArgCheckbox{
// 							Name:        "todo",
// 							Value:       "lorem",
// 							Description: p1,
// 							URL:         "/other",
// 						},
// 					},
// 					{
// 						Type: webref.InputTypeCheckbox,
// 						Arg: webref.InputArgCheckbox{
// 							Name:        "todo",
// 							Value:       "nolink",
// 							Description: "It has no URL/link.",
// 						},
// 					},
// 				},
// 				Submit: "Done",
// 			},
// 		},
// 	}
// 	if err := tpl.ExecuteTemplate(w, "page_main", data); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }
