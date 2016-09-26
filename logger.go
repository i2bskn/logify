package logify

import (
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
)

type Logger interface {
	Level() LogLevel
	SetLevel(LogLevel)
	SetOutput(io.Writer)
	With(fields ...Field) Logger
	Debug(string, ...Field)
	Info(string, ...Field)
	Warn(string, ...Field)
	Error(string, ...Field)
	Fatal(string, ...Field)
	Panic(string, ...Field)
}

type logger struct {
	mu         sync.Mutex
	level      LogLevel
	serializer Serializer
	out        Writer
}

func New(w io.Writer, s Serializer, lv LogLevel) Logger {
	l := &logger{
		level:      lv,
		serializer: s,
		out:        NewWriter(w),
	}

	return l
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
	l.out = NewWriter(w)
}

func (l *logger) With(fields ...Field) Logger {
	return newFieldedLogger(l, fields)
}

func (l *logger) Debug(msg string, fields ...Field) {
	l.log(DebugLevel, msg, fields)
}

func (l *logger) Info(msg string, fields ...Field) {
	l.log(InfoLevel, msg, fields)
}

func (l *logger) Warn(msg string, fields ...Field) {
	l.log(WarnLevel, msg, fields)
}

func (l *logger) Error(msg string, fields ...Field) {
	l.log(ErrorLevel, msg, fields)
}

func (l *logger) Fatal(msg string, fields ...Field) {
	l.log(FatalLevel, msg, fields)
	os.Exit(1)
}

func (l *logger) Panic(msg string, fields ...Field) {
	l.log(PanicLevel, msg, fields)
	panic(msg)
}

func (l *logger) log(lv LogLevel, msg string, fields []Field) {
	if lv < l.Level() {
		return
	}

	e := newEntry(lv, msg, fields)
	defer e.free()
	err := l.serializer.Serialize(e)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Serialize error: %v\n", err)
		return
	}

	_, err = l.write(e.Buffer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Write error: %v\n", err)
		return
	}

	if lv >= FatalLevel {
		l.out.Sync()
	}
}

func (l *logger) write(b []byte) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	n, err := l.out.Write(b)
	return n, err
}

var std = New(os.Stdout, new(LTSVSerializer), DebugLevel)

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
	std.Debug(msg, fields...)
}

func Info(msg string, fields ...Field) {
	std.Info(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	std.Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	std.Error(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	std.Fatal(msg, fields...)
}

func Panic(msg string, fields ...Field) {
	std.Panic(msg, fields...)
}
