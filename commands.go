package main

import (
	"errors"
	"strings"
)

var (
	ErrEmptyCommand   = errors.New("empty command")
	ErrInvalidCommand = errors.New("invalid command")
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the program.",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "See the list of available commands.",
			callback:    commandHelp,
		},
		"store": {
			name:        "store",
			description: "Store item(s) to inventory.",
			callback:    commandStore,
		},
	}
}

func validateCommand(userInput string) (string, error) {
	// Ignore empty command
	if userInput == "" {
		return "", ErrEmptyCommand
	}

	// Only accept one argument
	userInputSplit := strings.Split(userInput, " ")
	if len(userInputSplit) != 1 {
		return "", errors.New("too many arguments. Use 'help' to view available commands")
	}

	// Commands are case-insensitive
	command := strings.ToLower(userInputSplit[0])

	// List of valid commands
	commandList := getCommands()

	// Check for valid command
	_, commandExist := commandList[command]
	if !commandExist {
		return "", errors.New("invalid command. Use 'help' to view available commands")
	}

	return command, nil
}
