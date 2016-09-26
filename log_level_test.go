package logify

import (
	"testing"
)

func TestString__debug(t *testing.T) {
	expected := "debug"
	if lvStr := DebugLevel.String(); lvStr != expected {
		t.Fatalf("Expected %v, but %v", expected, lvStr)
	}
}

func TestString__info(t *testing.T) {
	expected := "info"
	if lvStr := InfoLevel.String(); lvStr != expected {
		t.Fatalf("Expected %v, but %v", expected, lvStr)
	}
}

func TestString__warn(t *testing.T) {
	expected := "warn"
	if lvStr := WarnLevel.String(); lvStr != expected {
		t.Fatalf("Expected %v, but %v", expected, lvStr)
	}
}

func TestString__error(t *testing.T) {
	expected := "error"
	if lvStr := ErrorLevel.String(); lvStr != expected {
		t.Fatalf("Expected %v, but %v", expected, lvStr)
	}
}

func TestString__fatal(t *testing.T) {
	expected := "fatal"
	if lvStr := FatalLevel.String(); lvStr != expected {
		t.Fatalf("Expected %v, but %v", expected, lvStr)
	}
}

func TestString__panic(t *testing.T) {
	expected := "panic"
	if lvStr := PanicLevel.String(); lvStr != expected {
		t.Fatalf("Expected %v, but %v", expected, lvStr)
	}
}

func TestString__unknown(t *testing.T) {
	var unknown LogLevel = 0
	expected := "unknown"
	if lvStr := unknown.String(); lvStr != expected {
		t.Fatalf("Expected %v, but %v", expected, lvStr)
	}
}
