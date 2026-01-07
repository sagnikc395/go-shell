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

		switch cmd {
		case "exit":
			runShell = false
			break
		case "echo":
			for _, v := range args {
				fmt.Printf("%v ", v)
			}
			fmt.Printf("\n")
			break
		case "type":
			type_check := args[0]
			if type_check == "echo" || type_check == "exit" || type_check == "type" {
				fmt.Printf("%v is a shell builtin\n", type_check)
			} else {
				fmt.Printf("%v: not found\n", type_check)
			}

		default:
			fmt.Printf(cmd + ": not found")
		}

	}
}
