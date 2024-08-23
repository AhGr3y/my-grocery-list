package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func commandExit(app *cli.App, ctx *cli.Context, cfg *config) error {
	os.Exit(1)
	return nil
}
