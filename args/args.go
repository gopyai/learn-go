package main

import (
	"os"
	"fmt"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, ":", arg)
	}
}
