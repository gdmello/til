package cli_test

import (
	"github.com/codegangsta/cli"
	. "github.com/kevinjqiu/til/cli"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("GenCfg", func() {
	It("Generates the sample config and output it to stdout", func() {
		GenCfgCommand(&cli.Context{})
	})
})
