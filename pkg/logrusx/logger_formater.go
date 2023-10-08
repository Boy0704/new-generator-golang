package logrusx

import (
	"bytes"
	"fmt"

	"github.com/sirupsen/logrus"
)

type LoggerFormatter struct{}

var levelList = []string{
	"PANIC",
	"FATAL",
	"ERROR",
	"WARN",
	"INFO",
	"DEBUG",
	"TRACE",
}

func (mf *LoggerFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	level := levelList[int(entry.Level)]
	data := entry.Data
	b.WriteString(fmt.Sprintf("%s - %s - %s - %s - %s\n",
		entry.Time.Format("2006-01-02 15:04:05,678"), level, data["who"], data["guid"], entry.Message))
	return b.Bytes(), nil
}
