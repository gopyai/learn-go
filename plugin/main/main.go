package main

import (
	"plugin"
	"fmt"
	"devx/iferr"
	"learn-go/plugin/something"
)

func main() {
	fmt.Println("This is main program")

	plug, err := plugin.Open("plug.so")
	iferr.Panic(err)

	f, err := plug.Lookup("Plug")
	iferr.Panic(err)
	f.(func(string))("Arief")

	v, err := plug.Lookup("Thing")
	iferr.Panic(err)
	t := *v.(*something.Something)
	t.DoThis()
	t.DoThat()
}
