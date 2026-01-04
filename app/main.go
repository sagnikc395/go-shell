package main

import (
	"bufio"
	"fmt"
	"os"
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

		if cmd[:len(cmd)-1] == "exit" {
			runShell = false
			break
		}

		fmt.Println(cmd[:len(cmd)-1] + ": command not found")

	}
}
