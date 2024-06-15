package aggregator

import (
	"strconv"

	"github.com/fabrizioperria/toll/aggregator/storage"
	"github.com/fabrizioperria/toll/shared/types"
)

type InvoiceAggregator struct {
	store storage.Storer
}

func NewInvoiceAggregator() Aggregator {
	return NewAggregatorLogMiddleware(&InvoiceAggregator{
		store: storage.NewMapStorage(),
	})
}

func (a *InvoiceAggregator) Aggregate(distance types.Distance) error {
	return a.store.Store(distance)
}

const pricePerKm = 0.1

func (a *InvoiceAggregator) GetInvoice(obuID string) (Invoice, error) {
	obu, err := strconv.ParseInt(obuID, 10, 64)
	if err != nil {
		return Invoice{}, err
	}
	distance, err := a.store.Get(int(obu))
	if err != nil {
		return Invoice{}, err
	}
	return Invoice{
		ObuID:    int(obu),
		Amount:   pricePerKm * distance,
		Distance: distance,
	}, nil
}
