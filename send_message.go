package utils

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/slack-go/slack"
)

func SendMessage(url string, message slack.Message) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)

	return err
}
