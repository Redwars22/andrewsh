package main

func getCommand() {
	var command = ""

	fmt.Scan(&command)

	if command {
		interpret(command)
	}
}

func interpret(command string) {
	
}