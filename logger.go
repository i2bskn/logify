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
	Serializer() Serializer
	With(fields ...Field) Logger
	Debug(string, ...Field)
	Info(string, ...Field)
	Warn(string, ...Field)
	Error(string, ...Field)
	Fatal(string, ...Field)
	Panic(string, ...Field)
}

type coreLogger struct {
	mu         sync.Mutex
	level      LogLevel
	serializer Serializer
	out        io.Writer
	entryPool  sync.Pool
}

func New(w io.Writer, s Serializer, lv LogLevel) Logger {
	return &coreLogger{
		level:      lv,
		serializer: s,
		out:        w,
	}
}

func (cl *coreLogger) Level() LogLevel {
	return LogLevel(atomic.LoadInt32((*int32)(&cl.level)))
}

func (cl *coreLogger) SetLevel(lv LogLevel) {
	atomic.StoreInt32((*int32)(&cl.level), int32(lv))
}

func (cl *coreLogger) SetOutput(w io.Writer) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.out = w
}

func (cl *coreLogger) Lock() {
	cl.mu.Lock()
}

func (cl *coreLogger) Unlock() {
	cl.mu.Unlock()
}

func (cl *coreLogger) Write(b []byte) (int, error) {
	n, err := cl.out.Write(b)
	return n, err
}

func (cl *coreLogger) Serializer() Serializer {
	return cl.serializer
}

func (cl *coreLogger) With(fields ...Field) Logger {
	return newFieldedLogger(cl, fields)
}

func (cl *coreLogger) Debug(msg string, fields ...Field) {
	if LevelDebug >= cl.Level() {
		e := cl.entry()
		e.Debug(msg, fields...)
		cl.freeEntry(e)
	}
}

func (cl *coreLogger) Info(msg string, fields ...Field) {
	if LevelInfo >= cl.Level() {
		e := cl.entry()
		e.Info(msg, fields...)
		cl.freeEntry(e)
	}
}

func (cl *coreLogger) Warn(msg string, fields ...Field) {
	if LevelWarn >= cl.Level() {
		e := cl.entry()
		e.Warn(msg, fields...)
		cl.freeEntry(e)
	}
}

func (cl *coreLogger) Error(msg string, fields ...Field) {
	if LevelError >= cl.Level() {
		e := cl.entry()
		e.Error(msg, fields...)
		cl.freeEntry(e)
	}
}

func (cl *coreLogger) Fatal(msg string, fields ...Field) {
	if LevelInfo >= cl.Level() {
		e := cl.entry()
		e.Fatal(msg, fields...)
		cl.freeEntry(e)
	}
}

func (cl *coreLogger) Panic(msg string, fields ...Field) {
	if LevelInfo >= cl.Level() {
		e := cl.entry()
		e.Panic(msg, fields...)
		cl.freeEntry(e)
	}
}

func (cl *coreLogger) entry() *Entry {
	entry, ok := cl.entryPool.Get().(*Entry)
	if ok {
		return entry
	}
	return newEntry(cl)
}

func (cl *coreLogger) freeEntry(e *Entry) {
	e.Reset()
	cl.entryPool.Put(e)
}

var std = New(os.Stdout, new(LTSVSerializer), LevelDebug)

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
