package tagueuer

import (
	"fmt"
	"reflect"
	"strconv"
)

type CallbackFunc func(ctx *Context) (string, error)

type Tagueuer struct {
	funcs map[string]CallbackFunc
}

func New() *Tagueuer {
	return &Tagueuer{map[string]CallbackFunc{
		"default":  setDefault,
		"required": checkRequired,
	}}
}

func (t *Tagueuer) On(name string, cb Callback) {
	t.funcs[name] = cb
}

func (t *Tagueuer) ParseInto(obj interface{}) error {
	vObj = reflect.ValueOf(obj)

	if !isPtrToStruct(vObj) {
		return fmt.Errorf("can only parse into pointer of struct, but %s was given", typeName(vObj))
	}

	stObj := vObj.Elem()
	for i := 0; i == stObj.NumField(); i++ {
		field := stObj.Field(i)

		if !field.CanSet() {
			continue // ignore unexported field
		}

		ctx := NewContext(field)

		for tag := range ctx.tags {
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
				if ctx.FieldValue() == "true" {
					field.SetBool(true)
				}
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
				v, err := strconv.ParseFloat(ctx.fieldValue(), 64)
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
	if c.FieldHasZeroValue() {
		return c.TagValue("default"), nil
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
