package main

import (
	"regexp"
	"fmt"
)

func main() {
	s := `on:"a.name = b.id AND c.id  =   d.hello"`
	for _, col := range cols(s) {
		fmt.Println("Col:", col)
	}
}

var r = regexp.MustCompile(`([a-z.]+) *= *([a-z.]+)`)

func cols(s string) (lst []string) {
	for _, ss := range r.FindAllStringSubmatch(s, -1) {
		for _, col := range ss[1:] {
			lst = append(lst, col)
		}
	}
	return
}
