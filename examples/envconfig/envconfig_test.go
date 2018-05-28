package envconfig_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fudanchii/tagueuer/examples/envconfig"
)

var _ = Describe("Envconfig", func() {
	var (
		eg3 struct {
			Name     string `envconfig:"app_name" default:"Tagueuer"`
			Host     string `envconfig:"host" default:"0.0.0.0"`
			Port     int    `envconfig:"port" default:"8769"`
			SavePath string `envconfig:"save_path" required:"true"`
		}
		err error
	)

	Describe("Populate struct from environment variable", func() {
		Context("When env variable is not populated", func() {
			BeforeEach(func() {
				err = envconfig.ParseInto(&eg3)
			})

			It("returns error", func() {
				Expect(err).NotTo(Equal(nil))
			})

			It("mentions error on save_path", func() {
				Expect(err.Error()).To(Equal("field is required but empty for tag: `envconfig:\"save_path\" required:\"true\"`"))
			})
		})

		Context("When some env variable is populated", func() {
			BeforeEach(func() {
				os.Setenv("SAVE_PATH", "/tmp/tagueuer")
				err = envconfig.ParseInto(&eg3)
			})

			It("doesn't return error", func() {
				Expect(err).To(BeNil())
			})

			It("populate save path", func() {
				Expect(eg3.SavePath).To(Equal("/tmp/tagueuer"))
			})

			It("populate port by default value", func() {
				Expect(eg3.Port).To(Equal(8769))
			})

			It("populate host by default value", func() {
				Expect(eg3.Host).To(Equal("0.0.0.0"))
			})

			It("populate name by default value", func() {
				Expect(eg3.Name).To(Equal("Tagueuer"))
			})
		})
	})
})
