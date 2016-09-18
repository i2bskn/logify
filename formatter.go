package logo

import (
	"bytes"
	"fmt"
	"time"
)

type (
	Formatter interface {
		Format(*Entry) ([]byte, error)
	}

	LTSVFormatter struct{}
)

func (f *LTSVFormatter) Format(e *Entry) ([]byte, error) {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "level:%s\ttime:%s\tmessage:%s", e.Level.String(), e.Time.Format(time.RFC3339), e.Message)
	return buf.Bytes(), nil
}
