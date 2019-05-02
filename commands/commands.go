package commands

import "github.com/urfave/cli"

// AllCommands contains the commands for the CLI
var AllCommands []cli.Command

// GlobalFlags contains the flags allowed for each command
var GlobalFlags []cli.Flag
var (
	buildNumFlag = cli.StringFlag{
		Name:  "num",
		Value: "",
		Usage: "build number",
	}
	branchFlag = cli.StringFlag{
		Name:   "branch, b",
		Value:  "",
		Usage:  "project branch name, default master",
		EnvVar: "CIRCLECI_BRANCH",
	}
	envVarFlag = cli.StringFlag{
		Name:  "var",
		Value: "",
		Usage: "use VAR name for get and VAR=value for set",
	}
	filterFlag = cli.StringFlag{
		Name:  "filter, f",
		Value: "",
		Usage: "filter by status of build ie. successful, failed, running",
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
	projectFlag = cli.StringFlag{
		Name:   "project, p",
		Value:  "",
		Usage:  "project name of repository",
		EnvVar: "CIRCLECI_PROJECT",
	}
	tokenFlag = cli.StringFlag{
		Name:   "token, t",
		Value:  "",
		Usage:  "circle-token param for authentication",
		EnvVar: "CIRCLECI_TOKEN",
	}
	userFlag = cli.StringFlag{
		Name:   "user, u",
		Value:  "",
		Usage:  "user or org",
		EnvVar: "CIRCLECI_ORG,CIRCLECI_USER",
	}
)

func init() {
	AllCommands = []cli.Command{
		BuildCmd,
		BuildsCmd,
		ClearCacheCmd,
		EnvCmd,
		RecentBuildsCmd,
	}
	GlobalFlags = []cli.Flag{
		tokenFlag,
	}
}
