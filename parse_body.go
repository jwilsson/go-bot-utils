package utils

import (
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/slack-go/slack"
)

func ParseBody(r events.LambdaFunctionURLRequest) (s slack.SlashCommand, err error) {
	body, err := decodeBody(r)
	if err != nil {
		return s, err
	}

	values, err := url.ParseQuery(body)
	if err != nil {
		return s, err
	}

	return slack.SlashCommandParse(&http.Request{
		PostForm: values,
	})
}
