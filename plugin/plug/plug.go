package main

import (
	"fmt"
	"learn-go/plugin/something"
)

type (
	thing byte
)

var Thing something.Something = thing(0)

func (t thing) DoThis() { fmt.Println("Plug this") }
func (t thing) DoThat() { fmt.Println("Plug that") }
func Plug(name string)  { fmt.Printf("Hello %s, how are you?\n", name) }
