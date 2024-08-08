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

	// Loop until user enters 'exit' command
	for {
		fmt.Print("MyGroceryList > ")
		reader.Scan()
		userInput := reader.Text()

		// Check if user entered valid command,
		// returns any subsequent arguments and
		// the list of valid commands.
		args, commandList, err := validateCommand(userInput)
		if err != nil {
			if errors.Is(err, ErrEmptyCommand) {
				continue
			} else {
				fmt.Printf("Error: %s.\n", err.Error())
				continue
			}
		}

		// Execute command
		err = commandList[userInput].callback(cfg, args...)
		if err != nil {
			fmt.Printf("MyGroceryList > " + err.Error())
			continue
		}
	}
}
