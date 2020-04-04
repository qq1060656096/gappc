package middleware

import (
	"errors"
	"fmt"
	"gappc/cmd/commands"
	"gappc/generate"
	"gappc/utils"
	"github.com/urfave/cli"
	"os"
)

var CmdMiddleware = cli.Command{
	Name:                   "middleware",
	ShortName:              "",
	Aliases:                nil,
	Usage:                  "middlewareName",
	UsageText:              "",
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
	commands.AvailableCommands = append(commands.AvailableCommands, CmdMiddleware)
}

func Run(ctx * cli.Context) (err error) {
	stdout := os.Stdout
	if ctx.NArg() < 1 {
		return errors.New(fmt.Sprintf("Argument [middlewareName] is missing"))
	}

	projectName, projectPath, err := utils.GetProjectInfo()
	if err != nil {
		return
	}
	projectPath = fmt.Sprintf("%s/%s", projectPath, projectName)

	middlewareName := ctx.Args().Get(0)
	generate.CreateMiddleware(projectPath, stdout, middlewareName)
	return nil
}