package schema

import (
	"encoding/json"
	"rci/orm/dialect"
	"testing"
)

type User struct {
	Name string `orm:"pk"`
	Age  int
}

var dial, _ = dialect.Get("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, dial)
	t.Log(schema.Name)
	t.Log(schema.Model)
	t.Log(schema.FieldMap)
	tJson, _ := json.MarshalIndent(schema, "", "\t")
	t.Log(string(tJson))
	t.Logf("%+v\n", schema.FieldMap)
}
