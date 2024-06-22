package producers

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fabrizioperria/toll/shared/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next   DataProducer
	logger *logrus.Logger
}

func NewLogMiddleware(next DataProducer) DataProducer {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory: ", err)
		return nil
	}
	os.MkdirAll(homeDir+"/log/toll", 0o755)
	f, err := os.OpenFile(homeDir+"/log/toll/obuReceiver.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o755)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil
	}
	l.SetOutput(io.MultiWriter(os.Stdout, f))

	return &LogMiddleware{
		next:   next,
		logger: l,
	}
}

func (lm *LogMiddleware) Produce(obuData types.OBUData) error {
	defer func(t time.Time) {
		lm.logger.WithFields(logrus.Fields{
			"obuId":     obuData.ObuId,
			"longitude": obuData.Longitude,
			"latitude":  obuData.Latitude,
			"timestamp": obuData.Timestamp,
			"duration":  time.Since(t),
		}).Info("Produced")
	}(time.Now())
	err := lm.next.Produce(obuData)
	if err != nil {
		lm.logger.WithFields(logrus.Fields{
			"obuId":     obuData.ObuId,
			"longitude": obuData.Longitude,
			"latitude":  obuData.Latitude,
			"timestamp": obuData.Timestamp,
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
