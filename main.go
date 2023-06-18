package main

import "fmt"
import "log"
import "ioutil"

func loadConfigs() {
	config, err := ioutil.ReadFile("andrewsh.dat");

    if err != nill {
        log.Fatalf("unable to load settings. Does andrewsh.dat exist?")
    }

    fmt.print(string(config))
}

func AndrewSH() {
	fmt.Print("andrew $: ")
}

func main() {
	loadConfigs()

	for x := 0; x < 1; x-- {
		AndrewSH();
	}
}
