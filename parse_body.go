package utils

import (
	"net/http"
	"net/url"

	"github.com/slack-go/slack"
)

func ParseBody(body string) (*slack.SlashCommand, error) {
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
