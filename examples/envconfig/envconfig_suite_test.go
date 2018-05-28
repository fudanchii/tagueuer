package envconfig_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEnvconfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Envconfig Suite")
}
