package new

import (
	"errors"
	"fmt"
	"gappc/cmd/commands"
	"gappc/utils"
	"github.com/urfave/cli"
	"os"
	"strings"
)

var CmdNew = cli.Command{
	Name:                   "new",
	ShortName:              "",
	Aliases:                nil,
	Usage:                  "project name",
	UsageText:              "",
	Description:            "",
	ArgsUsage:              "",
	Category:               "",
	BashComplete:           nil,
	Before:                 nil,
	After:                  nil,
	Action:                 CreateProject,
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
	commands.AvailableCommands = append(commands.AvailableCommands, CmdNew)
}

func CreateProject(ctx * cli.Context) (err error) {
	stdout := os.Stdout
	if ctx.NArg() < 1 {
		return errors.New(fmt.Sprintf("Argument [appname] is missing"))
	}
	projectName := strings.ToLower(ctx.Args().Get(0))
	cmdPath, err := utils.GetCmdPath()
	if err != nil {
		return
	}
	projectPath := fmt.Sprintf("%s/%s", cmdPath, projectName)

	projectConfigPath := projectPath + "/" + utils.ProjectDirs["config"]
	os.MkdirAll(projectConfigPath, os.ModePerm)
	configFile := projectConfigPath + "/.app.env"
	utils.WriteToFile(configFile, appConfigTemplate)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", configFile, "\x1b[0m")
	configFile = projectConfigPath + "/.cache.env"
	utils.WriteToFile(configFile, cacheConfigTemplate)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", configFile, "\x1b[0m")
	configFile = projectConfigPath + "/.db.env"
	utils.WriteToFile(configFile, dbConfigTemplate)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", configFile, "\x1b[0m")
	return nil
}



