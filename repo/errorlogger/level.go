package errorlogger

import "github.com/sirupsen/logrus"

// A constant exposing all logging levels
//
// Reference: https://github.com/sirupsen/logrus
var AllLevels []Level = logrus.AllLevels

// Level type
//
// Reference: https://github.com/sirupsen/logrus
type Level = logrus.Level

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
//
// Reference: github.com/sirupsen/logrus
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota

	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel

	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel

	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel

	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel

	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel

	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)
