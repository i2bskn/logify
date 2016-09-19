package logify

import (
	"io"
)

type fieldedLogger struct {
	logger Logger
	fields []Field
}

func newFieldedLogger(l Logger, fields []Field) Logger {
	return &fieldedLogger{
		logger: l,
		fields: fields,
	}
}

func (fl *fieldedLogger) Level() LogLevel {
	return fl.logger.Level()
}

func (fl *fieldedLogger) SetLevel(lv LogLevel) {
	fl.logger.SetLevel(lv)
}

func (fl *fieldedLogger) SetOutput(w io.Writer) {
	fl.logger.SetOutput(w)
}

func (fl *fieldedLogger) Lock() {
	fl.logger.Lock()
}

func (fl *fieldedLogger) Unlock() {
	fl.logger.Unlock()
}

func (fl *fieldedLogger) Write(b []byte) (int, error) {
	n, err := fl.logger.Write(b)
	return n, err
}

func (fl *fieldedLogger) Serializer() Serializer {
	return fl.logger.Serializer()
}

func (fl *fieldedLogger) With(fields ...Field) Logger {
	return fl.logger.With(fields...)
}

func (fl *fieldedLogger) Debug(msg string, fields ...Field) {
	fl.logger.Debug(msg, fl.fieldJoin(fields)...)
}

func (fl *fieldedLogger) Info(msg string, fields ...Field) {
	fl.logger.Info(msg, fl.fieldJoin(fields)...)
}

func (fl *fieldedLogger) Warn(msg string, fields ...Field) {
	fl.logger.Warn(msg, fl.fieldJoin(fields)...)
}

func (fl *fieldedLogger) Error(msg string, fields ...Field) {
	fl.logger.Error(msg, fl.fieldJoin(fields)...)
}

func (fl *fieldedLogger) Fatal(msg string, fields ...Field) {
	fl.logger.Fatal(msg, fl.fieldJoin(fields)...)
}

func (fl *fieldedLogger) Panic(msg string, fields ...Field) {
	fl.logger.Panic(msg, fl.fieldJoin(fields)...)
}

func (fl *fieldedLogger) fieldJoin(fields []Field) []Field {
	return append(fl.fields, fields...)
}
