package commands

import (
	"fmt"
	"net/url"
	"os"
	"text/tabwriter"

	circleci "github.com/mtchavez/circlecigo"
	"github.com/urfave/cli"
)

// BuildsBranchCmd - get recent builds for a branch of a project
var BuildsBranchCmd = cli.Command{
	Name:    "builds-branch",
	Aliases: []string{"bb"},
	Usage:   "get a list of recent builds for the branch of a project",
	Flags:   buildsBranchFlags,
	Action:  buildsBranchAction,
}

var buildsBranchFlags = []cli.Flag{
	userFlag,
	projectFlag,
	branchFlag,
	filterFlag,
	limitFlag,
	offsetFlag,
}

func buildsBranchAction(context *cli.Context) error {
	token := context.GlobalString("token")
	client := circleci.NewClient(token)
	padding := 1
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, padding, '\t', tabwriter.AlignRight)

	params := url.Values{}
	if context.String("filter") != "" {
		params.Set("filter", context.String("filter"))
	}
	if context.String("offset") != "" {
		params.Set("offset", context.String("offset"))
	}
	if context.String("limit") != "" {
		params.Set("limit", context.String("limit"))
	}
	builds, _ := client.ProjectRecentBuildsBranch(context.String("user"), context.String("project"), context.String("branch"), params)
	fmt.Fprintln(writer, "Build\tCommit\tUser\tStatus\t")
	for _, build := range builds {
		subject := fmt.Sprintf("%.50s", build.Subject)
		if len(build.Subject) > 50 {
			subject += "..."
		}
		fmt.Fprintln(writer, fmt.Sprintf("%d\t%s\t%s\t%s\t", build.BuildNum, subject, build.CommitterName, build.Status))
	}
	writer.Flush()
	return nil
}
