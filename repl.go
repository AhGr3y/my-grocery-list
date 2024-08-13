package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func startRepl(cfg *config) {
	// Reader for extracting user input
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to MyGroceryList!")
	fmt.Println("Use 'help' to see the list of available commands.")

	// Loop until user enters 'exit' command
	for {
		fmt.Print("MyGroceryList > ")

		// Get user input
		reader.Scan()
		userInput := reader.Text()

		// Check if user entered valid command,
		// returns back the user arguments.
		command, err := validateCommand(userInput)
		if err != nil {
			if errors.Is(err, ErrEmptyCommand) {
				continue
			}
			fmt.Printf("Error: %s.\n", err.Error())
			continue
		}

		// Execute command
		commandList := getCommands()
		err = commandList[command].callback(cfg)
		if err != nil {
			fmt.Printf("Error: %v.\n", err.Error())
			continue
		}
	}
}
