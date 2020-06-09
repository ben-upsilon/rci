package orm

import (
	"database/sql"
	"rci/log"
	"rci/orm/dialect"
	"rci/orm/schema"
	"strings"
)

//入口
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

//会话
type Session struct {
	db      *sql.DB
	sql     strings.Builder
	sqlArgs []interface{}

	dialect dialect.Dialect
	ref     *schema.Schema
}

//当前会话
func (s *Session) current() *Session {
	return s
}

//当前会话
func (s *Session) reset() {
	s.sqlArgs = nil
	s.sql.Reset()
}

//创建会话
func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{db: db, dialect: dialect}
}

//会话执行raw
func (s *Session) RawQuery(sql string, args ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlArgs = append(s.sqlArgs, args...)
	return s
}

//最终执行结果
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.reset()
	log.Info(s.sql.String(), s.sqlArgs)
	if result, err = s.db.Exec(s.sql.String(), s.sqlArgs...); err != nil {
		log.Error(err)
	}
	return
}

//新建数据库
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	PanicIf(err)
	err = db.Ping()
	PanicIf(err)
	dial, ok := dialect.Get(driver)
	if !ok {
		log.Error("数据库驱动模板找不到 > " + driver)
	}
	e = &Engine{db: db, dialect: dial}
	log.Info("数据库连上了")
	return
}

//同步数据表结构
func (e *Engine) SyncSchema() {

}

//关闭数据库
func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("数据库关闭失败")
	}
	log.Info("数据库关闭了")
}

//数据库开始会话,通过engine获取对应的数据库连接
func (e *Engine) NewSession() *Session {
	return New(e.db, e.dialect)
}

//错误处理日志
func PanicIf(err error) {
	if err != nil {
		log.Error(err)
		return
	}
}
