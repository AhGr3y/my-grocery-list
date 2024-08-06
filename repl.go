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

		if userInput == "exit" {
			os.Exit(1)
		}
		fmt.Printf("You typed %d characters.\n", len(userInput))
	}
}
