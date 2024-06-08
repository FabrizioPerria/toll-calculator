package producers

import "github.com/fabrizioperria/toll/shared/types"

type DataProducer interface {
	Produce(types.OBUData) error
}
