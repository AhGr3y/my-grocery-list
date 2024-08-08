package main

import (
	"errors"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("'exit' does not require additional arguments. Use 'exit' instead")
	}
	os.Exit(1)
	return nil
}
