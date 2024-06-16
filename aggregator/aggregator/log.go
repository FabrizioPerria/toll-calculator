package aggregator

import (
	"log"

	"github.com/fabrizioperria/toll/shared/types"
)

type AggregatorLogMiddleware struct {
	next Aggregator
}

func NewAggregatorLogMiddleware(next Aggregator) Aggregator {
	return &AggregatorLogMiddleware{
		next: next,
	}
}

func (l *AggregatorLogMiddleware) Aggregate(distance types.Distance) error {
	log.Printf("STORE: %+v\n", distance)
	return l.next.Aggregate(distance)
}

func (l *AggregatorLogMiddleware) GetInvoice(obuID string) (types.Invoice, error) {
	log.Printf("INVOICE: %+v\n", obuID)
	return l.next.GetInvoice(obuID)
}
