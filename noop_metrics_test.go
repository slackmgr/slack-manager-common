package types_test

import (
	"testing"

	"github.com/slackmgr/types"
)

func TestNoopMetrics(t *testing.T) {
	t.Parallel()

	// Ensure NoopMetrics implements the Metrics interface
	var m types.Metrics = &types.NoopMetrics{}

	// Ensure methods can be called without errors or panics
	m.RegisterCounter("", "")
	m.RegisterGauge("", "")
	m.RegisterHistogram("", "", []float64{})
	m.Add("", 0)
	m.Inc("", "")
	m.Set("", 0)
	m.Observe("", 0)
}
