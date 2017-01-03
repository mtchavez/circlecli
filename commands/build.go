package commands

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	circleci "github.com/mtchavez/circlecigo"
	"github.com/urfave/cli"
)

// BuildCmd - build project and interact with builds
var BuildCmd = cli.Command{
	Name:        "build",
	Usage:       "build project and interact with builds",
	Flags:       buildFlags,
	Action:      buildProjectAction,
	Subcommands: buildSubcommands,
}

var buildFlags = []cli.Flag{
	userFlag,
	projectFlag,
}

var buildSubcommands = []cli.Command{
	{
		Name:   "cancel",
		Usage:  "cancel a specific build",
		Action: buildCancelAction,
		Flags:  append(buildFlags, buildNumFlag),
	},
}

func buildProjectAction(context *cli.Context) error {
	token := context.GlobalString("token")
	client := circleci.NewClient(token)
	build, resp := client.NewBuild(context.String("user"), context.String("project"), nil)
	if resp.Success() {
		padding := 1
		writer := tabwriter.NewWriter(os.Stdout, 0, 8, padding, '\t', tabwriter.AlignRight)
		fmt.Fprintln(writer, "Build\tBranch\tQueued\tURL")
		fmt.Fprintln(writer, fmt.Sprintf("%d\t%s\t%s\t%s\t", build.BuildNum, build.Branch, build.QueuedAt, build.BuildURL))
		writer.Flush()
		return nil
	}
	failed := errors.New("Failed to build")
	return failed
}

func buildCancelAction(context *cli.Context) error {
	token := context.GlobalString("token")
	buildNum := context.Int("num")
	client := circleci.NewClient(token)
	build, resp := client.CancelBuild(context.String("user"), context.String("project"), buildNum)
	if resp.Success() {
		padding := 1
		writer := tabwriter.NewWriter(os.Stdout, 0, 8, padding, '\t', tabwriter.AlignRight)
		fmt.Fprintln(writer, "Build\tBranch\tStatus\tURL")
		fmt.Fprintln(writer, fmt.Sprintf("%d\t%s\t%s\t%s\t", build.BuildNum, build.Branch, build.Status, build.BuildURL))
		writer.Flush()
		return nil
	}
	failed := errors.New("Failed to cancel build")
	return failed
}
