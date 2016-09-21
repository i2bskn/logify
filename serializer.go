package logify

import (
	"fmt"
	"time"
)

type Serializer interface {
	Serialize(*Entry) error
}

type LTSVSerializer struct{}

func (f *LTSVSerializer) Serialize(e *Entry) error {
	e.Buffer = append(e.Buffer, "level:"...)
	e.Buffer = append(e.Buffer, e.Level.String()...)
	e.Buffer = append(e.Buffer, "\ttime:"...)
	e.Buffer = append(e.Buffer, e.Time.Format(time.RFC3339)...)
	e.Buffer = append(e.Buffer, "\tmessage:"...)
	e.Buffer = append(e.Buffer, e.Message...)

	if len(e.Fields) > 0 {
		for _, field := range e.Fields {
			e.Buffer = append(e.Buffer, '\t')
			e.Buffer = append(e.Buffer, field.key...)
			e.Buffer = append(e.Buffer, ':')
			e.Buffer = append(e.Buffer, fmt.Sprint(field.value)...)
		}
	}
	e.Buffer = append(e.Buffer, '\n')
	return nil
}
