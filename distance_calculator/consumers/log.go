package consumers

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fabrizioperria/toll/shared/types"
	"github.com/sirupsen/logrus"
)

type LogConsumerMiddleware struct {
	next   DataConsumer
	logger *logrus.Logger
}

func NewLogConsumerMiddleware(next DataConsumer) *LogConsumerMiddleware {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory: ", err)
		return nil
	}
	os.MkdirAll(homeDir+"/log/toll", 0o755)
	f, err := os.OpenFile(homeDir+"/log/toll/distance-calc.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o755)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil
	}
	l.SetOutput(io.MultiWriter(os.Stdout, f))

	return &LogConsumerMiddleware{
		next:   next,
		logger: l,
	}
}

func (lm *LogConsumerMiddleware) Consume() (data types.OBUData, err error) {
	defer func(t time.Time) {
		lm.logger.WithFields(logrus.Fields{
			"obuId":     data.ObuId,
			"latitude":  data.Latitude,
			"longitude": data.Longitude,
			"timestamp": data.Timestamp,
		}).Info("Consumed")
	}(time.Now())
	data, err = lm.next.Consume()
	if err != nil {
		lm.logger.Error("Error consuming message")
	}
	return
}
