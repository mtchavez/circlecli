package commands

import (
	"fmt"
	"os"
	"text/tabwriter"

	circleci "github.com/mtchavez/circlecigo"
	"github.com/urfave/cli"
)

// RecentBuildsCmd - Recent builds across all projects
var RecentBuildsCmd = cli.Command{
	Name:    "recent-builds",
	Aliases: []string{"rb"},
	Usage:   "get a list of recent builds across all projects",
	Action:  recentBuildAction,
}

func recentBuildAction(context *cli.Context) error {
	token := context.GlobalString("token")
	client := circleci.NewClient(token)
	padding := 1
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, padding, '\t', tabwriter.AlignRight)
	builds, _ := client.RecentBuilds(nil)
	fmt.Fprintln(writer, "Build\tProject\tBranch\tUser\tStatus\t")
	for _, build := range builds {
		project := fmt.Sprintf("%s/%s", build.Username, build.Reponame)
		fmt.Fprintln(writer, fmt.Sprintf("%d\t%s\t%s\t%s\t%s\t", build.BuildNum, project, build.Branch, build.CommitterName, build.Status))
	}
	writer.Flush()
	return nil
}
