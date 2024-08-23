package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func processMainMenuInput(userInput string, app *cli.App, ctx *cli.Context) error {
	for _, command := range app.Commands {
		if userInput == command.Name {
			err := command.Action(ctx)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("command doesn't exist")
}

func runApp(cfg *config) {
	app := cli.NewApp()
	app.Name = "MyGroceryList"
	app.Usage = "A grocery list and stocktaking command-line program."

	// Configure app commands
	configCommands(app, cfg)

	mainMenuPrompt := promptui.Select{
		Label: "Press <Enter> to select one of these options",
		Items: []string{
			"Store Item", "Help", "Exit",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		fmt.Println("-------------------------")
		fmt.Println("Welcome to MyGroceryList!")
		_, result, err := mainMenuPrompt.Run()
		if err != nil {
			log.Fatalf("Error running prompt: %s", err)
		}
		err = processMainMenuInput(result, app, ctx)
		if err != nil {
			return err
		}
		return nil
	}

	// Loop until user use Exit command
	for {
		if err := app.Run(os.Args); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}

}
