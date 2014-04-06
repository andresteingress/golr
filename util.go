package golr

import (
	"net/url"
	"reflect"
	"strconv"
	"unicode"
)

// lA converts the first character of the given string to lower-case
func lA(s string) string {
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

// structToMap creates url.Values from the given struct
func structToMap(i interface{}) url.Values {
	values := url.Values{}
	fieldValue := reflect.ValueOf(i).Elem()
	typ := fieldValue.Type()

	for i := 0; i < fieldValue.NumField(); i++ {
		f := fieldValue.Field(i)

		var v string
		switch f.Interface().(type) {
			case int, int8, int16, int32, int64:
				v = strconv.FormatInt(f.Int(), 10)
			case uint, uint8, uint16, uint32, uint64:
				v = strconv.FormatUint(f.Uint(), 10)
			case float32:
				v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
			case float64:
				v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
			case []byte:
				v = string(f.Bytes())
			case string:
				v = f.String()
		}

		vType := f.Type()
		if vType != nil {
			if reflect.Zero(vType).Interface() != f.Interface() {
				values.Set(lA(typ.Field(i).Name), v)
			}
		}
	}

	return values
}
