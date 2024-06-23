package storage

import (
	"os"

	"github.com/fabrizioperria/toll/shared/types"
)

type Storer interface {
	Store(types.Distance) error
	Get(string) (float64, error)
}

func AggregatorStorageFactory(storageType string) Storer {
	switch storageType {
	case "map":
		return NewMapStorage()
	case "mongo":
		return NewMongoStorage(os.Getenv("MONGO_URI"))
	default:
		panic("Unknown storage type " + storageType)
	}
}
