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

func validateCommand(userInput string) ([]string, map[string]cliCommand, error) {
	// Ignore empty command
	if userInput == "" {
		return []string{}, map[string]cliCommand{}, ErrEmptyCommand
	}

	// List of valid commands
	commandList := map[string]cliCommand{
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

	// Split command from arguments
	inputs := strings.Split(userInput, " ")
	command := inputs[0]

	// Command is case-insensitive
	commandLower := strings.ToLower(command)

	// Check for valid command
	_, commandExist := commandList[commandLower]
	if !commandExist {
		return []string{}, map[string]cliCommand{}, errors.New("invalid command. Use 'help' to view available commands")
	}

	// Return arguments if any
	if len(inputs) > 1 {
		return inputs[1:], commandList, nil
	}

	return []string{}, commandList, nil
}
