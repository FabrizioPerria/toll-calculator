package storage

import (
	"fmt"

	"github.com/fabrizioperria/toll/shared/types"
)

type MapStorage struct {
	storage map[string]float64
}

func NewMapStorage() Storer {
	return &MapStorage{
		storage: make(map[string]float64),
	}
}

func (m *MapStorage) Store(distance types.Distance) error {
	m.storage[distance.ObuId] += distance.Value
	return nil
}

func (m *MapStorage) Get(obuID string) (float64, error) {
	value, ok := m.storage[obuID]
	if !ok {
		return -1, fmt.Errorf("obuID %s not found", obuID)
	}
	return value, nil
}
