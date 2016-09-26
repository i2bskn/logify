package logify

import (
	"io"
)

type Writer interface {
	io.Writer
	Sync() error
}

type flushableWriter interface {
	io.Writer
	Flush() error
}

func NewWriter(w io.Writer) Writer {
	switch w := w.(type) {
	case Writer:
		return w
	case flushableWriter:
		return flushWriter{w}
	default:
		return defaultWriter{w}
	}
}

type flushWriter struct {
	flushableWriter
}

func (sw flushWriter) Sync() error {
	return sw.Flush()
}

type defaultWriter struct {
	io.Writer
}

func (dw defaultWriter) Sync() error {
	return nil
}
