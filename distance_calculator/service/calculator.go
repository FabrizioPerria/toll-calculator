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
	return &CalculatorService{}
}

var lastObuData map[int]types.OBUData = make(map[int]types.OBUData, 50)

func (c *CalculatorService) Distance(obuData types.OBUData) (float64, error) {
	lastObuDataForID := lastObuData[obuData.OBUID]
	distance := math.Sqrt(math.Pow(obuData.Latitude-lastObuDataForID.Latitude, 2) + math.Pow(obuData.Longitude-lastObuDataForID.Longitude, 2))
	lastObuData[obuData.OBUID] = obuData

	return distance, nil
}
