package main

import "fmt"
import "log"
import "ioutil"

func AndrewSH() {
	fmt.Print("andrew $: ")
}

func main() {
	loadConfigs()

	for x := 0; x < 1; x-- {
		AndrewSH();
	}
}
