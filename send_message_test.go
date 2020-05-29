package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/slack-go/slack"
)

func TestSendMessage(t *testing.T) {
	expected := "Hello, world"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var message slack.Message

		json.NewDecoder(r.Body).Decode(&message)

		if message.Text != expected {
			t.Fatalf("Expected %s, got %s", expected, message.Text)
		}
	}))

	defer ts.Close()

	SendMessage(ts.URL, slack.Message{
		Msg: slack.Msg{
			Text: expected,
		},
	})
}
