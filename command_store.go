package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func checkBrand(brand string) bool {
	return false
}

func commandStore(cfg *config, args ...string) error {
	// Store commands takes at least 3 arguments:
	// store <item_name> <brand_name> [quantity]
	if len(args) < 3 || len(args) > 4 {
		errorMsg := fmt.Sprintf("expecting 3 or 4 arguments, got %v argument(s). Use 'help' for command usage", len(args))
		return errors.New(errorMsg)
	}

	// Quantity defaults to 1
	var item_quantity int
	if len(args) == 4 {
		quantity, err := strconv.Atoi(args[3])
		if err != nil {
			return err
		}
		item_quantity = quantity
	} else {
		item_quantity = 1
	}

	// Get user confirmation to create new brand
	item_brand := args[2]
	brandExist := checkBrand(item_brand)
	if !brandExist {
		reader := bufio.NewScanner(os.Stdin)
		fmt.Printf("Brand '%s' does not exist. Do you want to continue? [Y/n] ", item_brand)
		reader.Scan()
		userInput := reader.Text()
		fmt.Printf("You have entered: %s\n", userInput)
	}

	// Output item quantity to user
	item_name := args[1]
	fmt.Printf("Number of %s %s on hand: %d\n", item_brand, item_name, item_quantity)

	return nil
}
