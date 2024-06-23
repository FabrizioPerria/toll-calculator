package consumers

import (
	"os"
	"time"

	"github.com/fabrizioperria/toll/shared/logger"
	"github.com/fabrizioperria/toll/shared/types"
	"github.com/sirupsen/logrus"
)

type LogConsumerMiddleware struct {
	next   DataConsumer
	logger *logrus.Logger
}

func NewLogConsumerMiddleware(next DataConsumer) *LogConsumerMiddleware {
	return &LogConsumerMiddleware{
		next:   next,
		logger: logger.LoggerFactory(os.Getenv("LOG_PATH")),
	}
}

func (lm *LogConsumerMiddleware) Consume() (data types.OBUData, err error) {
	defer func(t time.Time) {
		lm.logger.WithFields(logrus.Fields{
			"obuId":     data.ObuId,
			"latitude":  data.Latitude,
			"longitude": data.Longitude,
			"duration":  float64(time.Since(t).Microseconds()) / 1000,
		}).Info("Consumed")
	}(time.Now())
	data, err = lm.next.Consume()
	if err != nil {
		lm.logger.Error("Error consuming message")
	}
	return
}
