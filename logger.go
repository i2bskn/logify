package logo

import (
	"io"
	"os"
	"sync"
)

type Logger struct {
	mu        *sync.Mutex
	level     Level
	formatter Formatter
	out       io.Writer
}

func New() *Logger {
	return &Logger{
		mu:        new(sync.Mutex),
		out:       os.Stdout,
		formatter: &LTSVFormatter{},
	}
}

func (l *Logger) entry() *Entry {
	return newEntry(l)
}

func (l *Logger) Debug(v ...interface{}) {
	if LevelDebug >= l.level {
		l.entry().Debug(v...)
	}
}

var std = New()

func Debug(v ...interface{}) {
	if LevelDebug >= std.level {
		std.Debug(v...)
	}
}
