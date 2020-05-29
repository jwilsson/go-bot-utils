package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func VerifySecret(request events.APIGatewayProxyRequest, secret string) bool {
	signature, err := hex.DecodeString(strings.TrimPrefix(request.Headers["X-Slack-Signature"], "v0="))
	timestamp := request.Headers["X-Slack-Request-Timestamp"]
	body := request.Body

	if err != nil {
		return false
	}

	message := fmt.Sprintf("v0:%s:%s", timestamp, body)
	hash := hmac.New(sha256.New, []byte(secret))

	hash.Write([]byte(message))

	return hmac.Equal(hash.Sum(nil), []byte(signature))
}
