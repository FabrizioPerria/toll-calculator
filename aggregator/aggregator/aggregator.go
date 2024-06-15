package aggregator

import (
	"github.com/fabrizioperria/toll/shared/types"
)

type Invoice struct {
	ObuID    int     `json:"obu_id"`
	Amount   float64 `json:"amount"`
	Distance float64 `json:"distance"`
}

type Aggregator interface {
	Aggregate(types.Distance) error
	GetInvoice(string) (Invoice, error)
}
