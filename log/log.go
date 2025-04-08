package log

import (
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	log  *logrus.Logger
	once sync.Once
)

func Log() *logrus.Logger {
	once.Do(func() {
		log = logrus.New()
		log.SetOutput(&lumberjack.Logger{
			Filename: "./gonet.log",
		})
		log.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	})

	return log
}
