package main

import (
	"github.com/TwiN/go-color"
	"fmt"
	"os"
)

func AndrewSH() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print(color.Ize(color.Red, "@ANDREW - "));
	fmt.Print(color.Ize(color.Blue, path));
	fmt.Print(color.Ize(color.Red, " $: "));
	readAndInterpret();
}
