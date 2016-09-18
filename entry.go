package logify

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

type Entry struct {
	Logger  *Logger
	Time    time.Time
	Level   LogLevel
	Message string
	Buffer  *bytes.Buffer
}

func newEntry(logger *Logger) *Entry {
	return &Entry{
		Logger: logger,
		Buffer: new(bytes.Buffer),
	}
}

func (e *Entry) Free() {
	e.Buffer.Reset()
	e.Logger.entryPool.Put(e)
}

func (e *Entry) log(level LogLevel, msg string, fields []Field) {
	e.Time = time.Now()
	e.Level = level
	e.Message = msg

	err := e.Logger.formatter.Format(e, fields)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Format error: %v\n", err)
		return
	}

	e.Logger.mu.Lock()
	defer e.Logger.mu.Unlock()

	fmt.Fprint(e.Buffer, "\n")
	_, err = e.Logger.out.Write(e.Buffer.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Write error: %v\n", err)
	}
}

func (e *Entry) Debug(msg string, fields ...Field) {
	if LevelDebug >= e.Logger.level {
		e.log(LevelDebug, msg, fields)
	}
}
