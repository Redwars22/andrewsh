package main

import (
	"os/exec"
	"fmt"
)

func main(){
	exec.Command("clear")
	fmt.Println("Welcome to AndrewSH v0.1.3...");
	loadConfigs();
}