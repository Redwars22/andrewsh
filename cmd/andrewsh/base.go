package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"bufio"
	"runtime"
)

func readAndInterpret() {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		args := strings.Split(input, " ")

		if len(args) == 0 {
			continue
		}

		switch args[0] {
			case "exit":
				return
			case "cd":
				if len(args) < 2 {
					fmt.Println("Error: Directory argument is missing.")
					continue
				}
				
				err := changeDirectory(args[1])
				if err != nil {
					fmt.Println("Error:", err)
				}
			default:
				var cmd *exec.Cmd

				if runtime.GOOS == "windows" {
					cmdString := strings.Join(args, " ")
					cmd = exec.Command("cmd", "/c", cmdString)
				} else {
					cmd = exec.Command(args[0], args[1:]...)
				}

				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				err := cmd.Start()

				if err != nil {
					fmt.Println("Error:", err)
				}

				cmd.Wait()

				return
		}
	}

	return
}
