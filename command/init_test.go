package command_test

import (
	. "github.com/kevinjqiu/til/command"

	. "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"
)

var _ = Describe("Init Command", func() {

	It("Creates .til folder in the current directory", func() {
		cmd := InitCommand{}
		cmd.Run([]string{})
	})
})
