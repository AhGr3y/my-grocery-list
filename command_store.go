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

func promptUserForYesNo() (bool, error) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Do you want to continue? [Y/n] ")
	reader.Scan()
	userInput := reader.Text()
	proceed, err := processYesNo(userInput)
	if err != nil {
		return false, err
	}

	return proceed, nil
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

	// TODO: GET INVENTORY QUANTITY

	// Get user confirmation to create new brand
	item_name := args[1]
	item_brand := args[2]
	brandExist := checkBrand(item_brand)
	if !brandExist {
		// Loop until user input valid response
		for {
			fmt.Printf("Brand '%s' does not exist. ", item_brand)
			proceed, err := promptUserForYesNo()
			if err != nil {
				continue
			}
			if !proceed {
				fmt.Printf("%s %s not stored\n", item_brand, item_name)
				return nil
			}
			break
		}
	}
	// Add item to inventory

	// Output item quantity to user
	fmt.Printf("Number of %s %s on hand: %d\n", item_brand, item_name, item_quantity)

	return nil
}
