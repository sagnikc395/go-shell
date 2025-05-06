package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

var validTypes = map[string]string{
	"exit": "exit",
	"echo": "echo",
	"type": "type",
}

func main() {
	// Uncomment this block to pass the first stage

	// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		commandstr, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		cmds := strings.Split(commandstr[:len(commandstr)-1], " ")
		switch cmds[0] {
		case "exit":
			//exit command
			os.Exit(0)
		case "echo":
			fmt.Printf("%s\n", strings.Join(cmds[1:], " "))
			continue
		case "type":
			key := cmds[1]
			val, ok := validTypes[key]
			if ok {
				fmt.Printf("%s is a shell builtin\n", val)
				continue
			} else {
				fmt.Printf("%s", strings.Join(cmds[1:], " ")+": not found")
				fmt.Println()
				continue
			}
		}
		fmt.Printf("%s", commandstr[:len(commandstr)-1]+": command not found")
		fmt.Println()

	}

}
