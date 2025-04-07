package i2c

import (
	"sync"

	logger "github.com/d2r2/go-logger"
)

// You can manage verbosity of log output
// in the package by changing last parameter value
// (comment/uncomment corresponding lines).

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

var (
	lg   Logger = logger.NewPackageLogger("i2c", logger.WarnLevel)
	lock sync.RWMutex
)

func SetLogger(logger Logger) {
	lock.Lock()
	defer lock.Unlock()

	lg = logger
}
