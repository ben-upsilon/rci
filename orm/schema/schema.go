package schema

import (
	"go/ast"
	"rci/orm/dialect"
	"reflect"
)

//解析字段
type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Model    interface{}
	Name     string
	FieldMap map[string]*Field
}

func Parse(target interface{}, d dialect.Dialect) *Schema {
	tableType := reflect.Indirect(reflect.ValueOf(target)).Type()
	schema := &Schema{
		Model:    target,
		Name:     tableType.Name(),
		FieldMap: make(map[string]*Field),
	}
	for i := 0; i < tableType.NumField(); i++ {
		mField := tableType.Field(i)
		if !mField.Anonymous && ast.IsExported(mField.Name) {
			f := &Field{
				Name: mField.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(mField.Type))),
			}
			if t, ok := mField.Tag.Lookup("orm"); ok {
				f.Tag = t
			}
			schema.FieldMap[mField.Name] = f
		}
	}
	return schema
}
