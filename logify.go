package logify

type LogLevel int32

const (
	LevelDebug LogLevel = iota + 1
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (level LogLevel) String() string {
	switch level {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	default:
		return "unknown"
	}
}
