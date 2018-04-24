package main

import (
	"encoding/json"
	"fmt"
)

type (
	book struct {
		Title string
		Pages int
	}

	game struct {
		Name         string
		PlayDuration int
	}

	collection struct {
		Type string
		Item json.RawMessage
	}
)

func ExampleRaw() {
	js := []byte("{\n  \"Type\": \"book\",\n  \"Item\": {\n    \"Title\": \"Barney and Friends\",\n    \"Pages\": 10\n  }\n}")

	var c collection
	err := json.Unmarshal(js, &c)
	panicIf(err)

	switch c.Type {
	case "book":
		var b book
		panicIf(json.Unmarshal(c.Item, &b))
		fmt.Println(b)
	case "game":
		var g game
		panicIf(json.Unmarshal(c.Item, &g))
		fmt.Println(g)
	default:
		fmt.Println("Unknown type:", c.Type)
		fmt.Println(string(c.Item))
	}

	// Output:
	// {Barney and Friends 10}
}
