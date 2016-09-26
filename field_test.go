package logify

import (
	"testing"
)

const key = "test"

func TestBool(t *testing.T) {
	expected := true
	field := Bool(key, expected)
	if field.key != key {
		t.Fatalf("Expected %v, but %v", key, field.key)
	}

	if field.value.(bool) != expected {
		t.Fatalf("Expected %v, but %v", expected, field.value.(bool))
	}

	if field.fieldType != boolType {
		t.Fatalf("Expected %v, but %v", boolType, field.fieldType)
	}
}

func TestInt(t *testing.T) {
	expected := 1
	field := Int(key, expected)
	if field.key != key {
		t.Fatalf("Expected %v, but %v", key, field.key)
	}

	if field.value.(int) != expected {
		t.Fatalf("Expected %v, but %v", expected, field.value.(int))
	}

	if field.fieldType != intType {
		t.Fatalf("Expected %v, but %v", intType, field.fieldType)
	}
}

func TestString(t *testing.T) {
	expected := "example"
	field := String(key, expected)
	if field.key != key {
		t.Fatalf("Expected %v, but %v", key, field.key)
	}

	if field.value.(string) != expected {
		t.Fatalf("Expected %v, but %v", expected, field.value.(string))
	}

	if field.fieldType != stringType {
		t.Fatalf("Expected %v, but %v", stringType, field.fieldType)
	}
}
