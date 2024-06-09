package producers

import (
	"log"
	"time"

	"github.com/fabrizioperria/toll/shared/types"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) DataProducer {
	return &LogMiddleware{next: next}
}

func (lm *LogMiddleware) Produce(obuData types.OBUData) error {
	defer func(t time.Time) {
		log.Printf("PRODUCED: %+v in %v\n", obuData, time.Since(t))
	}(time.Now())
	return lm.next.Produce(obuData)
}

func (lm *LogMiddleware) Flush() int {
	if lm.next == nil {
		return 0
	}
	return lm.next.Flush()
}

func (lm *LogMiddleware) Close() {
	if lm.next == nil {
		return
	}
	lm.next.Close()
}
