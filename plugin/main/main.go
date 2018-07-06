package main

import (
	"plugin"
	"fmt"
)

func main() {
	fmt.Println("This is main program")

	p, err := plugin.Open("plug.so")
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("Plug")
	if err != nil {
		panic(err)
	}

	f.(func(string))("Arief")
}
