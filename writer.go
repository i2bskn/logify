package logify

import (
	"io"
)

type Writer interface {
	io.Writer
	Flush() error
}

type syncableWriter interface {
	io.Writer
	Sync() error
}

func NewWriter(w io.Writer) Writer {
	switch w := w.(type) {
	case Writer:
		return w
	case syncableWriter:
		return syncWriter{baseWriter: w}
	default:
		return defaultWriter{baseWriter: w}
	}
}

type syncWriter struct {
	baseWriter syncableWriter
}

func (sw syncWriter) Write(b []byte) (int, error) {
	n, err := sw.baseWriter.Write(b)
	return n, err
}

func (sw syncWriter) Flush() error {
	err := sw.baseWriter.Sync()
	return err
}

type defaultWriter struct {
	baseWriter io.Writer
}

func (dw defaultWriter) Write(b []byte) (int, error) {
	n, err := dw.Write(b)
	return n, err
}

func (dw defaultWriter) Flush() error {
	return nil
}
