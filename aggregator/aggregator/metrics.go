package aggregator

import (
	"time"

	"github.com/fabrizioperria/toll/shared/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type AggregatorMetricsMiddleware struct {
	next                    Aggregator
	requestCounterAggregate prometheus.Counter
	requestLatencyAggregate prometheus.Histogram
	requestCounterInvoice   prometheus.Counter
	requestLatencyInvoice   prometheus.Histogram
}

func NewAggregatorMetricsMiddleware(next Aggregator) Aggregator {
	return &AggregatorMetricsMiddleware{
		next: next,
		requestCounterAggregate: promauto.NewCounter(prometheus.CounterOpts{
			// Namespace: "aggregate_requestCounter",
			Name: "aggregate_cnt",
			Help: "Total number of requests",
		}),
		requestLatencyAggregate: promauto.NewHistogram(prometheus.HistogramOpts{
			// Namespace: "aggregate_requestLatency",
			Name:    "aggregate_lat",
			Help:    "Request latency in milliseconds",
			Buckets: []float64{0.1, 0.5, 1},
		}),
		requestCounterInvoice: promauto.NewCounter(prometheus.CounterOpts{
			// Namespace: "invoice_requestCounter",
			Name: "invoice_cnt",
			Help: "Total number of requests",
		}),
		requestLatencyInvoice: promauto.NewHistogram(prometheus.HistogramOpts{
			// Namespace: "invoice_requestLatency",
			Name:    "invoice_lat",
			Help:    "Request latency in milliseconds",
			Buckets: []float64{0.1, 0.5, 1},
		}),
	}
}

func (a *AggregatorMetricsMiddleware) Aggregate(distance types.Distance) error {
	defer func(begin time.Time) {
		a.requestCounterAggregate.Inc()
		a.requestLatencyAggregate.Observe(float64(time.Since(begin).Seconds()))
	}(time.Now())
	return a.next.Aggregate(distance)
}

func (a *AggregatorMetricsMiddleware) GetInvoice(obuID string) (types.Invoice, error) {
	defer func(begin time.Time) {
		a.requestCounterInvoice.Inc()
		a.requestLatencyInvoice.Observe(float64(time.Since(begin).Seconds()))
	}(time.Now())
	return a.next.GetInvoice(obuID)
}
