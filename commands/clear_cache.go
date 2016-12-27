package commands

import (
	"errors"
	"fmt"
	"strings"

	circleci "github.com/mtchavez/circlecigo"
	"github.com/urfave/cli"
)

// ClearCacheCmd - get env settings for a project
var ClearCacheCmd = cli.Command{
	Name:   "clear-cache",
	Usage:  "clear project cache",
	Flags:  clearCacheFlags,
	Action: clearCacheAction,
}

var clearCacheFlags = []cli.Flag{
	userFlag,
	projectFlag,
}

func clearCacheAction(context *cli.Context) error {
	token := context.GlobalString("token")
	client := circleci.NewClient(token)
	cache, resp := client.ProjectClearCache(context.String("user"), context.String("project"))
	if resp.Success() {
		fmt.Printf("%s\n", strings.Title(cache.Status))
		return nil
	}
	failed := errors.New("Failed to clear cache")
	return failed
}
