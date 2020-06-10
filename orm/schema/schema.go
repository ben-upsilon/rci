package schema

import (
	"go/ast"
	"rci/orm"
	"rci/orm/dialect"
	"reflect"
)

//解析字段
type Field struct {
	Name string
	Type string
	Tag  string
}

//对应构建表结构
type Schema struct {
	Model    interface{}       //模型
	Name     string            //表名
	FieldMap map[string]*Field //字段
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
