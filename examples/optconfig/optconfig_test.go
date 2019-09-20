package optconfig_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fudanchii/tagueuer/examples/optconfig"
)

var _ = Describe("Optconfig", func() {
	var (
		conf struct {
			host string `opt:"" default:"localhost" desc:"Host to bind and listen to, default to localhost."`
			port int    `opt:"" default:"9898" desc:"Port to bind and listen to, default to 9898."`
		}
		err     error
		optconf = optconfig.New()
	)

	Describe("Parsing opt flag", func() {
		BeforeEach(func() {
			os.Args = []string{"test", "--host", "0.0.0.0", "--port", "7890"}
			err = optconf.ParseInto(&conf)
		})

		It("doesn't return error", func() {
			Expect(err).To(Equal(nil))
		})
	})
})
