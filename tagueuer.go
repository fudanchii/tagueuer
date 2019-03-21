package tagueuer

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// CallbackFunc is a function which is called when
// tagueuer encounter tag declaration.
// The provided function will be called with `*tagueuer.Context`
// passed as parameter. This function is expected to return string
// which is the value to be used for this struct field, and an error,
// if any.
type CallbackFunc func(ctx *Context) (string, error)

// DefaultsFuncList is a function which is called when
// default tag is set to function call, e.g:  `funcname()`
// tagueuer will mach the function name listed here
type DefaultValues map[string]string

// Tagueuer main struct, stores a map of callback functions
// to be used when populating the struct.
type Tagueuer struct {
	funcs map[string]CallbackFunc
}

var (
	defaults = make(DefaultValues)
)

// New creates new Tagueuer struct. With default handler
// for `default` and `required` tag provided.
func New() *Tagueuer {
	return &Tagueuer{map[string]CallbackFunc{
		"default":  setDefault,
		"required": checkRequired,
	}}
}

func Defaults(name string, def string) {
	defaults[name] = def
}

// On assigns new callback for the given tag key. This can also be used
// to override the default handlers
func (t *Tagueuer) On(name string, cb CallbackFunc) {
	t.funcs[name] = cb
}

// ParseInto parse the struct tags from given `obj` and populate the resulted data
// back into `obj`. obj is expected to be a pointer of struct.
func (t *Tagueuer) ParseInto(obj interface{}) error {
	vObj := reflect.ValueOf(obj)

	if !isPtrToStruct(vObj) {
		return fmt.Errorf("can only parse into pointer of struct, but %s was given", typeName(vObj))
	}

	stObj := vObj.Elem()
	for i := 0; i < stObj.NumField(); i++ {
		field := stObj.Field(i)

		if !field.CanSet() {
			continue // ignore unexported field
		}

		ctx := NewContext(stObj.Type().Field(i))

		for _, tag := range ctx.tags {
			if fn, ok := t.funcs[tag]; ok {
				strFieldVal, err := fn(ctx)

				if err != nil {
					return err
				}

				ctx.set(strFieldVal)
			}
		}

		if !ctx.FieldHasZeroValue() {
			switch field.Kind() {
			case reflect.Bool:
				v, err := strconv.ParseBool(ctx.FieldValue())
				if err != nil {
					return err
				}
				field.SetBool(v)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				v, err := strconv.ParseInt(ctx.FieldValue(), 10, 64)
				if err != nil {
					return err
				}
				field.SetInt(v)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				v, err := strconv.ParseUint(ctx.FieldValue(), 10, 64)
				if err != nil {
					return err
				}
				field.SetUint(v)
			case reflect.Float32, reflect.Float64:
				v, err := strconv.ParseFloat(ctx.FieldValue(), 64)
				if err != nil {
					return err
				}
				field.SetFloat(v)
			case reflect.String:
				field.SetString(ctx.FieldValue())
			default:
				return fmt.Errorf("type `%s` is not supported yet", field.Type().Name())
			}
		}
	}

	return nil
}

func setDefault(c *Context) (string, error) {
	def := c.TagValue("default")
	if c.FieldHasZeroValue() {
		if strings.HasPrefix(def, "&") {
			return defaults[strings.Trim(def, "&")], nil
		}
		return def, nil
	}
	return c.FieldValue(), nil
}

func checkRequired(c *Context) (string, error) {
	if c.FieldHasZeroValue() {
		return c.FieldValue(), fmt.Errorf("field is required but empty for tag: `%s`", c.field.Tag)
	}
	return c.FieldValue(), nil
}

func isPtrToStruct(v reflect.Value) bool {
	return v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct
}

func typeName(v reflect.Value) string {
	if v.Kind() == reflect.Ptr {
		return "pointer to " + v.Elem().Type().Name()
	}
	return v.Type().Name()
}
