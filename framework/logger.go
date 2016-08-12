package framework

import (
  log "github.com/Sirupsen/logrus"
)

type Logger interface {
  Info(args ...interface{})
  Infof(format string, args ...interface{})
  Warn(args ...interface{})
  Warnf(format string, args ...interface{})
}

type LogrusLogger struct {
  *log.Logger
}

func (inner *LogrusLogger) Info(args ...interface{}) {
  inner.Logger.Info(args...)
}
func (inner *LogrusLogger) Infof(format string, args ...interface{}) {
  inner.Logger.Infof(format, args...)
}
func (inner *LogrusLogger) Warn(args ...interface{}) {
  inner.Logger.Warn(args...)
}
func (inner *LogrusLogger) Warnf(format string, args ...interface{}) {
  inner.Logger.Warnf(format, args...)
}

func DefaultLogger() Logger {
  logger := log.New()

  return &LogrusLogger {
    Logger: logger,
  }
}
