package app

import (
	"errors"
	"fmt"
	"gappc/cmd/commands"
	"gappc/generate"
	"gappc/utils"
	"github.com/urfave/cli"
	"os"
)

var CmdApp = cli.Command{
	Name:                   "app",
	ShortName:              "",
	Aliases:                nil,
	Usage:                  "appName",
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
	commands.AvailableCommands = append(commands.AvailableCommands, CmdApp)
}

func Run(ctx * cli.Context) (err error) {
	stdout := os.Stdout
	if ctx.NArg() < 1 {
		return errors.New(fmt.Sprintf("Argument [appName] is missing"))
	}

	projectName, projectPath, err := utils.GetProjectInfo()
	if err != nil {
		return
	}
	projectPath = fmt.Sprintf("%s/%s", projectPath, projectName)

	appName := ctx.Args().Get(0)
	generate.CreateApp(projectPath, stdout, appName)
	return nil
}