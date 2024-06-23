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
			Namespace: "request_counter",
			Name:      "aggregate",
			Help:      "Total number of aggregate requests",
		}),
		requestLatencyAggregate: promauto.NewHistogram(prometheus.HistogramOpts{
			Namespace: "request_latency",
			Name:      "aggregate",
			Help:      "Request latency aggregate in milliseconds",
		}),
		requestCounterInvoice: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "reuest_counter",
			Name:      "invoice",
			Help:      "Total number of invoice requests",
		}),
		requestLatencyInvoice: promauto.NewHistogram(prometheus.HistogramOpts{
			Namespace: "request_latency",
			Name:      "invoice",
			Help:      "Request latency invoice in milliseconds",
		}),
	}
}

func (a *AggregatorMetricsMiddleware) Aggregate(distance types.Distance) error {
	defer func(begin time.Time) {
		a.requestCounterAggregate.Inc()
		a.requestLatencyAggregate.Observe(float64(time.Since(begin).Microseconds()) / 1000)
	}(time.Now())
	return a.next.Aggregate(distance)
}

func (a *AggregatorMetricsMiddleware) GetInvoice(obuID string) (types.Invoice, error) {
	defer func(begin time.Time) {
		a.requestCounterInvoice.Inc()
		a.requestLatencyInvoice.Observe(float64(time.Since(begin).Microseconds()) / 1000)
	}(time.Now())
	return a.next.GetInvoice(obuID)
}
