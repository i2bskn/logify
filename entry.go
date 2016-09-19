package logify

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

type Entry struct {
	Logger  Logger
	Time    time.Time
	Level   LogLevel
	Message string
	Buffer  *bytes.Buffer
}

func newEntry(l Logger) *Entry {
	return &Entry{
		Logger: l,
		Buffer: new(bytes.Buffer),
	}
}

func (e *Entry) Reset() {
	e.Buffer.Reset()
}

func (e *Entry) log(level LogLevel, msg string, fields []Field) {
	e.Time = time.Now()
	e.Level = level
	e.Message = msg

	err := e.Logger.Serializer().Serialize(e, fields)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Serialize error: %v\n", err)
		return
	}

	e.Logger.Lock()
	defer e.Logger.Unlock()

	fmt.Fprint(e.Buffer, "\n")
	_, err = e.Logger.Write(e.Buffer.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Write error: %v\n", err)
	}
}

func (e *Entry) Debug(msg string, fields ...Field) {
	if LevelDebug >= e.Logger.Level() {
		e.log(LevelDebug, msg, fields)
	}
}

func (e *Entry) Info(msg string, fields ...Field) {
	if LevelInfo >= e.Logger.Level() {
		e.log(LevelInfo, msg, fields)
	}
}

func (e *Entry) Warn(msg string, fields ...Field) {
	if LevelWarn >= e.Logger.Level() {
		e.log(LevelWarn, msg, fields)
	}
}

func (e *Entry) Error(msg string, fields ...Field) {
	if LevelError >= e.Logger.Level() {
		e.log(LevelError, msg, fields)
	}
}

func (e *Entry) Fatal(msg string, fields ...Field) {
	if LevelFatal >= e.Logger.Level() {
		e.log(LevelFatal, msg, fields)
	}
}

func (e *Entry) Panic(msg string, fields ...Field) {
	if LevelPanic >= e.Logger.Level() {
		e.log(LevelPanic, msg, fields)
	}
}
