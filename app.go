package main

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func processMainMenuInput(userInput string, ctx *cli.Context, cfg *config) error {
	switch userInput {
	case "Exit":
		os.Exit(1)
	case "Help":
		return cli.ShowAppHelp(ctx)
	case "Store Item":
		return processStoreCommand(cfg)
	default:
		return nil
	}
	return nil
}

func runApp(cfg *config) {
	app := cli.NewApp()
	app.Name = "MyGroceryList"
	app.Usage = "A grocery list and stocktaking command-line program."

	app.CustomAppHelpTemplate = "NAME:\n\t{{.Name}} - {{.Usage}}\n"

	/*
		// Set commands
		app.Commands = []*cli.Command{
			{
				Name:    "punch",
				Aliases: []string{"p"},
				Usage:   "Options for punching",
				Subcommands: []*cli.Command{
					{
						Name:  "random",
						Usage: "Punch a random person",
						Action: func(ctx *cli.Context) error {
							fmt.Println("You punched a random person!")
							return nil
						},
					},
				},
			},
		}
	*/

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
		err = processMainMenuInput(result, ctx, cfg)
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
