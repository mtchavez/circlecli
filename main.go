package main

import (
	"os"

	"github.com/mtchavez/circlecli/commands"
	"github.com/urfave/cli"
)

const (
	// AppVersion is the version of the CLI
	AppVersion = "0.1.0"
	// AppName is the name of the CLI
	AppName = "circlecli"
)

func main() {
	app := cli.NewApp()
	app.Name = AppName
	app.Usage = "Interact with the CircleCI REST API"
	app.Version = AppVersion
	app.EnableBashCompletion = true
	app.Commands = commands.AllCommands
	app.Run(os.Args)
}
