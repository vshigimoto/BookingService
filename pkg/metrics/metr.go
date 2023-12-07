package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HttpBookTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "book_total",
			Help: "The total number of booking room",
		})
)
