package utils

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestVerifySecret(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		Body: "/command=foo&text=bar",
		Headers: map[string]string{
			"X-Slack-Request-Timestamp": "1531420618",
			"X-Slack-Signature": "v0=f1e1684b07f1ff7b0e09b2dcbeb48b461c20cc7743c7c4955fc537a63e077fda",
		},
	}

	got := VerifySecret(request, "my secret")
	if !got {
		t.Fatal("Expected true, but got false")
	}
}
