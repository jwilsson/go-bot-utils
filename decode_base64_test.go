package utils

import (
	"testing"
)

func TestDecodeBase64(t *testing.T) {
	expected := "hello"
	data := "aGVsbG8="

	got := decodeBase64(data)
	if got != expected {
		t.Fatalf("Expected '%s', but got '%s'", expected, got)
	}
}
