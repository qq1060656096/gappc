package main

import (
	"fmt"
	_ "gappc/cmd"
	"gappc/cmd/commands"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Name = "gappc"
	app.Usage = `gapp command`
	app.HelpName = app.Name
	app.Authors = []cli.Author{
		{
			Name: "andy",
			Email: "1060656096@qq.com",
		},
	}
	fmt.Println(commands.AvailableCommands)
	app.Commands = commands.AvailableCommands
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
