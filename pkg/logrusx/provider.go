package logrusx

import (
	"context"
	"fmt"
	"git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/pkg/logrusx/dailylogger"
	"io"
	"path"

	"github.com/sirupsen/logrus"
)

type Provider struct {
	ctx    *context.Context
	logger *logrus.Logger
}

func NewProvider(ctx *context.Context, cfg Config) *Provider {
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	logger.SetReportCaller(true)
	logger.SetFormatter(&LoggerFormatter{})

	// Send logs with level higher than warning to stderr
	logger.AddHook(&WriterHook{
		Writer: dailylogger.NewDailyRotateLogger(path.Join(cfg.Dir, fmt.Sprintf("%s.error.log", cfg.FileName)), cfg.MaxSize, cfg.LocalTime, cfg.Compress),
		LogLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
		},
	})

	// Send info and debug logs to stdout
	logger.AddHook(&WriterHook{
		Writer: dailylogger.NewDailyRotateLogger(path.Join(cfg.Dir, fmt.Sprintf("%s.info.log", cfg.FileName)), cfg.MaxSize, cfg.LocalTime, cfg.Compress),
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.DebugLevel,
		},
	})
	return &Provider{ctx: ctx, logger: logger}
}

func (p Provider) GetLogger(name string) *LoggerEntry {
	return NewLoggerEntry(p.logger.WithField("who", name))
}

// WriterHook is a hook that writes logs of specified LogLevels to specified Writer
type WriterHook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
}

// Fire will be called when some logging function is called with current hook
// It will format log entry to string and write it to appropriate writer
func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}

// Levels define on which log levels this hook would trigger
func (hook *WriterHook) Levels() []logrus.Level {
	return hook.LogLevels
}
