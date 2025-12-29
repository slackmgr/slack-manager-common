package common

type Metrics interface {
	// RegisterCounter registers a counter metric with the given name, help text, and optional labels.
	RegisterCounter(name, help string, labels ...string)

	// RegisterGauge registers a gauge metric with the given name, help text, and optional labels.
	RegisterGauge(name, help string, labels ...string)

	// RegisterHistogram registers a histogram metric with the given name, help text, buckets, and optional labels.
	RegisterHistogram(name, help string, buckets []float64, labels ...string)

	// Add adds the given value to the specified counter metric, with optional label values.
	Add(name string, value float64, labelValues ...string)

	// Inc increments the specified counter metric by 1, with optional label values.
	Inc(name string, labelValues ...string)

	// Set sets the specified gauge metric to the given value, with optional label values.
	Set(name string, value float64, labelValues ...string)

	// Observe records an observation for the specified histogram metric, with optional label values.
	Observe(name string, value float64, labelValues ...string)
}
