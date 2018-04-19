package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

const (
	// mysql    = "c:\\Program Files\\MySQL\\MySQL Server 5.7\\bin\\mysql.exe"
	rootPass = "root"
)

func main() {
	f, err := os.Open("create.sql")
	panicIf(err)
	defer f.Close()

	cmd := exec.Command("mysql", "-u", "root", "-p"+rootPass)
	cmd.Stdin = f

	var out bytes.Buffer
	cmd.Stdout = &out

	panicIf(cmd.Run())

	fmt.Printf("Output: ###\n%s\n###\n", out.String())
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
