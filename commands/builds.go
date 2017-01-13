package commands

import (
	"fmt"
	"net/url"
	"os"
	"text/tabwriter"
	"time"

	circleci "github.com/mtchavez/circlecigo"
	"github.com/urfave/cli"
)

const (
	timeShortFormat = "2017-01-10 19:59"
	timeAPIFormat   = "2017-01-10T19:59:03.526Z"
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
	branchFlag,
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
	branch := context.String("branch")
	var builds []*circleci.Build
	if branch != "" {
		builds, _ = client.ProjectRecentBuildsBranch(context.String("user"), context.String("project"), branch, params)
	} else {
		builds, _ = client.ProjectRecentBuilds(context.String("user"), context.String("project"), params)
	}
	fmt.Fprintln(writer, "Build\tBranch\tUser\tStatus\tTime\tFinished")
	for _, build := range builds {
		buildTime := fmt.Sprintf("%+v", build.RunTime())
		formattedStopTime := ""
		if build.StopTime != "" {
			parsedStopTime, _ := time.Parse(time.RFC3339, build.StopTime)
			formattedStopTime = parsedStopTime.Format(time.RFC822)
		}
		fmt.Fprintln(writer, fmt.Sprintf("%d\t%s\t%s\t%s\t%s\t%s\t", build.BuildNum, build.Branch, build.CommitterName, build.Status, buildTime, formattedStopTime))
	}
	writer.Flush()
	return nil
}
