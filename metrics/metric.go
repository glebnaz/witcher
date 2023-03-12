package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Handler() http.Handler {
	return promhttp.Handler()
}

func MustRegister(cs ...prometheus.Collector) {
	prometheus.MustRegister(cs...)
}

// NewCounter creates a new Counter.
func NewCounter(name, namespace, help string) prometheus.Counter {
	return prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	})
}

// NewCounterVec creates a new CounterVec partitioned by the given label names.
func NewCounterVec(name, namespace, help string, labelNames []string) *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	}, labelNames)
}

// NewGauge creates a new Gauge.
func NewGauge(name, namespace, help string) prometheus.Gauge {
	return prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	})
}

// NewGaugeVec creates a new GaugeVec partitioned by the given label names.
func NewGaugeVec(name, namespace, help string, labelNames []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	}, labelNames)
}

// NewHistogram creates a new Histogram.
func NewHistogram(name, namespace, help string, buckets []float64) prometheus.Histogram {
	return prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
		Buckets:   buckets,
	})
}

// NewHistogramVec creates a new HistogramVec partitioned by the given label names.
func NewHistogramVec(name, namespace, help string, buckets []float64, labelNames []string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
		Buckets:   buckets,
	}, labelNames)
}

// NewSummary creates a new Summary.
func NewSummary(name, namespace, help string) prometheus.Summary {
	return prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace:  namespace,
		Name:       name,
		Help:       help,
		Objectives: SummaryObjectives,
		MaxAge:     prometheus.DefMaxAge,
	})
}

// NewSummaryVec creates a new SummaryVec partitioned by the given label names.
func NewSummaryVec(name, namespace, help string, labelNames []string) *prometheus.SummaryVec {
	return prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace:  namespace,
		Name:       name,
		Help:       help,
		Objectives: SummaryObjectives,
		MaxAge:     prometheus.DefMaxAge,
	}, labelNames)
}

// MustRegisterCounter creates and registers a new Counter.
// Panics if any error occurs.
func MustRegisterCounter(name, namespace, help string) prometheus.Counter {
	c := NewCounter(name, namespace, help)
	MustRegister(c)
	return c
}

// MustRegisterCounterVec creates and registers a new CounterVec.
// Panics if any error occurs.
func MustRegisterCounterVec(name, namespace, help string, labelNames []string) *prometheus.CounterVec {
	c := NewCounterVec(name, namespace, help, labelNames)
	MustRegister(c)
	return c
}

// MustRegisterGauge creates and registers a new Gauge.
// Panics if any error occurs.
func MustRegisterGauge(name, namespace, help string) prometheus.Gauge {
	c := NewGauge(name, namespace, help)
	MustRegister(c)
	return c
}

// MustRegisterGaugeVec creates and registers a new GaugeVec.
// Panics if any error occurs.
func MustRegisterGaugeVec(name, namespace, help string, labelNames []string) *prometheus.GaugeVec {
	c := NewGaugeVec(name, namespace, help, labelNames)
	MustRegister(c)
	return c
}

// MustRegisterHistogram creates and registers a new Histogram.
// Panics if any error occurs.
func MustRegisterHistogram(name, namespace, help string, buckets []float64) prometheus.Histogram {
	c := NewHistogram(name, namespace, help, buckets)
	MustRegister(c)
	return c
}

// MustRegisterHistogramVec creates and registers a new HistogramVec.
// Panics if any error occurs.
func MustRegisterHistogramVec(name, namespace, help string, buckets []float64, labelNames []string) *prometheus.HistogramVec {
	c := NewHistogramVec(name, namespace, help, buckets, labelNames)
	MustRegister(c)
	return c
}

// MustRegisterSummary creates and registers a new Summary.
// Panics if any error occurs.
func MustRegisterSummary(name, namespace, help string) prometheus.Summary {
	c := NewSummary(name, namespace, help)
	MustRegister(c)
	return c
}

// MustRegisterSummaryVec creates and registers a new SummaryVec.
// Panics if any error occurs.
func MustRegisterSummaryVec(name, namespace, help string, labelNames []string) *prometheus.SummaryVec {
	c := NewSummaryVec(name, namespace, help, labelNames)
	MustRegister(c)
	return c
}
