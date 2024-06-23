package producers

import (
	"os"
	"time"

	"github.com/fabrizioperria/toll/shared/logger"
	"github.com/fabrizioperria/toll/shared/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next   DataProducer
	logger *logrus.Logger
}

func NewLogMiddleware(next DataProducer) DataProducer {
	return &LogMiddleware{
		next:   next,
		logger: logger.LoggerFactory(os.Getenv("LOG_PATH")),
	}
}

func (lm *LogMiddleware) Produce(obuData types.OBUData) error {
	defer func(t time.Time) {
		lm.logger.WithFields(logrus.Fields{
			"obuId":     obuData.ObuId,
			"longitude": obuData.Longitude,
			"latitude":  obuData.Latitude,
			"duration":  float64(time.Since(t).Microseconds()) / 1000,
		}).Info("Produced")
	}(time.Now())
	err := lm.next.Produce(obuData)
	if err != nil {
		lm.logger.WithFields(logrus.Fields{
			"obuId":     obuData.ObuId,
			"longitude": obuData.Longitude,
			"latitude":  obuData.Latitude,
		}).Error("Error while producing")
	}
	return err
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
