package orm

import (
	"database/sql"
	"rci/log"
	"strings"
)

//入口
type Engine struct {
	db *sql.DB
}

//会话
type Session struct {
	db      *sql.DB
	sql     strings.Builder
	sqlArgs []interface{}
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
func New(db *sql.DB) *Session {
	return &Session{db: db}
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
	e = &Engine{db: db}
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

//数据库开始会话
func (e *Engine) NewSession() *Session {
	return New(e.db)
}

//错误处理日志
func PanicIf(err error) {
	if err != nil {
		log.Error(err)
		return
	}
}
