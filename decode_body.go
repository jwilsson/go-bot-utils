package utils

import (
	"encoding/base64"

	"github.com/aws/aws-lambda-go/events"
)

func decodeBody(r events.LambdaFunctionURLRequest) string {
	if !r.IsBase64Encoded {
		return r.Body
	}

	result, err := base64.StdEncoding.DecodeString(r.Body)
	if err != nil {
		return ""
	}

	return string(result)
}
