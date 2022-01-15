package tagueuer_test

import (
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/fudanchii/tagueuer"
)

var _ = Describe("Context", func() {
	var (
		ctx *tagueuer.Context
		eg1 struct {
			Name string `field:"name" required:"true"`
		}
		ftype reflect.StructField
	)

	BeforeEach(func() {
		t := reflect.TypeOf(eg1)
		ftype = t.Field(0)
		ctx = tagueuer.NewContext(ftype)
	})

	It("has empty value by default", func() {
		Expect(ctx.FieldValue()).To(Equal(""))
	})

	It("has field tag", func() {
		Expect(ctx.TagValue("field")).To(Equal("name"))
	})

	It("has required tag", func() {
		Expect(ctx.TagValue("required")).To(Equal("true"))
	})
})
