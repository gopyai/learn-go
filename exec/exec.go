package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	runAndSignal()
}

func runAndSignal() {
	cmd := exec.Command("./signal")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println(cmd.Start())

	fmt.Println("Wait ...")
	go func() {
		time.Sleep(time.Second * 2)
		cmd.Process.Signal(os.Interrupt)
	}()
	fmt.Println(cmd.Wait())
}
