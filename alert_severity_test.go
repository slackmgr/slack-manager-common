package types_test

import (
	"testing"

	"github.com/slackmgr/types"
	"github.com/stretchr/testify/assert"
)

func TestAlertSeverityValidation(t *testing.T) {
	t.Parallel()

	assert.True(t, types.SeverityIsValid(types.AlertPanic))
	assert.True(t, types.SeverityIsValid(types.AlertError))
	assert.True(t, types.SeverityIsValid(types.AlertWarning))
	assert.True(t, types.SeverityIsValid(types.AlertResolved))
	assert.True(t, types.SeverityIsValid(types.AlertInfo))
	assert.False(t, types.SeverityIsValid("invalid"))
}

func TestAlertPriority(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 3, types.SeverityPriority(types.AlertPanic))
	assert.Equal(t, 2, types.SeverityPriority(types.AlertError))
	assert.Equal(t, 1, types.SeverityPriority(types.AlertWarning))
	assert.Equal(t, 0, types.SeverityPriority(types.AlertResolved))
	assert.Equal(t, 0, types.SeverityPriority(types.AlertInfo))
	assert.Equal(t, -1, types.SeverityPriority("invalid"))
}

func TestValidSeverities(t *testing.T) {
	t.Parallel()

	s := types.ValidSeverities()
	assert.Len(t, s, 5)
	assert.Contains(t, s, "panic")
	assert.Contains(t, s, "error")
	assert.Contains(t, s, "warning")
	assert.Contains(t, s, "resolved")
	assert.Contains(t, s, "info")
}
