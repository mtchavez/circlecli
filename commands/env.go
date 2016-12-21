package commands

import (
	"fmt"
	"os"
	"strings"
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

var envSubcommands = []cli.Command{
	{
		Name:   "list",
		Usage:  "all environment variables",
		Action: envListAction,
		Flags:  envFlags,
	},
	{
		Name:   "get",
		Usage:  "value of an environment variable",
		Action: envGetAction,
		Flags:  append(envFlags, envVarFlag),
	},
	{
		Name:   "set",
		Usage:  "a value for an environment variable",
		Action: envSetAction,
		Flags:  append(envFlags, envVarFlag),
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
	token := context.GlobalString("token")
	client := circleci.NewClient(token)
	padding := 1
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, padding, '\t', tabwriter.AlignRight)

	envVar, _ := client.GetEnvVar(context.String("user"), context.String("project"), context.String("var"))
	fmt.Fprintln(writer, "Name\tValue\t")
	fmt.Fprintln(writer, fmt.Sprintf("%s\t%s\t", envVar.Name, envVar.Value))
	writer.Flush()
	return nil
}

func envSetAction(context *cli.Context) error {
	varInput := strings.Split(context.String("var"), "=")
	envName := ""
	envValue := ""
	if len(varInput) == 2 {
		envName = varInput[0]
		envValue = varInput[1]
	}
	token := context.GlobalString("token")
	client := circleci.NewClient(token)
	padding := 1
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, padding, '\t', tabwriter.AlignRight)

	envVar, _ := client.AddEnvVar(context.String("user"), context.String("project"), envName, envValue)
	fmt.Fprintln(writer, "Name\tValue\t")
	fmt.Fprintln(writer, fmt.Sprintf("%s\t%s\t", envVar.Name, envVar.Value))
	writer.Flush()
	return nil
}
