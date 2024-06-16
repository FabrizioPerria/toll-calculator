package aggregator

import (
	"github.com/fabrizioperria/toll/shared/types"
)

type Aggregator interface {
	Aggregate(types.Distance) error
	GetInvoice(string) (types.Invoice, error)
}
