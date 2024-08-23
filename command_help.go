package main

import "github.com/urfave/cli/v2"

func commandHelp(app *cli.App, ctx *cli.Context, cfg *config) error {
	app.CustomAppHelpTemplate = "NAME:\n\t{{.Name}} - {{.Usage}}\n"

	return cli.ShowAppHelp(ctx)
}
