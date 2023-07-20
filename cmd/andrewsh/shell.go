package main

import (
	"github.com/TwiN/go-color"
	"fmt"
	"os"
	"os/user"
)

func AndrewSH() {
	path, err := os.Getwd()
	user, usErr := user.Current()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if usErr != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print(color.Ize(color.Red, "@" + user.Username));
	fmt.Print(color.Ize(color.Purple, " ANDRWSH "));
	fmt.Print(color.Ize(color.Blue, "~" + path));
	fmt.Print("\n$ ");
	readAndInterpret();
}
