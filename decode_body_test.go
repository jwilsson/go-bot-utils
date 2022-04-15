package utils

import (
	"encoding/base64"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestDecodeBody(t *testing.T) {
	expected := "hello"
	request := events.LambdaFunctionURLRequest{
		Body: base64.StdEncoding.EncodeToString([]byte(expected)),
		IsBase64Encoded: true,
	}

	got := decodeBody(request)
	if got != expected {
		t.Fatalf("Expected '%s', but got '%s'", expected, got)
	}
}
