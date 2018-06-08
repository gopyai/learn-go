package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	fullPath := path.Join(os.Getenv("GOPATH"), "src/devx/main/html/*.gohtml")
	fmt.Println(fullPath)
}
