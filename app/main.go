package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	runShell := true
	for runShell {
		fmt.Print("$ ")
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}

		command := strings.Fields(cmd)
		cmd = command[0]
		args := command[1:]

		if cmd == "exit" {
			runShell = false
			break
		}

		if cmd == "echo" {
			for _, v := range args {
				fmt.Printf("%v ", v)
			}
			fmt.Printf("\n")
		} else {
			fmt.Println(cmd + ": command not found")
		}
	}
}
