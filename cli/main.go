package cli

import (
	"log"
	_ "os"

	"github.com/mitchellh/cli"
)

func Run(args []string) int {
	c := cli.NewCLI("til", Version)
	c.Args = args

	c.Commands = Commands()

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	return exitStatus
}
