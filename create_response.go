package utils

import (
	"github.com/aws/aws-lambda-go/events"
)

func CreateResponse(statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       "",
		Headers:    map[string]string{},
		StatusCode: statusCode,
	}
}
