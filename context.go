package tagueuer

import (
	"reflect"
	"strings"
)

type Context struct {
	tags  []string
	field reflect.StructField
	value string
}

func NewContext(val reflect.StructField) *Context {
	return &Context{
		tags:  getTagKeys(val.Tag),
		field: val,
	}
}

func (c *Context) TagValue(key string) string {
	return c.val.Tag.Get(key)
}

func (c *Context) FieldValue() string {
	return c.value
}

func (c *Context) FieldHasZeroValue() bool {
	return c.value == ""
}

func (c *Context) set(val string) {
	c.value = val
}

func getTagKeys(tagString string) []string {
	result := []string{}
	tagFields := strings.Split(tagString, " ")
	for tf := range tagFields {
		key := strings.Split(tf, ":")[0]
		result = append(result, key)
	}
	return result
}
