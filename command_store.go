package main

import (
	"errors"
	"fmt"
)

func commandStore(cfg *config, args ...string) error {
	// Store commands takes at least 3 arguments:
	// store <item_name> <brand_name> [quantity]
	if len(args) < 4 {
		errorMsg := fmt.Sprintf("expecting 3 or 4 arguments, got %v argument(s)\n", len(args))
		return errors.New(errorMsg)
	}
	return nil
}
