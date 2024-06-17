package aggregator

import (
	"fmt"
	"io"
	"os"

	"github.com/fabrizioperria/toll/shared/types"
	"github.com/sirupsen/logrus"
)

type AggregatorLogMiddleware struct {
	next   Aggregator
	logger *logrus.Logger
}

func NewAggregatorLogMiddleware(next Aggregator) Aggregator {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory: ", err)
		return nil
	}
	os.MkdirAll(homeDir+"/log/toll", 0o755)
	f, err := os.OpenFile(homeDir+"/log/toll/aggregator.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o755)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil
	}
	l.SetOutput(io.MultiWriter(os.Stdout, f))

	return &AggregatorLogMiddleware{
		next:   next,
		logger: l,
	}
}

func (l *AggregatorLogMiddleware) Aggregate(distance types.Distance) error {
	l.logger.WithFields(logrus.Fields{
		"obu_id":    distance.ObuId,
		"value":     distance.Value,
		"timestamp": distance.Timestamp,
	}).Info("Aggregating distance")

	return l.next.Aggregate(distance)
}

func (l *AggregatorLogMiddleware) GetInvoice(obuID string) (types.Invoice, error) {
	l.logger.WithFields(logrus.Fields{
		"obu_id": obuID,
	}).Info("Getting invoice")

	return l.next.GetInvoice(obuID)
}
