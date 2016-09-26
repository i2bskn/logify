package logify

import (
	"testing"
)

func TestNewEntry(t *testing.T) {
	expectedMessage := "example message"
	entry := newEntry(InfoLevel, expectedMessage, []Field{})
	if entry == nil {
		t.Fatal("Expected Entry object, but nil")
	}

	if entry.Level != InfoLevel {
		t.Fatalf("Expected %v, but %v", InfoLevel, entry.Level)
	}

	if entry.Message != expectedMessage {
		t.Fatalf("Expected %v, but %v", expectedMessage, entry.Message)
	}
}
