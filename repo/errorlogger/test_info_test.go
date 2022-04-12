package errorlogger

import (
	"testing"

	"github.com/sirupsen/logrus"
)

type errorloggerTestArgs struct {
	enabled bool
	msg     string
	wrap    error
	fn      func(args ...interface{})
	logger  *Logger
}

// internal tests directly on private structs
var (
	errorLoggerTestStruct = newTestStruct(true, "", nil, nil, nil)
	wrapTestStruct        = newTestStruct(true, "", fakeSysCallError, nil, nil)
	messageTestStruct     = newTestStruct(true, "fake test message", nil, nil, nil)

	privateStructTests = []struct {
		name string
		e    *errorLogger
	}{
		{"errorLoggerTestStruct", errorLoggerTestStruct},
		{"wrapTestStruct", wrapTestStruct},
		{"messageTestStruct", messageTestStruct},
	}

	// errorloggerTests provide a set of instantiated errorloggers
	// used for tests.
	// input uses type interface{} in order to allow testing with
	// a variety of types that may or may not implement ErrorLogger.
	//
	// If ErrorLogger is not implemented, wantErr bool should be
	// set to true.
	errorloggerTests = []struct {
		name    string
		args    errorloggerTestArgs
		want    ErrorLogger
		wantErr bool
	}{
		// control
		{"global ErrorLogger", errorloggerTestArgs{}, testDefaultLogger, false},

		// Check for false positive and false negative errors
		// Test New() should pass and nil should fail
		{"New()", errorloggerTestArgs{}, New(), false},
		{"nil", errorloggerTestArgs{}, nil, true},

		// NewWithOptions() is also tested here
		{"NewWithOptions(false, nil, nil, nil) (should pass)", errorloggerTestArgs{}, NewWithOptions(false, "", nil, nil, nil), false},

		{"NewWithOptions(true, nil, nil, nil)", errorloggerTestArgs{}, NewWithOptions(true, "", nil, nil, nil), false},
		{"NewWithOptions(false, nil, nil, nil)", errorloggerTestArgs{}, NewWithOptions(false, "", nil, nil, nil), false},
		{"NewWithOptions(true, nil, nil, string)", errorloggerTestArgs{}, NewWithOptions(true, "", nil, nil, nil), false},
		{"NewWithOptions(true, nil, nil, integer)", errorloggerTestArgs{}, NewWithOptions(true, "", nil, nil, nil), false},
		{"NewWithOptions(all defaults ...)", errorloggerTestArgs{}, NewWithOptions(true, "", defaultLogFunc, defaultErrWrap, defaultlogger), false},
		{"NewWithOptions(false, DefaultLogFunc, nil)", errorloggerTestArgs{}, NewWithOptions(true, "", defaultLogFunc, nil, nil), false},

		// Various tests using private struct
		{"logrus logger in errorLogger (not public)", errorloggerTestArgs{}, &errorLogger{Logger: &logrus.Logger{}}, false},
		{"default ErrorLogger with nil wrapper (not public)", errorloggerTestArgs{}, &errorLogger{wrap: nil}, false},
		// Do not need a check for this in the constructor since errorLogger is not exported
		// But something to be aware of ...
		// {"ErrorLogger with nil logger (should fail)", &errorLogger{Logger: nil}, true},
	}
)

func Test_newTestStruct(t *testing.T) {
	for _, tt := range errorloggerTests {
		t.Run(tt.name, func(t *testing.T) {
			got := newTestStruct(tt.args.enabled, tt.args.msg, tt.args.wrap, tt.args.fn, tt.args.logger)
			var want ErrorLogger
			if tt.want == nil {
				if !tt.wantErr {
					t.Logf("newTestStruct() logger cannot be nil: %v(%T)", tt.want, tt.want)
					t.FailNow()
					// } else {
					// 	want = New()
				}
			}

			// TODO check test logic here ...
			want = New()

			g := got.errFunc(errFake)
			w := want.Err(errFake)

			if w.Error() != g.Error() {
				t.Errorf("newTestStruct() error string not equal = %v, want %v", g, w)
			}
		})
	}
}
