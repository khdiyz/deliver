package logger

import (
	"fmt"
	"path"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	instance *Logger
	once     sync.Once
)

// Logger wraps logrus.Entry for structured logging
type Logger struct {
	*logrus.Entry
}

// GetLogger returns the singleton logger instance
func GetLogger() *Logger {
	once.Do(func() {
		baseLogger := logrus.New()
		baseLogger.SetReportCaller(true)
		baseLogger.Formatter = &logrus.TextFormatter{
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				filename := path.Base(frame.File)
				return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
			},
			EnvironmentOverrideColors: true,
			DisableColors:             false,
			FullTimestamp:             true,
		}
		baseLogger.SetLevel(logrus.TraceLevel)
		instance = &Logger{logrus.NewEntry(baseLogger)}
	})
	return instance
}

// WithField creates a logger with an additional field
func (log *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{log.Entry.WithField(key, value)}
}

// WithFields creates a logger with multiple additional fields
func (log *Logger) WithFields(fields logrus.Fields) *Logger {
	return &Logger{log.Entry.WithFields(fields)}
}
