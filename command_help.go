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

	intro := `--------------------------------------------------------------------------
MyGroceryList is a inventory management tool that acts as a grocery list.

Usage:

	<command> [args ...]

List of commands:

	help
		See help
	exit
		Exit this program.
--------------------------------------------------------------------------
`

	fmt.Println(intro)
	return nil
}
