package logify

import (
	"io"
	"os"
	"sync"
)

type Logger struct {
	mu        sync.Mutex
	level     LogLevel
	formatter Formatter
	out       io.Writer
	entryPool sync.Pool
}

func New() *Logger {
	return &Logger{
		level:     LevelDebug,
		formatter: new(LTSVFormatter),
		out:       os.Stdout,
	}
}

func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = w
}

func (l *Logger) entry() *Entry {
	entry, ok := l.entryPool.Get().(*Entry)
	if ok {
		return entry
	}
	return newEntry(l)
}

func (l *Logger) Debug(msg string, fields ...Field) {
	if LevelDebug >= l.level {
		entry := l.entry()
		entry.Debug(msg, fields...)
		entry.Free()
	}
}

var std = New()

func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

func Debug(msg string, fields ...Field) {
	if LevelDebug >= std.level {
		std.Debug(msg, fields...)
	}
}
