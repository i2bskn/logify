package logify

import (
	"io"
	"os"
	"sync"
	"sync/atomic"
)

type Logger interface {
	Level() LogLevel
	SetLevel(LogLevel)
	SetOutput(io.Writer)
	Lock()
	Unlock()
	Write([]byte) (int, error)
	Formatter() Formatter
	// With(fields ...Field) Logger
	Debug(string, ...Field)
	Info(string, ...Field)
	Warn(string, ...Field)
	Error(string, ...Field)
	Fatal(string, ...Field)
	Panic(string, ...Field)
}

type logger struct {
	mu        sync.Mutex
	level     LogLevel
	formatter Formatter
	out       io.Writer
	entryPool sync.Pool
}

func New(w io.Writer, f Formatter, lv LogLevel) Logger {
	return &logger{
		level:     lv,
		formatter: f,
		out:       w,
	}
}

func (l *logger) Level() LogLevel {
	return LogLevel(atomic.LoadInt32((*int32)(&l.level)))
}

func (l *logger) SetLevel(lv LogLevel) {
	atomic.StoreInt32((*int32)(&l.level), int32(lv))
}

func (l *logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = w
}

func (l *logger) Lock() {
	l.mu.Lock()
}

func (l *logger) Unlock() {
	l.mu.Unlock()
}

func (l *logger) Write(b []byte) (int, error) {
	n, err := l.out.Write(b)
	return n, err
}

func (l *logger) Formatter() Formatter {
	return l.formatter
}

func (l *logger) Debug(msg string, fields ...Field) {
	if LevelDebug >= l.Level() {
		e := l.entry()
		e.Debug(msg, fields...)
		l.freeEntry(e)
	}
}

func (l *logger) Info(msg string, fields ...Field) {
	if LevelInfo >= l.Level() {
		e := l.entry()
		e.Info(msg, fields...)
		l.freeEntry(e)
	}
}

func (l *logger) Warn(msg string, fields ...Field) {
	if LevelWarn >= l.Level() {
		e := l.entry()
		e.Warn(msg, fields...)
		l.freeEntry(e)
	}
}

func (l *logger) Error(msg string, fields ...Field) {
	if LevelError >= l.Level() {
		e := l.entry()
		e.Error(msg, fields...)
		l.freeEntry(e)
	}
}

func (l *logger) Fatal(msg string, fields ...Field) {
	if LevelInfo >= l.Level() {
		e := l.entry()
		e.Fatal(msg, fields...)
		l.freeEntry(e)
	}
}

func (l *logger) Panic(msg string, fields ...Field) {
	if LevelInfo >= l.Level() {
		e := l.entry()
		e.Panic(msg, fields...)
		l.freeEntry(e)
	}
}

func (l *logger) entry() *Entry {
	entry, ok := l.entryPool.Get().(*Entry)
	if ok {
		return entry
	}
	return newEntry(l)
}

func (l *logger) freeEntry(e *Entry) {
	e.Reset()
	l.entryPool.Put(e)
}

var std = New(os.Stdout, new(LTSVFormatter), LevelDebug)

func Level() LogLevel {
	return std.Level()
}

func SetLevel(lv LogLevel) {
	std.SetLevel(lv)
}

func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

func Debug(msg string, fields ...Field) {
	if LevelDebug >= std.Level() {
		std.Debug(msg, fields...)
	}
}

func Info(msg string, fields ...Field) {
	if LevelInfo >= std.Level() {
		std.Info(msg, fields...)
	}
}

func Warn(msg string, fields ...Field) {
	if LevelWarn >= std.Level() {
		std.Warn(msg, fields...)
	}
}

func Error(msg string, fields ...Field) {
	if LevelError >= std.Level() {
		std.Error(msg, fields...)
	}
}

func Fatal(msg string, fields ...Field) {
	if LevelFatal >= std.Level() {
		std.Fatal(msg, fields...)
	}
}

func Panic(msg string, fields ...Field) {
	if LevelPanic >= std.Level() {
		std.Panic(msg, fields...)
	}
}
