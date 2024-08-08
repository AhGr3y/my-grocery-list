package main

import (
	"errors"
	"strings"
)

var (
	ErrEmptyCommand = errors.New("empty command")
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
}

func validateCommand(userInput string) ([]string, error) {
	// Ignore empty command
	if userInput == "" {
		return []string{}, ErrEmptyCommand
	}

	// List of valid commands
	commandList := getCommands()

	// Split command from arguments
	inputs := strings.Split(userInput, " ")
	command := inputs[0]

	// Command is case-insensitive
	commandLower := strings.ToLower(command)

	// Check for valid command
	_, commandExist := commandList[commandLower]
	if !commandExist {
		return []string{}, errors.New("invalid command. Use 'help' to view available commands")
	}

	// Return arguments if any
	if len(inputs) > 1 {
		return inputs[1:], nil
	}

	return []string{}, nil
}
