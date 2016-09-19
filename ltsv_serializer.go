package logify

import (
	"fmt"
	"time"
)

type LTSVSerializer struct{}

func (f *LTSVSerializer) Serialize(e *Entry, fields []Field) error {
	fmt.Fprintf(
		e.Buffer,
		"level:%s\ttime:%s\tmessage:%s",
		e.Level.String(),
		e.Time.Format(time.RFC3339),
		e.Message,
	)
	if len(fields) > 0 {
		b := make([]byte, 0, 100)
		for _, field := range fields {
			b = append(b, '\t')
			b = append(b, field.key...)
			b = append(b, ':')
			b = append(b, fmt.Sprint(field.value)...)
		}
		e.Buffer.Write(b)
	}
	return nil
}
