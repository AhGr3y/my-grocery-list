package main

import (
	"bufio"
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

		// Check if user entered valid command
		_, commandExist := commandList[userInput]
		if !commandExist {
			fmt.Println("Please enter a valid command.")
			continue
		}

		// Execute command
		err := commandList[userInput].callback(cfg)
		if err != nil {
			fmt.Printf("MyGroceryList > " + err.Error())
			continue
		}
	}
}
