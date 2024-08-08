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
			description: "Store item(s) to inventory. Replace spaces ' ' with underscores '_' for multi-word name and brand.\n\t    For example, store hand_soap dettol 2.",
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

func processYesNo(userInput string) (bool, error) {
	// Empty string defaults to yes
	if userInput == "" {
		return true, nil
	}

	// userInput is case-insensitive
	inputLower := strings.ToLower(userInput)

	// 'yes' or 'y' is acceptable
	if inputLower == "y" {
		return true, nil
	}
	if inputLower == "yes" {
		return true, nil
	}

	// 'no' or 'n' is acceptable
	if inputLower == "n" {
		return false, nil
	}
	if inputLower == "no" {
		return false, nil
	}

	return false, ErrInvalidCommand
}
