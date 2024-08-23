package main

import (
	"github.com/urfave/cli/v2"
)

var helpNameTemplate = `{{.Name}} - {{.Usage}}{{end}}`

func commandHelp(app *cli.App, ctx *cli.Context, cfg *config) error {

	// Configure text/template for help
	app.CustomAppHelpTemplate = `NAME:
	{{template "helpNameTemplate" .}}

COMMANDS:


`

	return cli.ShowAppHelp(ctx)
}
