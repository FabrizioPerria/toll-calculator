package consumers

import "github.com/fabrizioperria/toll/shared/types"

type DataConsumer interface {
	Consume() (types.OBUData, error)
}
