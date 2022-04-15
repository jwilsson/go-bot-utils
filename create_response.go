package utils

import (
	"github.com/aws/aws-lambda-go/events"
)

func CreateResponse(statusCode int) events.LambdaFunctionURLResponse {
	return events.LambdaFunctionURLResponse{
		Body:       "",
		Headers:    map[string]string{},
		StatusCode: statusCode,
	}
}
