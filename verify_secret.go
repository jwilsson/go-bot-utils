package utils

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/slack-go/slack"
)

func VerifySecret(r events.APIGatewayProxyRequest, secret string) bool {
	verifier, err := slack.NewSecretsVerifier(r.MultiValueHeaders, secret)
	if err != nil {
		return false
	}

	_, err = verifier.Write([]byte(r.Body))
	if err != nil {
		return false
	}

	return verifier.Ensure() == nil
}
