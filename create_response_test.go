package utils

import (
	"testing"
)

func TestCreateResponse(t *testing.T) {
	expected := 200

	got := CreateResponse(expected)
	if got.StatusCode != expected {
		t.Fatalf("Expected %d, but got %d", expected, got.StatusCode)
	}
}
