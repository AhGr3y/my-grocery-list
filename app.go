package main

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func runApp(cfg *config) {
	app := cli.NewApp()
	app.Name = "MyGroceryList"
	app.Usage = "Your very own household inventory grocery list."

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

	mainMenuPrompt := promptui.Select{
		Label: "Press <Enter> to select one of these options",
		Items: []string{
			"Store Item", "Help", "Exit",
		},
		/*
			Templates: &promptui.SelectTemplates{
				Help: "Hello",
			},
		*/
	}

	app.Action = func(ctx *cli.Context) error {
		fmt.Println("Welcome to MyGroceryList!")
		_, result, err := mainMenuPrompt.Run()
		if err != nil {
			log.Fatalf("Error running prompt: %s", err)
		}
		err = processMainMenuInput(result, ctx)
		if err != nil {
			return err
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func processMainMenuInput(userInput string, ctx *cli.Context) error {

	switch userInput {
	case "Exit":
		os.Exit(1)
	case "Help":
		cli.ShowAppHelp(ctx)
	default:
		return nil
	}
	return nil
}
