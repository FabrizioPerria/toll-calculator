package service

import (
	"github.com/fabrizioperria/toll/shared/types"
)

type LogServiceMiddleware struct {
	next Calculator
}

func NewLogServiceMiddleware(next Calculator) *LogServiceMiddleware {
	return &LogServiceMiddleware{next: next}
}

func (lm *LogServiceMiddleware) Distance(obuData types.OBUData) (distance float64, err error) {
	// defer func() {
	// log.Println("CALCULATED: ", distance)
	// }()
	distance, err = lm.next.Distance(obuData)
	return
}
