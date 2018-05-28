package tagueuer_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTagueuer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tagueuer Suite")
}
