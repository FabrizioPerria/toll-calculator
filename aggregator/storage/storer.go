package storage

import "github.com/fabrizioperria/toll/shared/types"

type Storer interface {
	Store(types.Distance) error
	Get(string) (float64, error)
}
