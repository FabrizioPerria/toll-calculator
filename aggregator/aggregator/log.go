package aggregator

import (
	"os"
	"time"

	"github.com/fabrizioperria/toll/shared/logger"
	"github.com/fabrizioperria/toll/shared/types"
	"github.com/sirupsen/logrus"
)

type AggregatorLogMiddleware struct {
	next   Aggregator
	logger *logrus.Logger
}

func NewAggregatorLogMiddleware(next Aggregator) Aggregator {
	return &AggregatorLogMiddleware{
		next:   next,
		logger: logger.LoggerFactory(os.Getenv("LOG_PATH")),
	}
}

func (l *AggregatorLogMiddleware) Aggregate(distance types.Distance) error {
	defer func(t time.Time) {
		l.logger.WithFields(logrus.Fields{
			"latency": float64(time.Since(t).Microseconds()) / 1000,
		}).Info("Latency Aggregating distance")
	}(time.Now())

	l.logger.WithFields(logrus.Fields{
		"obu_id":    distance.ObuId,
		"value":     distance.Value,
		"timestamp": distance.Timestamp,
	}).Info("Aggregating distance")

	err := l.next.Aggregate(distance)
	if err != nil {
		l.logger.WithFields(logrus.Fields{
			"obu_id":    distance.ObuId,
			"value":     distance.Value,
			"timestamp": distance.Timestamp,
		}).Error("Failed to aggregate distance")
	}

	return err
}

func (l *AggregatorLogMiddleware) GetInvoice(obuID string) (types.Invoice, error) {
	defer func(t time.Time) {
		l.logger.WithFields(logrus.Fields{
			"latency": float64(time.Since(t).Microseconds()) / 1000,
		}).Info("Latency Getting invoice")
	}(time.Now())

	l.logger.WithFields(logrus.Fields{
		"obu_id": obuID,
	}).Info("Getting invoice")

	invoice, err := l.next.GetInvoice(obuID)
	if err != nil {
		l.logger.WithFields(logrus.Fields{
			"obu_id": obuID,
		}).Error("Failed to get invoice")
	}

	return invoice, err
}
