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
	Flags:       append(buildFlags, branchFlag),
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
	{
		Name:   "retry",
		Usage:  "retry a specific build",
		Action: buildRetryAction,
		Flags:  append(buildFlags, buildNumFlag),
	},
}

func buildProjectAction(context *cli.Context) error {
	token := context.GlobalString("token")
	client := circleci.NewClient(token)
	user := context.String("user")
	project := context.String("project")
	branch := context.String("branch")
	var build *circleci.Build
	var resp *circleci.APIResponse
	if branch != "" {
		build, resp = client.BuildBranch(user, project, branch, nil)
	} else {
		build, resp = client.NewBuild(user, project, nil)
	}
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

func buildRetryAction(context *cli.Context) error {
	token := context.GlobalString("token")
	buildNum := context.Int("num")
	client := circleci.NewClient(token)
	build, resp := client.RetryBuild(context.String("user"), context.String("project"), buildNum)
	if resp.Success() {
		padding := 1
		writer := tabwriter.NewWriter(os.Stdout, 0, 8, padding, '\t', tabwriter.AlignRight)
		fmt.Fprintln(writer, "Build\tBranch\tStatus\tURL")
		fmt.Fprintln(writer, fmt.Sprintf("%d\t%s\t%s\t%s\t", build.BuildNum, build.Branch, build.Status, build.BuildURL))
		writer.Flush()
		return nil
	}
	failed := errors.New("Failed to retry build")
	return failed
}
