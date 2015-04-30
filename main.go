package main

import (
	"os"

	"github.com/kevinjqiu/til/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
