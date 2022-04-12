package errorlogger

import (
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// test_info provides samples and test cases for the tests
// and benchmarks in this package.

var (
	testLogrusLogger *Logger = &logrus.Logger{
		Out:       io.Discard,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
	testDefaultLogger ErrorLogger = NewWithOptions(true, "", nil, nil, testLogrusLogger)
	errFake           error       = errors.New("fake")
	fakeSysCallError  error       = os.NewSyscallError("fake syscall error", fmt.Errorf("fake syscall error"))
	blankSysCallError error       = new(os.SyscallError)
	blankPathError    error       = new(fs.PathError)
)

func newTestStruct(enabled bool, msg string, wrap error, _ func(args ...interface{}), logger *Logger) *errorLogger {
	if logger == nil {
		logger = defaultlogger
	}

	e := errorLogger{
		msg:    msg,
		Logger: logger,
	}

	if enabled {
		e.Enable()
	} else {
		e.Disable()
	}

	// e.Logger = logger
	e.logFunc = e.Error

	if wrap == nil {
		// the defaultErrWrap is actually nil ... so this is not needed.
		// However, if the default is later changed to a package-wide
		// wrapper, this will be a valid check
		wrap = defaultErrWrap
	}
	e.wrap = wrap
	return &e
}
