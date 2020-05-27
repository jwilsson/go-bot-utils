package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/slack-go/slack"
)

func CreateResponse(statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       "",
		Headers:    map[string]string{},
		StatusCode: statusCode,
	}
}

func ParseBody(body string) (*slack.SlashCommand, error) {
	log.Printf("Received body: %v", body)

	values, err := url.ParseQuery(body)
	if err != nil {
		return nil, err
	}

	s := &slack.SlashCommand{
		ChannelID:      values.Get("channel_id"),
		ChannelName:    values.Get("channel_name"),
		Command:        values.Get("command"),
		EnterpriseID:   values.Get("enterprise_id"),
		EnterpriseName: values.Get("enterprise_name"),
		ResponseURL:    values.Get("response_url"),
		TeamDomain:     values.Get("team_domain"),
		TeamID:         values.Get("team_id"),
		Text:           values.Get("text"),
		Token:          values.Get("token"),
		TriggerID:      values.Get("trigger_id"),
		UserID:         values.Get("user_id"),
		UserName:       values.Get("user_name"),
	}

	return s, nil
}

func SendMessage(url string, message slack.Message) error {
	body, _ := json.Marshal(message)

	log.Printf("Posting %s to %s", body, url)

	_, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)

	return err
}

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
