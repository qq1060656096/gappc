package new

import (
	"errors"
	"fmt"
	"gappc/cmd/commands"
	"gappc/utils"
	"github.com/urfave/cli"
	"io"
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
	// crate bootstrap
	createBootstrap(projectPath, stdout)
	// create config
	createConfig(projectPath, stdout)
	// create middleware
	CreateMiddleware(projectPath, stdout, "DemoAuth")
	return nil
}

func createBootstrap(projectPath string, stdout io.Writer) {
	projectBootstrapPath := projectPath + "/" + utils.ProjectDirs[utils.BootstrapDir]
	os.MkdirAll(projectBootstrapPath, os.ModePerm)
	projectBootstrapFile := projectBootstrapPath + "/application.go"
	utils.WriteToFile(projectBootstrapFile, bootstrapTemplate)
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", projectBootstrapFile, "\x1b[0m")
}

func createConfig(projectPath string, stdout io.Writer)  {
	projectConfigPath := projectPath + "/" + utils.ProjectDirs[utils.ConfigDir]
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
}

func CreateMiddleware(projectPath string, stdout io.Writer, middlewareName string) {
	path := projectPath + "/" + utils.ProjectDirs[utils.MiddlewareDir]
	os.MkdirAll(path, os.ModePerm)
	file := path + fmt.Sprintf("/%s.go", utils.ToSnakeCase(middlewareName))
	utils.WriteToFile(file, strings.Replace(MiddlewareTemplate, "{{middlewareName}}", utils.ToCamelCase(middlewareName), -1))
	fmt.Fprintf(stdout, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", file, "\x1b[0m")
}