package main

import (
	"github.com/urfave/cli/v2"
)

func commandHelp(app *cli.App, ctx *cli.Context, cfg *config) error {
	// Configure text/template for help
	app.CustomAppHelpTemplate = `NAME:
	{{.Name}} - {{.Usage}}

COMMANDS:
{{range .Commands}}
	{{.Name}}:
		{{.Usage}}
{{end}}
`

	return cli.ShowAppHelp(ctx)
}
