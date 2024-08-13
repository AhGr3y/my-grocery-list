package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {

	commandList := getCommands()

	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("MyGroceryList is a inventory management tool that acts as a grocery list.")
	fmt.Println()
	fmt.Println("List of commands:")
	for _, command := range commandList {
		fmt.Println()
		fmt.Printf("\t%v\n", command.name)
		fmt.Println()
		fmt.Printf("\t    %v\n", command.description)
	}
	fmt.Println("--------------------------------------------------------------------------")

	return nil
}
