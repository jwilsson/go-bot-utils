package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

func TestVerifySecret(t *testing.T) {
	secret := "my secret"
	body := []byte("command=foo&text=bar")
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	message := fmt.Sprintf("v0:%s:%s", timestamp, body)
	hash := hmac.New(sha256.New, []byte(secret))

	hash.Write([]byte(message))

	signature := fmt.Sprintf("%x", hash.Sum(nil))

	request := events.LambdaFunctionURLRequest{
		Body: base64.StdEncoding.EncodeToString(body),
		Headers: map[string]string{
			"X-Slack-Request-Timestamp": timestamp,
			"X-Slack-Signature":         "v0=" + signature,
		},
	}

	got := VerifySecret(request, secret)
	if !got {
		t.Fatal("Expected true, but got false")
	}
}
