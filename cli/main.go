package cli

import (
	_ "log"

	"github.com/codegangsta/cli"
)

func Run(args []string) error {
	app := cli.NewApp()
	app.Name = "TIL"
	app.Version = Version
	app.Usage = "Manage your Today-I-Learned entries"
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "Initialize a TIL repo",
			Action: InitCommand,
		},
		{
			Name:   "new",
			Usage:  "Create a new TIL entry",
			Action: NewCommand,
		},
		{
			Name:   "show",
			Usage:  "Show existing TIL entries",
			Action: ShowCommand,
		},
	}

	return app.Run(args)
}
