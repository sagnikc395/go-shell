package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
		case "echo":
			for _, v := range args {
				fmt.Printf("%v ", v)
			}
			fmt.Printf("\n")
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Printf("error: %v\n", err)
			} else {
				fmt.Printf("%s\n", dir)
			}

		case "type":
			if len(args) == 0 {
				continue
			}

			typeCheck := args[0]

			builtins := map[string]bool{"echo": true, "exit": true, "type": true, "pwd": true}
			if builtins[typeCheck] {
				fmt.Printf("%s is a shell builtin\n", typeCheck)
				continue
			}

			// check it in PATH
			// 1.Go through every directory in PATH. For each directory
			// 2. check if a file with the command name exists
			// 3. check if the file has execute permissions
			// 4. if the file exists and has execute permissions, print <command> is <full_path> and stop.
			// 5. If the file exists but lacks execute permissions, skip it and continue to the next directory.
			pathEnv := os.Getenv("PATH")
			paths := strings.Split(pathEnv, string(os.PathListSeparator))
			found := false

			for _, dir := range paths {
				fullPath := filepath.Join(dir, typeCheck)

				// check if file exists and if we can access it
				info, err := os.Stat(fullPath)
				if err != nil {
					continue // file doesnt exist condn
				}

				// check if its is a regular file and has executbale permissions
				// 0111 repr the executable bit
				if !info.IsDir() && info.Mode()&0o111 != 0 {
					fmt.Printf("%s is %s\n", args[0], fullPath)
					found = true
					break
				}
			}

			if !found {
				fmt.Printf("%v: not found\n", typeCheck)
			}

		default:
			// search for an executable with the given name in the directories listed in PATH
			// if found, execute the program
			// pass any arguments from the cli to the program

			// prepare the external program
			// cmd is the executable ame, args is the slice of arguments
			externalCmd := exec.Command(cmd, args...)

			externalCmd.Stdout = os.Stdout
			externalCmd.Stderr = os.Stderr
			externalCmd.Stdin = os.Stdin

			// running the command and checking for errors
			err := externalCmd.Run()
			if err != nil {
				// if error is "executable file not found", print the command not found stmt
				if _, ok := err.(*exec.Error); ok {
					fmt.Printf("%s: command not found\n", cmd)
				} else {
					fmt.Printf("%s: %v\n", cmd, err)
				}
			}

		}

	}
}
