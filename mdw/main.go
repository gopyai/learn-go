package main

import (
	"fmt"
)

type (
	handler func(string)
	mdw func(handler) handler
	chain []mdw
)

func endPoint(s string) {
	fmt.Println("EndPoint:", s)
}

func mdw1(h handler) handler {
	return func(s string) {
		fmt.Println("mdw1:", s)
		h("mdw1 " + s)
	}
}

func mdw2(h handler) handler {
	return func(s string) {
		fmt.Println("mdw2:", s)
		h("mdw2 " + s)
	}
}

func main() {
	mux("Direct end point", endPoint)
	mux("MDW 1", mdw1(endPoint))
	mux("MDW 1, 2", mdw1(mdw2(endPoint)))
	mux("Chain", Chain(mdw1, mdw2).Then(endPoint))
}

func mux(name string, h handler) {
	fmt.Printf("### %s ###\n", name)
	h("Hello")
}

func Chain(c ...mdw) chain {
	return c
}

func (c chain) Then(h handler) handler {
	return func(s string) {
		i := len(c) - 1
		f := c[i](h)
		i--
		for i >= 0 {
			f = c[i](f)
			i--
		}
		f(s)
	}
}
