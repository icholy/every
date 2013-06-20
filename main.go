package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {

	// make sure we have enough arguments
	if len(os.Args) < 3 {
		fmt.Println("usage: every <interval> <command>")
		return
	}

	// parse the parameters
	interval, err := time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Println("error: invalid interval format")
	}
	command := os.Args[2]
	arguments := []string{}
	if len(os.Args) > 3 {
		arguments = os.Args[3:]
	}

	// start looping
	for {
		cmd := exec.Command(command, arguments...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		_ = cmd.Run()
		time.Sleep(interval)
	}
}
