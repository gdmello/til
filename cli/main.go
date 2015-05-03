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
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Path to the config file. Default: $HOME/.tilrc.toml",
			Value: "$HOME/.tilrc.toml",
		},
		cli.BoolTFlag{
			Name:  "verbose",
			Usage: "Verbose output",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "gencfg",
			Usage:  "Generate a sample config",
			Action: GenCfgCommand,
		},
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
