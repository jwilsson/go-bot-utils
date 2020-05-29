package utils

import (
	"net/url"

	"github.com/slack-go/slack"
)


func ParseBody(body string) (*slack.SlashCommand, error) {
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
