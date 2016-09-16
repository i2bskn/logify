package logo

import (
	"fmt"
	"os"
	"time"
)

type Entry struct {
	Logger  *Logger
	Time    time.Time
	Level   Level
	Message string
}

func newEntry(logger *Logger) *Entry {
	return &Entry{
		Logger: logger,
	}
}

func (e *Entry) log(level Level, msg string) {
	e.Time = time.Now()
	e.Level = level
	e.Message = msg
	b, err := e.Logger.formatter.Format(e)

	e.Logger.mu.Lock()
	defer e.Logger.mu.Unlock()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Format error: %v\n", err)
		return
	}

	_, err = e.Logger.out.Write(b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Write error: %v\n", err)
	}
}

func (e *Entry) Debug(v ...interface{}) {
	if LevelDebug >= e.Logger.level {
		e.log(LevelDebug, fmt.Sprint(v...))
	}
}
