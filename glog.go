package glog

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
)

var (
	exception = color.New(color.FgRed).Add(color.Bold).FprintfFunc()
	success   = color.New(color.FgGreen).Add(color.Bold).FprintfFunc()
	info      = color.New(color.FgMagenta).Add(color.Bold).FprintfFunc()
	warn      = color.New(color.FgYellow).Add(color.Bold).FprintfFunc()
	normal    = color.New(color.FgWhite).Add(color.Bold).FprintfFunc()

	timeFormat = "2006-01-02 15:04:05.000 (T)"

	mutex sync.Mutex
)

func log(_log func(w io.Writer, format string, args ...interface{}), status string, format string, args ...interface{}) {
	// Lock mutex to prevent logs interfering with eachother
	mutex.Lock()
	defer mutex.Unlock()

	_log(os.Stdout, "[%s] [%s] ", time.Now().Format(timeFormat), status)
	_log(os.Stdout, format, args...)
	_log(os.Stdout, "\n")
}

func Exception(format string, args ...interface{}) {
	log(exception, "ERR", format, args...)
}

func Success(format string, args ...interface{}) {
	// Heh, succ（＞_＜）
	log(success, "SUCC", format, args...)
}

func Info(format string, args ...interface{}) {
	log(info, "INFO", format, args...)
}

func Warn(format string, args ...interface{}) {
	log(warn, "WARN", format, args...)
}

func Log(format string, args ...interface{}) {
	log(normal, "LOG", format, args)
}
