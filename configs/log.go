package configs

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLog(fn string) *logrus.Entry {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	l.Out = file
	return logrus.NewEntry(l)

}
