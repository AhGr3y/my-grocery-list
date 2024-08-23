package main

import (
	"github.com/urfave/cli/v2"
)

func configCommands(app *cli.App, cfg *config) {
	app.Commands = []*cli.Command{
		{
			Name:  "Exit",
			Usage: "Exit the program",
			Action: func(ctx *cli.Context) error {
				return commandExit(app, ctx, cfg)
			},
		},
		{
			Name:  "Help",
			Usage: "Display program usage information",
			Action: func(ctx *cli.Context) error {
				return commandHelp(app, ctx, cfg)
			},
		},
		{
			Name:  "Store Item",
			Usage: "Add an item to your inventory",
			Action: func(ctx *cli.Context) error {
				return commandStoreItem(app, ctx, cfg)
			},
		},
	}

}
