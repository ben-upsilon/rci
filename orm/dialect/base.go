package dialect

import "reflect"

var dialectMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(t reflect.Value) string
}

func Register(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

func Get(name string) (d Dialect, ok bool) {
	d, ok = dialectMap[name]
	return
}
