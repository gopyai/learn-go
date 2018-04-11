package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var (
	pidFile = flag.String("p", "my.pid", "PID file")
	stop    = flag.Bool("stop", false, "Stop this process")
)

var (
	config struct {
		PID int
	}
)

func init() { flag.Parse() }

func checkAndKillRunningProcess() {
	// Check previous running process
	b, e := ioutil.ReadFile(*pidFile)
	if e != nil {
		return
	}
	e = os.Remove(*pidFile)
	if e != nil {
		panic(e)
	}
	e = json.Unmarshal(b, &config)
	if e != nil {
		return
	}

	// Kill running process
	p, e := os.FindProcess(config.PID)
	if e != nil {
		panic(e)
	}
	e = p.Kill()
	if e != nil {
		fmt.Println("Failed to kill PID", config.PID)
	}
}

func savepid() {
	config.PID = os.Getpid()
	b, e := json.Marshal(&config)
	if e != nil {
		panic(e)
	}
	e = ioutil.WriteFile(*pidFile, b, 0600)
	if e != nil {
		panic(e)
	}
}

func main() {
	checkAndKillRunningProcess()
	if *stop {
		return
	}
	savepid()

	for {
		fmt.Println("PID", config.PID, "is running")
		time.Sleep(time.Second)
	}

}
