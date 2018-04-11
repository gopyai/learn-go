// main
package main

import (
	"fmt"
)

type (
	data struct {
		name string
		age  int
	}
)

func (my data) Asik() {
	my.name = "Arief"
	my.age = 17
}

func (my data) Print() {
	fmt.Println(my.name, my.age)
}

func main() {
	var d data
	d.Asik()
	d.Print()
}
