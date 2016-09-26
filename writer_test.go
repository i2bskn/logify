package logify

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"testing"
)

func TestNewWriter__file(t *testing.T) {
	logifyWriter := NewWriter(os.Stdout)
	switch logifyWriter.(type) {
	case Writer:
	default:
		t.Fatalf("Expected Writer object, but %v", logifyWriter)
	}
}

func TestNewWriter__bufio(t *testing.T) {
	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)
	logifyWriter := NewWriter(writer)
	switch logifyWriter.(type) {
	case Writer:
	default:
		t.Fatalf("Expected logify.Writer object, but %v", logifyWriter)
	}
}

func TestNewWriter__multiWriter(t *testing.T) {
	buf := new(bytes.Buffer)
	multiWriter := io.MultiWriter(buf, os.Stdout)
	logifyWriter := NewWriter(multiWriter)
	switch logifyWriter.(type) {
	case Writer:
	default:
		t.Fatalf("Expected logify.Writer object, but %v", logifyWriter)
	}
}
