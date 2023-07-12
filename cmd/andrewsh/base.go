package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"bufio"
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

		cmd := exec.Command(args[0], args[1:]...)

		switch args[0] {
		case "exit":
			return
		case "cd":
			if len(args) < 2 {
				fmt.Println("Error: Directory argument is missing.")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("Error:", err)
			}
			return
		default:
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			return
		}
	}

	return
}
