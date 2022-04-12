// Copyright (c) 2021 Michael Treanor
// https://
// MIT License

// Package errorlogger implements error logging to a variety of
// output formats. The goal of this package is to provide a simple
// and efficient mechanism for managing, testing, debugging, and
// changing options for error logging throughout a program.
//
// It is a drop-in replacement for the standard library log package
// and the popular Logrus package.
//
// Code like this works as expected without any changes:
//  log.Errorf("this is an error: %v", err)
//  log.Fatal("no input file provided.")
//
// Usage
//
// A global logger with a default logging function is supplied:
//  var log = errorlogger.Log
//  log.Error("sample log error message")
//
// using the variable 'log' matches most code using the standard
// library 'log' package or Logrus package.
//
// Logging
//
// The default global error logging function is supplied.
//  var Err = errorlogger.Err
//
// Err wraps errors, adds custom messages, formats errors, and outputs
// messages to the correct io.Writer as specified in the options.
//
// Calling this function will perform all package-level logging
// and error wrapping. It will then return the error otherwise
// unchanged and ready to propagate up.
//
// If you do not intend to use any options or disable the logger,
// it may be more convenient to use only the function alias to call the
// most common method, Err(), like this:
//  var Err = errorlogger.Err
//
// then, just call the function within error blocks:
//  err := someProcess(stuff)
//  if err != nil {
//   return Err(err)
//  }
// or
//  return Err(someProcess(stuff))
//
// or even this
//  _ = Err(someProcess(stuff)) // log errors only and continue
//
// if the error does not need to be propagated (bubbled) up. (This is
// not generally recommended.)
//
// Examples
//
// file open
//  f, err := os.Open("myfile")
//  if err != nil {
//  	return Err(err)
//  }
// get environment variable
//  env := os.Getenv("PATH")
//  if env == "" {
//  	return "", Err(os.ErrNotExist)
//  }
//  return env, nil
// check return value while returning an error
//  return Err(os.Chmod("myfile", 420))
//
// Defaults
//
// The global defaults may be aliased if there is a concern
// about name collisions:
//  var LogThatWontConflict = errorlogger.Log
//  var ErrThatWontConflict = errorlogger.Err
//
// By default, logging is enabled and ANSI colorized text is sent to
// stderr of the TTY. If it is changed and you wish to return to the
// default text formatter, use
//  log.SetText()
//
// Logging can also be redirected to a file or any io.Writer
//  log.SetLogOutput(w io.Writer)
//
// To create a new logger with default behaviors, use:
//  var log = errorlogger.New()
// and start logging!
//
// (The defaults are output to os.Stderr, ANSI color, include
// timestamps, logging enabled, default log level(INFO),
// no error wrapping, default log function, and use default
// Logrus logger as pass-through.)
//
// Customize
//
// If you want to customize the logger, use:
//  NewWithOptions(enabled bool, fn LoggerFunc, wrap error, logger interface{}) ErrorLogger
//
// Some additional features of this package include:
//
// - easy configuration of JSON logging:
//  log.EnableJSON(true) // true for pretty printing
//
// - return to the default text formatting
//  log.SetText() // change to default text formatter
//
// - easy configuration of custom output formatting:
//  log.SetFormatter(myJSONformatter) // set a custom formatter
//
// - easy configuration of numerous third party formatters.
//
//
// - Set log level - the verbosity of the logging may be adjusted.
// Allowed values are Panic, Fatal, Error, Warn, Info, Debug, Trace.
// The default is "Info"
//  log.SetLogLevel("INFO") // Set log level - uppercase string ...
//  log.SetLogLevel("error") // ... or lowercase string accepted
//
// Performance
//
// Error logging may be disabled during performance critical operations:
//  log.Disable() // temporarily disable logging
//  defer log.Enable()  // enable logging after critical code
//
// In this case, the error function is replaced with a noop function. This
// removed any enabled/disabled check and usually results in a performance
// gain when compared to checking a flag during every possible operation
// that may request logging.
//
// Logging is deferred or reenabled with
//  log.Enable() // after performance sensitive portion, enable logging
//
// This may be done at any time and as often as desired.
//
// - SetLoggerFunc allows setting of a custom logger function.
// The default is log.Error(), which is compatible with
// the standard library log package and logrus.
//  log.SetLoggerFunc(fn LoggerFunc)
//
// - SetErrorWrap allows ErrorLogger to wrap errors in a
// specified custom type for use with errors.Is():
//  log.SetErrorWrap(wrap error)
//
// For example, if you want all errors returned to be
// considered type *os.PathError, use:
//  log.SetErrorWrap(&os.PathError{})
//
// To wrap all errors in a custom type, use:
//  log.SetErrorWrap(myErrType{}) // wrap all errors in a custom type
package errorlogger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

const defaultEnabled bool = true

var (
	// Log is the default global ErrorLogger. It implements
	// the ErrorLogger interface as well as the basic
	// logrus.Logger interface, which is compatible with the
	// standard library "log" package.
	//
	// In the case of name collisions with 'Log', use an alias
	// instead of creating a new instance. For example:
	//  var mylogthatwontmessthingsup = errorlogger.Log
	Log = New()

	// Err is the logging function for the global ErrorLogger.
	Err = Log.Err

	// ErrInvalidWriter is returned when an output writer is
	// nil or does not implement io.Writer.
	ErrInvalidWriter = os.ErrInvalid

	// defaultLogFunc is Log.Error, which will log messages
	// of level ErrorLevel or higher.
	defaultLogFunc LoggerFunc = defaultlogger.Error

	// defaultErrWrap is the default error used to wrap
	// errors processed with Err. A <nil> value disables
	// error wrapping.
	defaultErrWrap error = nil

	// Discard is a Writer on which all Write calls succeed without doing anything.
	DiscardWriter io.Writer = io.Discard
)

type (

	// LoggerFunc defines the function signature for logging functions.
	LoggerFunc = func(args ...interface{})

	// ErrorFunc defines the function signature to choose the logging function.
	ErrorFunc = func(err error) error

	// ErrorLogger implements error logging to a logrus log
	// (or a standard library log) by providing convenience
	// methods, advanced formatting options, more automated
	// logging, a more efficient way to log errors within
	// code, and methods to temporarily disable/enable
	// logging, such as in the case of performance
	// optimization or during critical code blocks.
	ErrorLogger interface {

		// Disable disables logging and sets a no-op function for
		// Err() to prevent slowdowns while logging is disabled.
		Disable()

		// Enable enables logging and restores the Err() logging functionality.
		Enable()

		// EnableText enables text formatting of log errors (default)
		SetText()

		// EnableJSON enables JSON formatting of log errors
		SetJSON(pretty bool)

		// LogLevel sets the logging level from a string value.
		// Allowed values: Panic, Fatal, Error, Warn, Info, Debug, Trace
		SetLogLevel(lvl string) error

		// Err logs an error to the provided logger, if it is enabled,
		// and returns the error unchanged.
		Err(err error) error

		// SetLoggerFunc allows setting of the logger function.
		// The default is log.Error(), which is compatible with
		// the standard library log package and logrus.
		SetLoggerFunc(fn LoggerFunc)

		// SetErrorWrap allows ErrorLogger to wrap errors in a
		// specified custom type. For example, if you want all errors
		// returned to be of type *os.PathError
		SetErrorWrap(wrap error)

		// SetCustomMessage allows automated addition of a custom
		// message to all log messages generated by this
		// logger.
		SetCustomMessage(msg string)

		logrusLogger
	}

	// errorLogger implements ErrorLogger with logrus or the
	// standard library log package.
	errorLogger struct {
		wrap    error      // `default:"nil"` // nil = disabled
		msg     string     // `default:""` // the empty string = disabled
		errFunc ErrorFunc  // `default:"()yesErr"`
		logFunc LoggerFunc // `default:"defaultLogFunc"`
		*Logger            // `default:"defaultlogger"`
	}
)

// New returns a new ErrorLogger with default options and
// logging enabled.
// Most users will not need to call this, since the default
// global ErrorLogger 'Log' is provided.
//
// In the case of name collisions with 'Log', use an alias
// instead of creating a new instance. For example:
//  var mylogthatwontmessthingsup = errorlogger.Log
func New() ErrorLogger {
	return NewWithOptions(defaultEnabled, "", defaultLogFunc, defaultErrWrap, defaultlogger)
}

// NewWithOptions returns a new ErrorLogger with options
// determined by parameters. To use defaults, use nil for
// any option except 'enabled'.
//
// - enabled: defines the initial logging state.
//
// - fn: defines a custom logging function used to log information.
//
// - wrap: defines a custom error type to wrap all errors in.
//
// - logger: defines a custom logger to use.
func NewWithOptions(enabled bool, msg string, fn LoggerFunc, wrap error, logger *Logger) ErrorLogger {
	return newTestStruct(enabled, msg, wrap, fn, logger)
}

func (e *errorLogger) Writer() io.Writer {
	return e.Out
}

// SetErrorWrap allows ErrorLogger to wrap all errors in a
// specified custom error type.
// Example:
//  log.SetErrorWrap(&os.PathError{})
// Setting wrap == nil will disable wrapping of errors:
//  log.SetErrorWrap(nil)
func (e *errorLogger) SetErrorWrap(wrap error) { e.wrap = wrap }

// SetCustomMessage allows ErrorLogger to add a specified
// custom string to all errors.
// Example:
//  log.SetCustomMessage("MyApp error occurred!")
// Setting msg == "" will disable this feature:
//  log.SetCustomMessage("")
//
// TODO not implemented in logging code yet ...
func (e *errorLogger) SetCustomMessage(msg string) { e.msg = msg }

// SetJSON sets the log format to JSON. The JSON output conforms
// to RFC 7159 (https://www.rfc-editor.org/rfc/rfc7159.html) from
// March 2014.
//
// It should be noted that this format has been obsoleted the
// latest version of the JSON standard from December 2017,
// RFC 8259 (https://datatracker.ietf.org/doc/html/rfc8259)
//
// The default is compact "ugly" json. A "pretty" format can be
// selected with
//  Log.SetOptions()
//
// Use
//  Log.SetText()
// to return to the default Text formatter.
//
// In general,
//  Log.Setformatter(myformatter)
// can be used to set any custom formatter.
//
// Many other third party logging formatters are available.
//
// - FluentdFormatter. Formats entries that can be parsed by Kubernetes and Google Container Engine.
//
// - GELF. Formats entries so they comply to Graylog's GELF 1.1 specification.
//
// - logstash. Logs fields as Logstash Events.
//
// - prefixed. Displays log entry source along with alternative layout.
//
// - zalgo. Invoking the Power of Zalgo.
//
// - nested-logrus-formatter. Converts logrus fields to a nested structure.
//
// - powerful-logrus-formatter. get fileName, log's line number and the latest function's name when print log; Sava log to files.
//
// - caption-json-formatter. logrus's message json formatter with human-readable caption added.
//
// Reference: https://pkg.go.dev/github.com/sirupsen/logrus#JSONFormatter
func (e *errorLogger) SetJSON(pretty bool) {
	// e.SetErrorWrap(&os.PathError{})
	f := NewJSONFormatter(pretty)
	e.SetFormatter(f)
}

// SetText sets the log format to Text. This is the default
// formatter.
//
// It provides ANSI colorized (if available) TTY output to os.Stderr.
//
// Use
//  Log.SetJSON()
// to switch to the JSON formatter.
//
// In general,
//  Log.Setformatter(myformatter)
// can be used to set any custom formatter.
//
// Many other third party logging formatters are available.
//
// - FluentdFormatter. Formats entries that can be parsed by Kubernetes and Google Container Engine.
//
// - GELF. Formats entries so they comply to Graylog's GELF 1.1 specification.
//
// - logstash. Logs fields as Logstash Events.
//
// - prefixed. Displays log entry source along with alternative layout.
//
// - zalgo. Invoking the Power of Zalgo.
//
// - nested-logrus-formatter. Converts logrus fields to a nested structure.
//
// - powerful-logrus-formatter. get fileName, log's line number and the latest function's name when print log; Sava log to files.
//
// - caption-json-formatter. logrus's message json formatter with human-readable caption added.
// Reference: https://pkg.go.dev/github.com/sirupsen/logrus#TextFormatter
func (e *errorLogger) SetText() { e.SetFormatter(DefaultTextFormatter) }

// SetLoggerFunc sets the logger function that is used to
// write log messages. This allows rapid switching between loggers
// as well as turning the logging off and on regularly.
//
// The default is Log.Error(err), which is compatible with
// the standard library log package and logrus. Setting this to
// a no-op function allows fast pass through of logging
// information without the need for checks.
//
// The function signature must be of type LoggerFunc:
//  func(args ...interface{}).
func (e *errorLogger) SetLoggerFunc(fn LoggerFunc) {
	if fn == nil {
		e.logFunc = e.Error
	} else {
		e.logFunc = fn
	}
}

// SetLogLevel converts lvl to a compatible log level and sets the log level.
//
// Allowed values: Panic, Fatal, Error, Warn, Info, Debug, Trace
func (e *errorLogger) SetLogLevel(lvl string) error {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		return Err(err)
	}
	e.Logger.SetLevel(level)
	return nil
}

// SetLogOutput sets the output writer for logging.
// The default is os.Stderr. Any io.Writer can be setup
// to receive messages.
func (e *errorLogger) SetLogOutput(w io.Writer) error {
	if w == nil {
		return Err(ErrInvalidWriter)
	}

	// this is a redundant check, but I was considering
	// allowing other values to be passed and type
	// checked internally
	switch v := w.(type) {
	case io.Writer:
		e.SetOutput(v)
		return nil
	default:
		return Err(ErrInvalidWriter)
	}
}
