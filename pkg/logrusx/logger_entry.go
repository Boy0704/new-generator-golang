package logrusx

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type LoggerEntry struct {
	entry *logrus.Entry
}

func NewLoggerEntry(entry *logrus.Entry) *LoggerEntry {
	return &LoggerEntry{entry: entry}
}

func (le LoggerEntry) WithField(key string, value interface{}) *LoggerEntry {
	le.entry = le.entry.WithField(key, value)
	return &le
}

func (le LoggerEntry) Debug(guid string, args ...interface{}) {
	le.entry.WithField("guid", guid).Debug(args...)
}

func (le LoggerEntry) Info(guid string, args ...interface{}) {
	le.entry.WithField("guid", guid).Info(args...)
}

func (le LoggerEntry) Warn(guid string, message interface{}, err error) {
	le.entry.WithField("guid", guid).Warn(fmt.Errorf("%s %+v", message, err))
}

func (le LoggerEntry) Error(guid string, message interface{}, err error) {
	le.entry.WithField("guid", guid).Error(fmt.Errorf("%s %+v", message, err))
}

func (le LoggerEntry) Fatal(guid string, message interface{}, err error) {
	le.entry.WithField("guid", guid).Fatal(fmt.Errorf("%s %+v", message, err))
}

func (le LoggerEntry) Panic(guid string, message interface{}, err error) {
	le.entry.WithField("guid", guid).Panic(fmt.Errorf("%s %+v", message, err))
}
