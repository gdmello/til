package cli

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func InitCommand(c *cli.Context) {
}

func NewCommand(c *cli.Context) {
}

func ShowCommand(c *cli.Context) {
}

const SampleConfig = `["til"]
repository_path="$HOME/.til"
`

func GenCfgCommand(c *cli.Context) {
	fmt.Println(SampleConfig)
}
