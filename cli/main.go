package cli

import (
	"log"
	_ "os"

	"github.com/mitchellh/cli"
)

func Run(args []string) int {
	c := cli.NewCLI("til", Version)
	c.Args = args

	// c.Commands = map[string]cli.CommandFactory{
	//     "foo": fooCommandFactory,
	//     "bar": barCommandFactory,
	// }

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	return exitStatus
}
