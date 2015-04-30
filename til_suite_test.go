package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Til Suite")
}
