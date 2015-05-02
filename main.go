package main

import (
	"os"

	"github.com/kevinjqiu/til/cli"
)

func main() {
	if err := cli.Run(os.Args); err != nil {
		panic(err)
	}
	os.Exit(0)
}
