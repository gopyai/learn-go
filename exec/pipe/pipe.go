package main

import (
	"os/exec"
	"bitbucket.org/stefarf/iferr"
	"bufio"
	"fmt"
	"io"
)

func main() {
	cmd := exec.Command("mysqldump", "-u", "root", "-proot", "itsp")
	r, err := cmd.StdoutPipe()
	iferr.Panic(err)

	go func() {
		br := bufio.NewReader(r)
	loop:
		for {
			s, err := br.ReadString('\n')
			switch err {
			case nil:
				fmt.Print(s)
			case io.EOF:
				break loop
			default:
				iferr.Panic(err)
			}
		}
	}()

	err = cmd.Run()
	iferr.Panic(err)
}
