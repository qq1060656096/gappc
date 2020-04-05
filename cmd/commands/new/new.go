package new

import (
	"errors"
	"fmt"
	"gappc/cmd/commands"
	"gappc/generate"
	"gappc/utils"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"strings"
)

var CmdNew = cli.Command{
	Name:                   "new",
	ShortName:              "",
	Aliases:                nil,
	Usage:                  "projectName",
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
		return errors.New(fmt.Sprintf("Argument [projectName] is missing"))
	}
	projectName := strings.ToLower(ctx.Args().Get(0))
	cmdPath, err := utils.GetCmdPath()
	if err != nil {
		return
	}
	projectPath := fmt.Sprintf("%s/%s", cmdPath, projectName)

	ok, err := utils.CheckPathExists(projectPath)
	if ok {
		fmt.Fprintf(stdout, "\t%s%sfail%s\t%s project exist%s\n", "\x1b[33m", "\x1b[31m", "\x1b[21m", projectPath, "\x1b[0m")
		return
	}

	// create api
	generate.CreateApi(projectPath, stdout, "v1/demo/DemoApi")
	// create app
	generate.CreateApp(projectPath, stdout, "v1/demo/DemoApp")
	// create theme
	generate.CreateTheme(projectPath, stdout, "default", "demo", "create.tmpl")
	generate.CreateTheme(projectPath, stdout, "default", "demo", "delete.tmpl")
	generate.CreateTheme(projectPath, stdout, "default", "demo", "detail.tmpl")
	generate.CreateTheme(projectPath, stdout, "default", "demo", "list.tmpl")
	generate.CreateTheme(projectPath, stdout, "default", "demo", "update.tmpl")


	// crate bootstrap
	generate.CreateBootstrap(projectPath, stdout)
	// create config
	generate.CreateConfig(projectPath, stdout)
	// create middleware
	generate.CreateMiddleware(projectPath, stdout, "DemoAuth")

	// create routers
	generate.CreateRouters(projectPath, stdout)
	// create log
	generate.CreateLog(projectPath, stdout)
	// create main
	generate.CreateMain(projectPath, stdout, projectName)

	// 2. 初始化项目
	cmd := exec.Command( "go", "mod", "init", projectName)
	cmd.Dir = cmdPath+"/"+projectName
	cmd.Stdout = os.Stdout
	err = cmd.Start()
	if err != nil {
		fmt.Fprintf(stdout, "\t%s%sfail%s\t %s/go.mod%s\n", "\x1b[33m", "\x1b[31m", "\x1b[21m", cmd.Dir, "\x1b[0m")
		return
	}
	return nil
}

