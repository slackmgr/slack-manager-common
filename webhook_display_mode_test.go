package types_test

import (
	"testing"

	"github.com/slackmgr/types"
	"github.com/stretchr/testify/assert"
)

func TestWebhookDisplayMode(t *testing.T) {
	t.Parallel()

	assert.True(t, types.WebhookDisplayModeIsValid(types.WebhookDisplayModeAlways))
	assert.True(t, types.WebhookDisplayModeIsValid(types.WebhookDisplayModeOpenIssue))
	assert.True(t, types.WebhookDisplayModeIsValid(types.WebhookDisplayModeResolvedIssue))
	assert.False(t, types.WebhookDisplayModeIsValid("invalid"))
}

func TestWebhookDisplayModeString(t *testing.T) {
	t.Parallel()

	s := types.ValidWebhookDisplayModes()
	assert.Len(t, s, 3)
	assert.Contains(t, s, "always")
	assert.Contains(t, s, "open_issue")
	assert.Contains(t, s, "resolved_issue")
}
