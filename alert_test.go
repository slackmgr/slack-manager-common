package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlackChannelIDValidation(t *testing.T) {
	a := Alert{SlackChannelID: "abcdefghi"}
	assert.NoError(t, a.ValidateSlackChannelID())

	a = Alert{SlackChannelID: "ABab129cf"}
	assert.NoError(t, a.ValidateSlackChannelID())

	a = Alert{SlackChannelID: "abcdefghi9238yr"}
	assert.NoError(t, a.ValidateSlackChannelID())

	a = Alert{SlackChannelID: ""}
	assert.Error(t, a.ValidateSlackChannelID())

	// Channel names are allowed
	a = Alert{SlackChannelID: "12345678"}
	assert.NoError(t, a.ValidateSlackChannelID())

	// Channel names are allowed
	a = Alert{SlackChannelID: "foo-something"}
	assert.NoError(t, a.ValidateSlackChannelID())

	a = Alert{SlackChannelID: "sdkjsdf asdfasdf"}
	assert.Error(t, a.ValidateSlackChannelID())

	a = Alert{SlackChannelID: "abcdefghi9238yrh923hr29fh92qu4fh92q4hf9qabcdefghi9238yrh923hr29fh92qu4fh92q4hf9qabcdefghi9238yrh923hr29fh92qu4fh92q4hf9q123sdfsdfsdf"}
	assert.Error(t, a.ValidateSlackChannelID())
}

func TestAlertEscalation(t *testing.T) {
	a := Alert{
		Escalation: []*Escalation{
			{DelaySeconds: 30, Severity: AlertWarning, SlackMentions: []string{"<@foo>"}},
			{DelaySeconds: 60, Severity: AlertError, SlackMentions: []string{"<@bar>"}},
		},
	}

	a.Clean()
	err := a.ValidateEscalation()
	assert.NoError(t, err)
	assert.Equal(t, 30, a.Escalation[0].DelaySeconds)
	assert.Equal(t, "<@foo>", a.Escalation[0].SlackMentions[0])
	assert.Equal(t, 60, a.Escalation[1].DelaySeconds)

	a = Alert{
		Escalation: []*Escalation{
			{DelaySeconds: 60, Severity: AlertWarning, SlackMentions: []string{"<@foo>"}},
			{DelaySeconds: 30, Severity: AlertError, SlackMentions: []string{"<@bar>"}},
		},
	}

	a.Clean()
	err = a.ValidateEscalation()
	assert.NoError(t, err)
	assert.Equal(t, 30, a.Escalation[0].DelaySeconds)
	assert.Equal(t, "<@bar>", a.Escalation[0].SlackMentions[0])
	assert.Equal(t, 60, a.Escalation[1].DelaySeconds)

	a = Alert{
		Escalation: []*Escalation{
			{DelaySeconds: 59, Severity: AlertWarning, SlackMentions: []string{"<@foo>"}},
			{DelaySeconds: 30, Severity: AlertError, SlackMentions: []string{"<@bar>"}},
		},
	}

	a.Clean()
	err = a.ValidateEscalation()
	assert.Error(t, err)
	assert.ErrorContains(t, err, "expected diff >=30")
}
