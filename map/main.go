package main

import "fmt"

func main() {
	var m1 map[string]bool
	var m2 map[string]*string

	fmt.Println(m1["hello"])
	fmt.Println(m2["hello"])
}
