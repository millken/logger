package logger

import "strings"

var logx Logger

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

type Level string

const (
	//Debug has verbose message
	Debug Level = "debug"
	//Info is default log level
	Info Level = "info"
	//Warn is for logging messages about possible issues
	Warn Level = "warn"
	//Error is for logging errors
	Error Level = "error"
	//Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	Fatal Level = "fatal"
)

//Logger is our contract for the logger
type Logger interface {
	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})

	Panicf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger
}

// Configuration stores the config for the Logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      Level
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         Level
	FileLocation      string
}

func GetLevel(level string) Level {
	level = strings.ToLower(level)
	switch level {
	case "info":
		return Info
	case "warn":
		return Warn
	case "debug":
		return Debug
	case "error":
		return Error
	case "fatal":
		return Fatal
	default:
		return Info
	}
}

//New returns an instance of Logger
func New(logger Logger) {
	logx = logger
}

func Debugf(format string, args ...interface{}) {
	logx.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logx.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logx.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logx.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logx.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logx.Panicf(format, args...)
}

func WithFields(keyValues Fields) Logger {
	return logx.WithFields(keyValues)
}
