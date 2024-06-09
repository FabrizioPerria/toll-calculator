package consumers

import (
	"log"
	"time"

	"github.com/fabrizioperria/toll/shared/types"
)

type LogConsumerMiddleware struct {
	next DataConsumer
}

func NewLogConsumerMiddleware(next DataConsumer) *LogConsumerMiddleware {
	return &LogConsumerMiddleware{next: next}
}

func (lm *LogConsumerMiddleware) Consume() (data types.OBUData, err error) {
	defer func(t time.Time) {
		log.Printf("CONSUMED: %+v in %v\n", data, time.Since(t))
	}(time.Now())
	data, err = lm.next.Consume()
	return
}
