package api

import (
	"errors"
	"fmt"
	"gappc/cmd/commands"
	"gappc/generate"
	"gappc/utils"
	"github.com/urfave/cli"
	"os"
)

var CmdApi = cli.Command{
	Name:                   "api",
	ShortName:              "",
	Aliases:                nil,
	Usage:                  "apiName",
	UsageText:              "v1/demo/user",
	Description:            "",
	ArgsUsage:              "",
	Category:               "",
	BashComplete:           nil,
	Before:                 nil,
	After:                  nil,
	Action:                 Run,
	OnUsageError:           nil,
	Subcommands:            nil,
	Flags:                  nil,
	SkipFlagParsing:        false,
	SkipArgReorder:         false,
	HideHelp:               false,
	Hidden:                 false,
	UseShortOptionHandling: false,
	HelpName:               "",
	CustomHelpTemplate:     "",
}

func init()  {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdApi)
}

func Run(ctx * cli.Context) (err error) {
	stdout := os.Stdout
	if ctx.NArg() < 1 {
		return errors.New(fmt.Sprintf("Argument [apiName] is missing"))
	}

	projectName, projectPath, err := utils.GetProjectInfo()
	if err != nil {
		return
	}
	projectPath = fmt.Sprintf("%s/%s", projectPath, projectName)

	apiName := ctx.Args().Get(0)
	generate.CreateApi(projectPath, stdout, apiName)
	return nil
}