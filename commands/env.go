package commands

import (
	"fmt"
	"os"
	"text/tabwriter"

	circleci "github.com/mtchavez/circlecigo"
	"github.com/urfave/cli"
)

// EnvCmd - get env settings for a project
var EnvCmd = cli.Command{
	Name:        "env",
	Usage:       "get env settings for a project",
	Flags:       envFlags,
	Action:      envListAction,
	Subcommands: envSubcommands,
}

var envFlags = []cli.Flag{
	userFlag,
	projectFlag,
}

var envSubcommands = cli.Commands{
	cli.Command{
		Name:   "list",
		Action: envListAction,
		Flags:  envFlags,
	},
	cli.Command{
		Name:   "get",
		Action: envGetAction,
		Flags:  envFlags,
	},
	cli.Command{
		Name:   "set",
		Action: envSetAction,
		Flags:  envFlags,
	},
}

func envListAction(context *cli.Context) error {
	token := context.GlobalString("token")
	client := circleci.NewClient(token)
	padding := 1
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, padding, '\t', tabwriter.AlignRight)

	envVars, _ := client.EnvVars(context.String("user"), context.String("project"))
	fmt.Fprintln(writer, "Var\tValue\t")
	for _, envVar := range envVars {
		fmt.Fprintln(writer, fmt.Sprintf("%s\t%s\t", envVar.Name, envVar.Value))
	}
	writer.Flush()
	return nil
}

func envGetAction(context *cli.Context) error {
	return nil
}

func envSetAction(context *cli.Context) error {
	return nil
}
