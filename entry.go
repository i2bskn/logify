package logify

import (
	"sync"
	"time"
)

const entryBufferSize = 512

var entryPool = sync.Pool{
	New: func() interface{} {
		return &Entry{
			Buffer: make([]byte, 0, entryBufferSize),
		}
	},
}

type Entry struct {
	Level   LogLevel
	Time    time.Time
	Message string
	Fields  []Field
	Buffer  []byte
}

func newEntry(lv LogLevel, msg string, fields []Field) *Entry {
	e := entryPool.Get().(*Entry)
	e.Level = lv
	e.Time = time.Now()
	e.Message = msg
	e.Fields = fields
	return e
}

func (e *Entry) free() {
	e.Buffer = e.Buffer[:0]
}
