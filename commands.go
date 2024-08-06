package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

var commandList = map[string]cliCommand{
	"exit": cliCommand{
		name:        "exit",
		description: "Exit the program",
		callback:    commandExit,
	},
}

func validateCommand(userInput string) error {
	return nil
}
