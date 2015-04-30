package cli

import (
	"github.com/kevinjqiu/til/command"
	"github.com/mitchellh/cli"
)

func Commands() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &command.InitCommand{}, nil
		},
	}
}
