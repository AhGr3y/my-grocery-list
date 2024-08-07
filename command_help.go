package main

import (
	"errors"
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {

	// help command does not require additional arguments
	if len(args) != 0 {
		return errors.New("'help' command does not require additional arguments. Try 'help' instead")
	}

	intro := `
MyGroceryList is a inventory management tool that acts as a grocery list.
`

	fmt.Println(intro)
	return nil
}
