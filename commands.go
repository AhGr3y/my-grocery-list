package main

import "github.com/urfave/cli/v2"

type cliCommand struct {
	name        string
	description string
	callback    func(*cli.App, *cli.Context, *config) error
}

var cliCommands = map[string]cliCommand{
	"Exit": {
		name:        "Exit",
		description: "Exit the program",
		callback:    commandExit,
	},
	"Help": {
		name:        "Help",
		description: "Display program usage information",
		callback:    commandHelp,
	},
	"Store Item": {
		name:        "Store Item",
		description: "Add an item to your inventory",
		callback:    commandStoreItem,
	},
}
