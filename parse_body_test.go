package utils

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/slack-go/slack"
)

func TestParseBody(t *testing.T) {
	expected := &slack.SlashCommand{
		ChannelID:      "C1234ABCD",
		ChannelName:    "channel",
		Command:        "/command",
		EnterpriseID:   "E0001",
		EnterpriseName: "Globular%20Construct%20Inc",
		ResponseURL:    "https://hooks.slack.com/commands/XXXXXXXX/00000000000/YYYYYYYYYYYYYY",
		TeamDomain:     "team",
		TeamID:         "T1234ABCD",
		Text:           "text",
		Token:          "valid",
		TriggerID:      "0000000000.1111111111.222222222222aaaaaaaaaaaaaa",
		UserID:         "U1234ABCD",
		UserName:       "username",
	}

	values := url.Values{
		"channel_id":      []string{"C1234ABCD"},
		"channel_name":    []string{"channel"},
		"command":         []string{"/command"},
		"enterprise_id":   []string{"E0001"},
		"enterprise_name": []string{"Globular%20Construct%20Inc"},
		"response_url":    []string{"https://hooks.slack.com/commands/XXXXXXXX/00000000000/YYYYYYYYYYYYYY"},
		"team_domain":     []string{"team"},
		"team_id":         []string{"T1234ABCD"},
		"text":            []string{"text"},
		"token":           []string{"valid"},
		"trigger_id":      []string{"0000000000.1111111111.222222222222aaaaaaaaaaaaaa"},
		"user_id":         []string{"U1234ABCD"},
		"user_name":       []string{"username"},
	}

	got, _ := ParseBody(values.Encode())
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, got %v", got, expected)
	}
}
