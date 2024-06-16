package service

import (
	"math"

	"github.com/fabrizioperria/toll/shared/types"
)

type Calculator interface {
	Distance(types.OBUData) (float64, error)
}

type CalculatorService struct{}

func NewCalculatorService() Calculator {
	return NewLogServiceMiddleware(&CalculatorService{})
}

var lastObuData map[string]types.OBUData = make(map[string]types.OBUData, 50)

func (c *CalculatorService) Distance(obuData types.OBUData) (float64, error) {
	lastObuDataForID := lastObuData[obuData.ObuId]
	distance := math.Sqrt(math.Pow(obuData.Latitude-lastObuDataForID.Latitude, 2) + math.Pow(obuData.Longitude-lastObuDataForID.Longitude, 2))
	lastObuData[obuData.ObuId] = obuData

	return distance, nil
}
