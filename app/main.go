package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

var validTypes = map[string]string{
	"exit": "exit",
	"echo": "echo",
	"type": "type",
}

func findBinInPath(bin string) (string, bool) {
	//checks in the binary is in path or not
	paths := os.Getenv("PATH")
	for _, path := range strings.Split(paths, ":") {
		file := path + "/" + bin
		if _, err := os.Stat(file); err == nil {
			return file, true
		}
	}
	return "", false
}

func ExitCommand(argv []string) {
	code := 0
	if len(argv) > 1 {
		argcode, err := strconv.Atoi(argv[1])
		if err != nil {
			code = argcode
		}
	}

	os.Exit(code)
}

func EchoCommand(argv []string) {
	output := strings.Join(argv[1:], " ")
	fmt.Fprintf(os.Stdout, "%s\n", output)
}

func TypeCommand(argv []string) {
	if len(argv) == 1 {
		return
	}

	value := argv[1]

	if key, ok := validTypes[value]; ok {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", key)
		return
	}

	if file, exists := findBinInPath(value); exists {
		fmt.Fprintf(os.Stdout, "%s is %s\n", value, file)
		return
	}

	fmt.Fprintf(os.Stdout, "%s: not found\n", value)
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

		argv := cmds
		cmd := cmds[0]

		switch cmd {
		case "exit":
			//exit command
			// os.Exit(0)
			ExitCommand(argv)
		case "echo":
			// fmt.Printf("%s\n", strings.Join(cmds[1:], " "))
			// continue
			EchoCommand(argv)
		case "type":
			// key := cmds[1]
			// val, ok := validTypes[key]
			// if file, exists := findBinInPath(val); exists {
			// 	fmt.Printf("%s is %s\n", val, file)
			// 	continue
			// }
			// if ok {
			// 	fmt.Printf("%s is a shell builtin\n", val)
			// 	continue
			// } else {
			// 	fmt.Printf("%s", strings.Join(cmds[1:], " ")+": not found")
			// 	fmt.Println()
			// 	continue
			// }
			TypeCommand(argv)
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
		}
	}
}
