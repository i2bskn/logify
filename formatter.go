package logo

import (
	"strings"
	"time"
)

type (
	Formatter interface {
		Format(*Entry) ([]byte, error)
	}

	LTSVFormatter struct{}
)

func (f *LTSVFormatter) Format(e *Entry) ([]byte, error) {
	items := make([]string, 0, 0)
	items = append(items, strings.Join([]string{"level", string(e.Level)}, ":"))
	items = append(items, strings.Join([]string{"time", e.Time.Format(time.RFC3339)}, ":"))
	items = append(items, strings.Join([]string{"message", e.Message}, ":"))
	return []byte(strings.Join(items, "\t") + "\n"), nil
}
