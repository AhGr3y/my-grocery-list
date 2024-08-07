package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

var commandList = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the program",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "See the guide for using this program",
		callback:    commandHelp,
	},
}

func validateCommand(userInput string) error {
	return nil
}
