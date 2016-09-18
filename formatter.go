package logify

import (
	"fmt"
	"time"
)

type (
	Formatter interface {
		Format(*Entry, []Field) error
	}

	LTSVFormatter struct{}
)

func (f *LTSVFormatter) Format(e *Entry, fields []Field) error {
	fmt.Fprintf(
		e.Buffer,
		"level:%s\ttime:%s\tmessage:%s",
		e.Level.String(),
		e.Time.Format(time.RFC3339),
		e.Message,
	)
	for _, field := range fields {
		fmt.Fprintf(e.Buffer, "\t%v:%v", field.key, field.value)
	}
	return nil
}
