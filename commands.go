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
			description: "Exit the program.",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "See the guide for using this program.",
			callback:    commandHelp,
		},
		"store": {
			name:        "store name brand [quantity]",
			description: "Add item(s) to inventory. For example, store hand_soap dettol 2.",
			callback:    commandStore,
		},
	}
}

func wordsToLower(words []string) []string {
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
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

	// Inputs are case-insensitive
	inputsLower := wordsToLower(inputs)

	// Check for valid command
	command := inputsLower[0]
	_, commandExist := commandList[command]
	if !commandExist {
		return []string{}, errors.New("invalid command. Use 'help' to view available commands")
	}

	return inputsLower, nil
}
