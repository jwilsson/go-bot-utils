package utils

import (
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/slack-go/slack"
)

func ParseBody(r events.LambdaFunctionURLRequest) (*slack.SlashCommand, error) {
	body := decodeBody(r)
	values, err := url.ParseQuery(body)
	if err != nil {
		return nil, err
	}

	s, err := slack.SlashCommandParse(&http.Request{
		PostForm: values,
	})

	if err != nil {
		return nil, err
	}

	return &s, nil
}
