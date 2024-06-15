package storage

import (
	"fmt"

	"github.com/fabrizioperria/toll/shared/types"
)

type MapStorage struct {
	storage map[int]float64
}

func NewMapStorage() Storer {
	return &MapStorage{
		storage: make(map[int]float64),
	}
}

func (m *MapStorage) Store(distance types.Distance) error {
	m.storage[distance.OBUID] += distance.Value
	return nil
}

func (m *MapStorage) Get(obuID int) (float64, error) {
	value, ok := m.storage[obuID]
	if !ok {
		return -1, fmt.Errorf("obuID %d not found", obuID)
	}
	return value, nil
}
