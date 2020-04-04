package theme

import (
	"errors"
	"fmt"
	"gappc/cmd/commands"
	"gappc/generate"
	"gappc/utils"
	"github.com/urfave/cli"
	"os"
)

var CmdTheme = cli.Command{
	Name:                   "theme",
	ShortName:              "",
	Aliases:                nil,
	Usage:                  "themeName subDirName fileName",
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
	commands.AvailableCommands = append(commands.AvailableCommands, CmdTheme)
}

func Run(ctx * cli.Context) (err error) {
	stdout := os.Stdout
	if ctx.NArg() < 1 {
		return errors.New(fmt.Sprintf("Argument [themeName] is missing"))
	}

	if ctx.NArg() < 2 {
		return errors.New(fmt.Sprintf("Argument [subDirName] is missing"))
	}

	if ctx.NArg() < 3 {
		return errors.New(fmt.Sprintf("Argument [fileName] is missing"))
	}

	projectName, projectPath, err := utils.GetProjectInfo()
	if err != nil {
		return
	}
	projectPath = fmt.Sprintf("%s/%s", projectPath, projectName)

	themeName := ctx.Args().Get(0)
	subDirName := ctx.Args().Get(1)
	fileName := ctx.Args().Get(2)
	generate.CreateTheme(projectPath, stdout, themeName, subDirName, fileName)
	return nil
}