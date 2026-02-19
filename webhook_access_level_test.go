package types_test

import (
	"testing"

	"github.com/slackmgr/types"
	"github.com/stretchr/testify/assert"
)

func TestWebhookAccessLevelValidation(t *testing.T) {
	t.Parallel()

	assert.True(t, types.WebhookAccessLevelIsValid(types.WebhookAccessLevelGlobalAdmins))
	assert.True(t, types.WebhookAccessLevelIsValid(types.WebhookAccessLevelChannelAdmins))
	assert.True(t, types.WebhookAccessLevelIsValid(types.WebhookAccessLevelChannelMembers))
	assert.False(t, types.WebhookAccessLevelIsValid("invalid"))
}

func TestWebhookAccessLevelString(t *testing.T) {
	t.Parallel()

	s := types.ValidWebhookAccessLevels()
	assert.Len(t, s, 3)
	assert.Contains(t, s, "global_admins")
	assert.Contains(t, s, "channel_admins")
	assert.Contains(t, s, "channel_members")
}
