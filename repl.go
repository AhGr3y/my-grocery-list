package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to MyGroceryList!")
	for {
		fmt.Print("MyGroceryList > ")
		reader.Scan()
		userInput := reader.Text()

		_, commandExist := commandList[userInput]
		if !commandExist {
			fmt.Println("Please enter a valid command.")
			continue
		}
	}
}
