package utils

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/slack-go/slack"
)

func SendMessage(url string, message slack.Message) error {
	body, _ := json.Marshal(message)

	_, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)

	return err
}
