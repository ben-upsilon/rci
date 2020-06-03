package dialect

import (
	"reflect"
	"time"
)

type sqlite3 struct {
}

func (s sqlite3) DataTypeOf(t reflect.Value) string {

	switch t.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "integer"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Struct:
		if _, ok := t.Interface().(time.Time); ok {
			return "datetime"
		}
	}

	panic("implement me")
}

var _ Dialect = (*sqlite3)(nil)

func init() {
	Register("sqlite3", &sqlite3{})
}
