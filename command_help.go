package main

import (
	"errors"
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	// help command does not require additional arguments
	if len(args) != 0 {
		return errors.New("'help' command does not require additional arguments. Use 'help' instead")
	}

	commandList := getCommands()

	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("MyGroceryList is a inventory management tool that acts as a grocery list.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Printf("\t<command> [args ...]\n")
	fmt.Println()
	fmt.Println("List of commands:")
	for name, command := range commandList {
		fmt.Println()
		fmt.Printf("\t%v\n", name)
		fmt.Printf("\t    %v\n", command.description)
	}
	fmt.Println("--------------------------------------------------------------------------")

	return nil
}
