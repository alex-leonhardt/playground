package metrics

import (
	"fmt"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

// NewPusher returns a new metric pusher
func NewPusher(job string) *push.Pusher {
	pusher := push.New("http://localhost:9091", job)
	return pusher
}

// newGauge returns a new gauge with name and labels applied
func newGauge(name string, help string, labels prometheus.Labels) prometheus.Gauge {
	g := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        name,
		Help:        help,
		ConstLabels: labels,
	})
	return g
}

func sanitize(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), "/", "_")
}

// WrappedTestWithGaugeMetric decorates a testing Func with a prometheus gauge metric
func WrappedTestWithGaugeMetric(t *testing.T, f func(t *testing.T)) func(*testing.T) {

	return func(t *testing.T) {
		pusher := NewPusher(sanitize(t.Name()))

		g := newGauge(
			sanitize(t.Name()),
			t.Name(),
			prometheus.Labels{},
		)

		f(t)

		g.SetToCurrentTime()
		pusher.Collector(g)

		if err := pusher.Push(); err != nil {
			fmt.Println("Could not push to prom pushgateway:", err)
		}

	}
}
