// main
package main

import (
	"fmt"
)

type (
	Data struct {
		Setter
	}

	Setter struct{ name string }
)

func (my *Setter) Set(name string) { my.name = name }
func (my *Setter) Print()          { fmt.Println(my.name) }
func (my *Data) Print() {
	fmt.Println("Data")
	my.Setter.Print()
}

func main() {
	var d Data
	d.Set("Arief")
	d.Print()
}
