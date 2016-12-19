package commands

import "github.com/urfave/cli"

// AllCommands contains the commands for the CLI
var AllCommands []cli.Command
var (
	userFlag = cli.StringFlag{
		Name:   "user, u",
		Value:  "",
		Usage:  "user or org",
		EnvVar: "CIRCLECI_ORG,CIRCLECI_USER",
	}
	projectFlag = cli.StringFlag{
		Name:   "project, p",
		Value:  "",
		Usage:  "project name of repository",
		EnvVar: "CIRCLECI_PROJECT",
	}
	branchFlag = cli.StringFlag{
		Name:   "branch, b",
		Value:  "master",
		Usage:  "project branch name, default master",
		EnvVar: "CIRCLECI_BRANCH",
	}
	filterFlag = cli.StringFlag{
		Name:   "filter, f",
		Value:  "",
		Usage:  "filter by status of build ie. successful, failed, running",
		EnvVar: "CIRCLECI_ORG,CIRCLECI_USER",
	}
	limitFlag = cli.StringFlag{
		Name:  "limit, l",
		Value: "",
		Usage: "limit of builds to return, default 100",
	}
	offsetFlag = cli.StringFlag{
		Name:  "offset, o",
		Value: "",
		Usage: "offset of builds to start from",
	}
)

func init() {
	AllCommands = []cli.Command{
		RecentBuildsCmd,
		BuildsCmd,
		BuildsBranchCmd,
	}
}
