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

	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(string(output))
	}
}
