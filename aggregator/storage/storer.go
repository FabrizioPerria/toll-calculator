package storage

import "github.com/fabrizioperria/toll/shared/types"

type Storer interface {
	Store(types.Distance) error
	Get(int) (float64, error)
}
