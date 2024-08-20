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

	mutex sync.Mutex
)

type Logger struct {
	prefix     string
	timeFormat string
}

type Arguments struct {
	prefix  string
	showYMD bool
}

func New(arguments Arguments) *Logger {
	logger := &Logger{}

	if arguments.prefix != "" {
		logger.prefix = "[" + arguments.prefix + "] "
	} else {
		logger.prefix = ""
	}

	if arguments.showYMD {
		logger.timeFormat = "2006-01-02 15:04:05.000 (T)"
	} else {
		logger.timeFormat = "15:04:05.000 (T)"
	}

	return logger
}

func (l *Logger) log(_log func(w io.Writer, format string, args ...interface{}), status string, format string, args ...interface{}) {
	// Lock mutex to prevent logs interfering with eachother
	mutex.Lock()
	defer mutex.Unlock()

	_log(os.Stdout, "%s[%s] [%s] ", l.prefix, time.Now().Format(l.timeFormat), status)
	_log(os.Stdout, format, args...)
	_log(os.Stdout, "\n")
}

func (l *Logger) Exception(format string, args ...interface{}) {
	l.log(exception, "ERR", format, args...)
}

func (l *Logger) Success(format string, args ...interface{}) {
	// Heh, succ（＞_＜）
	l.log(success, "SUCC", format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.log(info, "INFO", format, args...)
}

func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(warn, "WARN", format, args...)
}

func (l *Logger) Log(format string, args ...interface{}) {
	l.log(normal, "LOG", format, args)
}
