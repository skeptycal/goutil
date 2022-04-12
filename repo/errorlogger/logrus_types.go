package errorlogger

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const (

	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	DefaultLogLevel Level = InfoLevel

	// DefaultTimestampFormat is time.RFC3339FA
	//
	// Note that this is not the most current standard but it is the
	// most stable and recommended with the Go standard library.
	//
	// Additional notes
	//
	// The RFC822, RFC850, and RFC1123 formats should be applied only to
	// local times. Applying them to UTC times will use "UTC" as the time
	// zone abbreviation, while strictly speaking those RFCs require the
	// use of "GMT" in that case.
	//
	// In general RFC1123Z should be used instead of RFC1123 for servers
	// that insist on that format, and RFC3339 should be preferred for
	// new protocols.
	//
	// While RFC3339, RFC822, RFC822Z, RFC1123, and RFC1123Z are useful
	// for formatting, when used with time.Parse they do not accept all
	// the time formats permitted by the RFCs and they do accept time
	// formats not formally defined.
	//
	// The RFC3339Nano format removes trailing zeros from the seconds
	// field and thus may not sort correctly once formatted.
	DefaultTimestampFormat string = time.RFC3339
)

var (

	// defaultlogger initializes a default logrus logger.
	// Reference: https://github.com/sirupsen/logrus/
	defaultlogger = &Logger{
		Out:       os.Stderr,
		Formatter: DefaultTextFormatter,
		Hooks:     make(logrus.LevelHooks),
		Level:     DefaultLogLevel,
	}
)

type (

	// Fields type, used to pass to `WithFields`.
	Fields = logrus.Fields

	// An entry is the final or intermediate Logrus logging entry. It contains all the fields passed with WithField{,s}. It's finally logged when Trace, Debug, Info, Warn, Error, Fatal or Panic is called on it. These objects can be reused and passed around as much as you wish to avoid field duplication.
	Entry = logrus.Entry

	// Logger is the main structure used by errorlogger. It is a thinly veiled
	// wrapper around logrus.Logger with some additional functionality.
	// 	type Logger struct {
	//     // The logs are `io.Copy`'d to this in a mutex. It's common to set this to a
	//     // file, or leave it default which is `os.Stderr`. You can also set this to
	//     // something more adventurous, such as logging to Kafka.
	//     Out io.Writer
	//     // Hooks for the logger instance. These allow firing events based on logging
	//     // levels and log entries. For example, to send errors to an error tracking
	//     // service, log to StatsD or dump the core on fatal errors.
	//     Hooks LevelHooks
	//     // All log entries pass through the formatter before logged to Out. The
	//     // included formatters are `TextFormatter` and `JSONFormatter` for which
	//     // TextFormatter is the default. In development (when a TTY is attached) it
	//     // logs with colors, but to a file it wouldn't. You can easily implement your
	//     // own that implements the `Formatter` interface, see the `README` or included
	//     // formatters for examples.
	//     Formatter Formatter
	//
	//     // Flag for whether to log caller info (off by default)
	//     ReportCaller bool
	//
	//     // The logging level the logger should log at. This is typically (and defaults
	//     // to) `logrus.Info`, which allows Info(), Warn(), Error() and Fatal() to be
	//     // logged.
	//     Level Level
	//     // Used to sync writing to the log. Locking is enabled by Default
	//     mu MutexWrap
	//     // Reusable empty entry
	//     entryPool sync.Pool
	//     // Function to exit the application, defaults to `os.Exit()`
	//     ExitFunc exitFunc
	// 	}
	Logger = logrus.Logger

	// The Formatter interface is used to implement a custom Formatter.
	// It takes an `Entry`. It exposes all the fields, including the
	// default ones:
	//
	// * `entry.Data["msg"]`. The message passed from Info, Warn, Error ..
	// * `entry.Data["time"]`. The timestamp.
	// * `entry.Data["level"]. The level the entry was logged at.
	//
	// Any additional fields added with `WithField` or `WithFields` are
	// also in `entry.Data`. Format is expected to return an array of
	// bytes which are then logged to `logger.Out`.
	//
	// Reference: logrus@v1.8.1 formatter.go
	// 	type Formatter interface {
	// 		Format(*Entry) ([]byte, error)
	// 	}
	Formatter interface{ logrus.Formatter }

	// logrusLogger implements the most common functionality
	// of the logging interface of the Logrus package.
	//
	// This is the minimum interface required for incorporation
	// in an ErrorLogger.
	//
	// It is a compatible superset of the standard library
	// log package and a compatible subset of the ErrorLogger
	// package.
	logrusLogger interface {
		basicErrorLogger
		logrusCommonOptions
	}

	basicErrorLogger interface {
		WithField(key string, value interface{}) *logrus.Entry
		WithFields(fields logrus.Fields) *logrus.Entry
		WithError(err error) *logrus.Entry

		Debugf(format string, args ...interface{})
		Infof(format string, args ...interface{})
		Printf(format string, args ...interface{})
		Warnf(format string, args ...interface{})
		Warningf(format string, args ...interface{})
		Errorf(format string, args ...interface{})
		Fatalf(format string, args ...interface{})
		Panicf(format string, args ...interface{})

		Debug(args ...interface{})
		Info(args ...interface{})
		Print(args ...interface{})
		Warn(args ...interface{})
		Warning(args ...interface{})
		Error(args ...interface{})
		Fatal(args ...interface{})
		Panic(args ...interface{})

		Debugln(args ...interface{})
		Infoln(args ...interface{})
		Println(args ...interface{})
		Warnln(args ...interface{})
		Warningln(args ...interface{})
		Errorln(args ...interface{})
		Fatalln(args ...interface{})
		Panicln(args ...interface{})
	}

	// logrusCommonOptions implements several common options
	// that should be in the basic LogrusLogger interface.
	logrusCommonOptions interface {
		SetLevel(level Level)
		GetLevel() Level
		SetFormatter(formatter logrus.Formatter)
		SetOutput(output io.Writer)
	}

	// logrusLoggerComplete implements the complete exported
	// interface implemented by the logrus.Logger struct.
	//
	// Most users will not need to use this. ErrorLogger
	// contains the most common functionality, including the
	// basic LogrusLogger interface.
	logrusLoggerComplete interface {
		logrusLogger
		logrusOptions
		logrusLogFunctions
	}

	// logrusOptions implements rarely used logging options.
	// Instead of using this directly, create your own custom
	// interface that uses the options required.
	logrusOptions interface {
		WithContext(ctx context.Context) *logrus.Entry
		WithTime(t time.Time) *logrus.Entry
		Exit(code int)
		SetNoLock()
		AddHook(hook logrus.Hook)
		IsLevelEnabled(level Level) bool
		SetReportCaller(reportCaller bool)
		ReplaceHooks(hooks logrus.LevelHooks) logrus.LevelHooks
	}

	// logrusLogFunctions implements logrus Logrus
	// LogFunctions.
	// Instead of using this directly, create your own custom
	// interface that uses the options required.
	logrusLogFunctions interface {
		DebugFn(fn logrus.LogFunction)
		InfoFn(fn logrus.LogFunction)
		PrintFn(fn logrus.LogFunction)
		WarnFn(fn logrus.LogFunction)
		WarningFn(fn logrus.LogFunction)
		ErrorFn(fn logrus.LogFunction)
		FatalFn(fn logrus.LogFunction)
		PanicFn(fn logrus.LogFunction)
	}
)

var _ logrusLoggerComplete
