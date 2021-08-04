package logging

import (
	"github.com/sirupsen/logrus"
)

type Fields logrus.Fields

type Logger struct {
	log *logrus.Entry
}

func (l *Logger) Info(fields map[string]interface{}, method, message string) {
	if method != "" {
		l.log.WithFields(logrus.Fields{
			"2-method": method,
		}).WithFields(fields).Info(message)
	} else {
		l.log.WithFields(fields).Info(message)
	}
}
func (l *Logger) Warn(fields map[string]interface{}, method, message string) {
	if method != "" {
		l.log.WithFields(logrus.Fields{
			"2-method": method,
		}).WithFields(fields).Warn(message)
	} else {
		l.log.WithFields(fields).Warn(message)
	}
}
func (l *Logger) Error(method string, err error) {
	if method != "" {
		l.log.WithFields(logrus.Fields{
			"2-method": method,
		}).Error(err)
	} else {
		l.log.Error(err)
	}
}
func (l *Logger) Fatal(method string, err error) {
	if method != "" {
		l.log.WithFields(logrus.Fields{
			"2-method": method,
		}).Fatal(err)
	} else {
		l.log.Fatal(err)
	}
}
func (l *Logger) Panic(method string, err error) {
	if method != "" {
		l.log.WithFields(logrus.Fields{
			"2-method": method,
		}).Panic(err)
	} else {
		l.log.Panic(err)
	}
}

func NewLogger(name string) (log *Logger) {
	l := logrus.New()
	log = &Logger{
		log: l.WithField("1-service", name),
	}
	return
}
