package utils

import (
	"encoding/base64"

	"github.com/aws/aws-lambda-go/events"
)

func decodeBody(r events.LambdaFunctionURLRequest) (string, error) {
	if !r.IsBase64Encoded {
		return r.Body, nil
	}

	result, err := base64.StdEncoding.DecodeString(r.Body)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
