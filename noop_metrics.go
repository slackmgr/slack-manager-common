package types

// NoopMetrics is a no-op implementation of the Metrics interface.
// It does not record any metrics. Use if no metrics are needed.
type NoopMetrics struct{}

func (m *NoopMetrics) RegisterCounter(_, _ string, _ ...string) {
}

func (m *NoopMetrics) RegisterGauge(_, _ string, _ ...string) {
}

func (m *NoopMetrics) RegisterHistogram(_, _ string, _ []float64, _ ...string) {
}

func (m *NoopMetrics) Add(_ string, _ float64, _ ...string) {
}

func (m *NoopMetrics) Inc(_ string, _ ...string) {
}

func (m *NoopMetrics) Set(_ string, _ float64, _ ...string) {
}

func (m *NoopMetrics) Observe(_ string, _ float64, _ ...string) {
}
