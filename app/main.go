package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage

	// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		actualCmd := cmd[:len(cmd)-1]
		if string(actualCmd[0:4]) == "exit" {
			os.Exit(0)
		} else {
			fmt.Printf("%s", cmd[:len(cmd)-1]+": command not found")
			fmt.Println()
		}
	}

}
