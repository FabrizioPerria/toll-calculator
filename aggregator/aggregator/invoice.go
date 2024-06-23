package aggregator

import (
	"os"
	"time"

	"github.com/fabrizioperria/toll/aggregator/storage"
	"github.com/fabrizioperria/toll/shared/types"
)

type InvoiceAggregator struct {
	store storage.Storer
}

func NewInvoiceAggregator() Aggregator {
	var agg Aggregator = &InvoiceAggregator{store: storage.AggregatorStorageFactory(os.Getenv("STORAGE_TYPE"))}
	agg = NewAggregatorMetricsMiddleware(agg)
	agg = NewAggregatorLogMiddleware(agg)
	return agg
}

func (a *InvoiceAggregator) Aggregate(distance types.Distance) error {
	return a.store.Store(distance)
}

const pricePerKm = 0.1

func (a *InvoiceAggregator) GetInvoice(obuID string) (types.Invoice, error) {
	distance, err := a.store.Get(obuID)
	if err != nil {
		return types.Invoice{}, err
	}
	return types.Invoice{
		ObuId:     obuID,
		Amount:    pricePerKm * distance,
		Distance:  distance,
		Timestamp: time.Now().Unix(),
	}, nil
}
