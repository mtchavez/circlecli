package main

import (
	"fmt"
	"net/url"
	"os"
	"text/tabwriter"

	"github.com/mtchavez/circlecigo"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "circlecli"
	app.Usage = "Interact with the CircleCI REST API"
	app.Version = "0.1.0"
	app.EnableBashCompletion = true
	token := os.Getenv("CIRCLECI_TOKEN")
	client := circleci.NewClient(token)
	padding := 3
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, padding, '\t', tabwriter.AlignRight)

	app.Commands = []cli.Command{
		{
			Name:    "recent-builds",
			Aliases: []string{"rb"},
			Usage:   "get a list of recent builds across all projects",
			Action: func(c *cli.Context) error {
				builds, _ := client.RecentBuilds(nil)
				fmt.Fprintln(writer, "Branch\tUser\tStatus\t")
				for _, build := range builds {
					fmt.Fprintln(writer, fmt.Sprintf("%s\t%s\t%s\t", build.Branch, build.CommitterName, build.Status))
				}
				writer.Flush()
				return nil
			},
		},
		{
			Name:    "builds",
			Aliases: []string{"b"},
			Usage:   "get a list of recent builds for a project",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "user, u",
					Value:  "",
					Usage:  "user or org",
					EnvVar: "CIRCLECI_ORG,CIRCLECI_USER",
				},
				cli.StringFlag{
					Name:   "project, p",
					Value:  "",
					Usage:  "project name of repository",
					EnvVar: "CIRCLECI_PROJECT",
				},
				cli.StringFlag{
					Name:   "filter, f",
					Value:  "",
					Usage:  "filter by status of build ie. successful, failed, running",
					EnvVar: "CIRCLECI_ORG,CIRCLECI_USER",
				},
				cli.StringFlag{
					Name:  "limit, l",
					Value: "",
					Usage: "limit of builds to return, default 100",
				},
				cli.StringFlag{
					Name:  "offset, o",
					Value: "",
					Usage: "offset of builds to start from",
				},
			},
			Action: func(c *cli.Context) error {
				params := url.Values{}
				if c.String("filter") != "" {
					params.Set("filter", c.String("filter"))
				}
				if c.String("offset") != "" {
					params.Set("offset", c.String("offset"))
				}
				if c.String("limit") != "" {
					params.Set("limit", c.String("limit"))
				}
				builds, _ := client.ProjectRecentBuilds(c.String("user"), c.String("project"), params)
				fmt.Fprintln(writer, "Branch\tUser\tStatus\t")
				for _, build := range builds {
					fmt.Fprintln(writer, fmt.Sprintf("%s\t%s\t%s\t", build.Branch, build.CommitterName, build.Status))
				}
				writer.Flush()
				return nil
			},
		},
		{
			Name:    "builds-branch",
			Aliases: []string{"bb"},
			Usage:   "get a list of recent builds for the branch of a project",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "user, u",
					Value:  "",
					Usage:  "user or org",
					EnvVar: "CIRCLECI_ORG,CIRCLECI_USER",
				},
				cli.StringFlag{
					Name:   "project, p",
					Value:  "",
					Usage:  "project name of repository",
					EnvVar: "CIRCLECI_PROJECT",
				},
				cli.StringFlag{
					Name:   "branch, b",
					Value:  "master",
					Usage:  "project branch name, default master",
					EnvVar: "CIRCLECI_BRANCH",
				},
				cli.StringFlag{
					Name:   "filter, f",
					Value:  "",
					Usage:  "filter by status of build ie. successful, failed, running",
					EnvVar: "CIRCLECI_ORG,CIRCLECI_USER",
				},
				cli.StringFlag{
					Name:  "limit, l",
					Value: "30",
					Usage: "limit of builds to return, default 30",
				},
				cli.StringFlag{
					Name:  "offset, o",
					Value: "",
					Usage: "offset of builds to start from",
				},
			},
			Action: func(c *cli.Context) error {
				params := url.Values{}
				if c.String("filter") != "" {
					params.Set("filter", c.String("filter"))
				}
				if c.String("offset") != "" {
					params.Set("offset", c.String("offset"))
				}
				if c.String("limit") != "" {
					params.Set("limit", c.String("limit"))
				}
				builds, _ := client.ProjectRecentBuildsBranch(c.String("user"), c.String("project"), c.String("branch"), params)
				fmt.Fprintln(writer, "Branch\tUser\tStatus\t")
				for _, build := range builds {
					fmt.Fprintln(writer, fmt.Sprintf("%s\t%s\t%s\t", build.Branch, build.CommitterName, build.Status))
				}
				writer.Flush()
				return nil
			},
		},
	}
	app.Run(os.Args)
}
