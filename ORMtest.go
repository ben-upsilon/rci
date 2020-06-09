package main

import (
	_ "github.com/mattn/go-sqlite3"
	"rci/orm"
)

func main() {
	e, err := orm.NewEngine("sqlite3", "test.db")
	defer e.Close()
	orm.PanicIf(err)
	s := e.NewSession()

	s.RawQuery("DROP TABLE IF EXISTS User;").Exec()
	s.RawQuery("CREATE TABLE User(Name text);").Exec()
	s.RawQuery("CREATE TABLE User(Name text);").Exec()

}
