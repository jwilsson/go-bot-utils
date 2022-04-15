package utils

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/slack-go/slack"
)

func VerifySecret(r events.LambdaFunctionURLRequest, secret string) bool {
	// Expected format of headers differ, so we need to fix it
	headers := http.Header{}
	for name, value := range r.Headers {
		headers.Add(name, value)
	}

	verifier, err := slack.NewSecretsVerifier(headers, secret)
	if err != nil {
		return false
	}

	_, err = verifier.Write([]byte(r.Body))
	if err != nil {
		return false
	}

	return verifier.Ensure() == nil
}
