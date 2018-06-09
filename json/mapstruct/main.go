package main

import (
	"encoding/json"
	"devx/iferr"
	"fmt"
	"github.com/kr/pretty"
	"github.com/mitchellh/mapstructure"
)

type (
	collection struct {
		Type string
		Item interface{}
	}

	book struct {
		Title string
		Pages int
	}

	game struct {
		Name         string
		PlayDuration int
	}
)

func main() {
	data := collection{
		Type: "book",
		Item: book{
			Title: "Golang",
			Pages: 100,
		},
	}
	b, err := json.Marshal(&data)
	iferr.Panic(err)
	fmt.Println(string(b))

	// Unmarshal

	var c collection
	iferr.Panic(json.Unmarshal(b, &c))

	pretty.Println(c)

	tbl := map[string]func(item interface{}){
		"book": func(item interface{}) {
			var bb book
			iferr.Panic(mapstructure.Decode(c.Item, &bb))
			pretty.Println(&bb)
		},
		"game": func(item interface{}) {
			var gg game
			iferr.Panic(mapstructure.Decode(c.Item, &gg))
			pretty.Println(&gg)
		},
	}
	tbl[c.Type](c.Item)

}
