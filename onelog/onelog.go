package onelog

import (
	"fmt"
	"os"

	one "github.com/francoispqt/onelog"
	"github.com/millken/logger"
)

type oneLogger struct {
	logger *one.Logger
}

func New(config logger.Configuration) (logger.Logger, error) {
	logger := one.New(
		os.Stdout,
		one.ALL, // shortcut for onelog.DEBUG|onelog.INFO|onelog.WARN|onelog.ERROR|onelog.FATAL,
	)

	return &oneLogger{
		logger: logger,
	}, nil
}

func (l *oneLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debug(fmt.Sprintf(format, args...))
}

func (l *oneLogger) Infof(format string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, args...))
}

func (l *oneLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, args...))
}

func (l *oneLogger) Errorf(format string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, args...))
}

func (l *oneLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal(fmt.Sprintf(format, args...))
}

func (l *oneLogger) Panicf(format string, args ...interface{}) {
	l.logger.Fatal(fmt.Sprintf(format, args...))
}

func (l *oneLogger) WithFields(fields logger.Fields) logger.Logger {
	return l
}
