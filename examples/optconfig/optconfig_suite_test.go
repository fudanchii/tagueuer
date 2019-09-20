package optconfig_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOptconfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Optconfig Suite")
}
