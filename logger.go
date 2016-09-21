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

type coreLogger struct {
	mu         sync.Mutex
	level      LogLevel
	serializer Serializer
	out        io.Writer
}

func New(w io.Writer, s Serializer, lv LogLevel) Logger {
	l := &coreLogger{
		level:      lv,
		serializer: s,
		out:        w,
	}

	return l
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

func (cl *coreLogger) With(fields ...Field) Logger {
	return newFieldedLogger(cl, fields)
}

func (cl *coreLogger) Debug(msg string, fields ...Field) {
	cl.log(DebugLevel, msg, fields)
}

func (cl *coreLogger) Info(msg string, fields ...Field) {
	cl.log(InfoLevel, msg, fields)
}

func (cl *coreLogger) Warn(msg string, fields ...Field) {
	cl.log(WarnLevel, msg, fields)
}

func (cl *coreLogger) Error(msg string, fields ...Field) {
	cl.log(ErrorLevel, msg, fields)
}

func (cl *coreLogger) Fatal(msg string, fields ...Field) {
	cl.log(FatalLevel, msg, fields)
	os.Exit(1)
}

func (cl *coreLogger) Panic(msg string, fields ...Field) {
	cl.log(PanicLevel, msg, fields)
	panic(msg)
}

func (cl *coreLogger) log(lv LogLevel, msg string, fields []Field) {
	if lv < cl.Level() {
		return
	}

	e := newEntry(lv, msg, fields)
	defer e.free()
	err := cl.serializer.Serialize(e)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Serialize error: %v\n", err)
		return
	}

	_, err = cl.write(e.Buffer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Write error: %v\n", err)
		return
	}
}

func (cl *coreLogger) write(b []byte) (int, error) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	n, err := cl.out.Write(b)
	return n, err
}

var defaultLogger = New(os.Stdout, new(LTSVSerializer), DebugLevel)

func Level() LogLevel {
	return defaultLogger.Level()
}

func SetLevel(lv LogLevel) {
	defaultLogger.SetLevel(lv)
}

func SetOutput(w io.Writer) {
	defaultLogger.SetOutput(w)
}

func Debug(msg string, fields ...Field) {
	defaultLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...Field) {
	defaultLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	defaultLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	defaultLogger.Error(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	defaultLogger.Fatal(msg, fields...)
}

func Panic(msg string, fields ...Field) {
	defaultLogger.Panic(msg, fields...)
}
