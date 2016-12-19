package commands

import (
	"fmt"
	"net/url"
	"os"
	"text/tabwriter"

	circleci "github.com/mtchavez/circlecigo"
	"github.com/urfave/cli"
)

// BuildsCmd - recent builds for a project
var BuildsCmd = cli.Command{
	Name:    "builds",
	Aliases: []string{"b"},
	Usage:   "get a list of recent builds for a project",
	Flags:   buildsFlags,
	Action:  buildsAction,
}

var buildsFlags = []cli.Flag{
	userFlag,
	projectFlag,
	filterFlag,
	limitFlag,
	offsetFlag,
}

func buildsAction(context *cli.Context) error {
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
	builds, _ := client.ProjectRecentBuilds(context.String("user"), context.String("project"), params)
	fmt.Fprintln(writer, "Branch\tUser\tStatus\t")
	for _, build := range builds {
		fmt.Fprintln(writer, fmt.Sprintf("%s\t%s\t%s\t", build.Branch, build.CommitterName, build.Status))
	}
	writer.Flush()
	return nil
}
