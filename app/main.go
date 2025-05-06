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
		}
		fmt.Printf("%s", commandstr[:len(commandstr)-1]+": command not found")
		fmt.Println()

	}

}
