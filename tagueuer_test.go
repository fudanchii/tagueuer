package tagueuer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fudanchii/tagueuer"
)

var _ = Describe("Tagueuer", func() {
	var (
		tagParser = tagueuer.New()
		eg2       struct {
			Name  string            `field_name:"student_name" required:"true"`
			Class int               `field_name:"class" required:"true"`
			Year  int               `field_name:"year" required:"true"`
			meta  map[string]string `field_name:"meta"`
		}

		data = map[string]string{
			"student_name": "Chitanda Eru",
			"class":        "3",
			"year":         "1",
		}
	)

	BeforeEach(func() {
		tagParser.On("field_name", func(c *tagueuer.Context) (string, error) {
			return data[c.TagValue("field_name")], nil
		})
		tagParser.ParseInto(&eg2)
	})

	Context("parse hash into struct", func() {
		It("has correct value for name", func() {
			Expect(eg2.Name).To(Equal("Chitanda Eru"))
		})

		It("has correct value for class", func() {
			Expect(eg2.Class).To(Equal(3))
		})

		It("has correct value for year", func() {
			Expect(eg2.Year).To(Equal(1))
		})
	})
})
