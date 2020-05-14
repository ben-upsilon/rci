package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

const (
	InfoLevel = iota
	ErrorLevel
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[ERROR]\033[0m >", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[INFO ]\033[0m >", log.LstdFlags|log.Lshortfile)
	mLogger  = []*log.Logger{errorLog, infoLog}
	mutex    sync.Mutex
)

var (
	Error = errorLog.Println

	Info = infoLog.Println
)

func SetLogLevel(level int) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, logger := range mLogger {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}
}

func info(info string) {

}
