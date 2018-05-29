package tagueuer

import (
	"reflect"
	"strings"
)

// Context represents the current context of the struct field
type Context struct {
	tags  []string
	field reflect.StructField
	value string
}

// NewContext creates new context from the given struct field
func NewContext(val reflect.StructField) *Context {
	return &Context{
		tags:  getTagKeys(string(val.Tag)),
		field: val,
	}
}

// TagValue returns the value of the given tag key
func (c *Context) TagValue(key string) string {
	return c.field.Tag.Get(key)
}

// FieldValue returns the current value to be set for the current field
func (c *Context) FieldValue() string {
	return c.value
}

// FieldHasZeroValue returns true if  c.value is empty
func (c *Context) FieldHasZeroValue() bool {
	return c.value == ""
}

func (c *Context) set(val string) {
	c.value = val
}

func getTagKeys(tagString string) []string {
	result := []string{}
	tagFields := strings.Split(tagString, " ")
	for _, tf := range tagFields {
		key := strings.Split(tf, ":")[0]
		result = append(result, key)
	}
	return result
}
