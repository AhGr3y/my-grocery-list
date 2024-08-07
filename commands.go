package main

import (
	"errors"
	"strings"
)

var (
	ErrEmptyCommand   = errors.New("empty command")
	ErrInvalidCommand = errors.New("invalid command. Use 'help' to view available commands")
)

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

func validateCommand(userInput string) ([]string, error) {
	if userInput == "" {
		return []string{}, ErrEmptyCommand
	}

	inputs := strings.Split(userInput, " ")
	command := inputs[0]

	_, commandExist := commandList[command]
	if !commandExist {
		return []string{}, ErrInvalidCommand
	}

	if len(inputs) > 1 {
		return inputs[1:], nil
	}

	return []string{}, nil
}
