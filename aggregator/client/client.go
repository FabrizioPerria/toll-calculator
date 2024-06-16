package client

import "github.com/fabrizioperria/toll/shared/types"

type AggregatorClient interface {
	Aggregate(distance types.Distance) error
	Invoice(obuID string) (types.Invoice, error)
}
