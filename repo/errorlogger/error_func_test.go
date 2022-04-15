// Copyright (c) 2021 Michael Treanor
// https://
// MIT License

package errorlogger

import (
	"fmt"
	"io"
	"testing"
)

var (
	nopWriter io.Writer = io.Discard // mock to avoid an extra dependency
	lenWriter io.Writer = io.Discard // mock to avoid an extra dependency

	yesnologger     = New()
	nopWriterlogger = NewWithOptions(true, "", nil, nil, nil)
	lenWriterlogger = NewWithOptions(true, "", nil, nil, nil)
	logrusonly      = New()
	nillogger       = &errorLogger{nil, "", nil, nil, nil}

	fakeOuter error
)

func loggerFromWriter(w io.Writer) ErrorLogger {
	return NewWithOptions(true, "", nil, nil, nil)
}

func init() {
	nillogger.Disable()
	// nillogger.SetOutput(nil)

	yesnologger.SetOutput(nopWriter)
	yesnologger.Enable()

	logrusonly.SetOutput(nopWriter)
	logrusonly.Enable()

	nopWriterlogger.SetOutput(nopWriter)
	lenWriterlogger.SetOutput(lenWriter)
}

// benchmark results
/*

Is it worth it? Yes, having separate functions and storing the current one in the struct
is 2.5 times faster than simply checking for a nil error in cases where there is no error

In cases where there is an error, processing the error takes 100 times as long as not
processing the error, so it is *critical* to have a way to efficiently enable and disable
logging for portions of the code where it is not required.

/// yesnologger has nopWriter from io package, nillogger has nil writer

outcomes:
* Wrap takes a long time ... 3500 vs 1900 ns ... this could be done differently by having a premade wrapper ...
* noErr always takes the minimum time ~1.5ns in this case
* yesErr with nil checking still takes 2.5 times as long as yes/no even with a "nillogger," so the yes/no is worth it no matter what"

yesnologger.noErr(err=fake)-8         	774125499	         1.544 ns/op	       0 B/op	       0 allocs/op
nillogger.noErr(err=nil)-8            	774155878	         1.543 ns/op	       0 B/op	       0 allocs/op
yesnologger.yesErr(err=fake)-8        	   599269	      	  1856 ns/op	     456 B/op	      15 allocs/op
nillogger.yesErr(err=nil)-8           	339053805	         3.536 ns/op	       0 B/op	       0 allocs/op

/// standard logger with noop writer performing logging with and without options (wrap / msg)

outcomes:
* noErr takes ~2.8 ns in these examples, instead of ~1.5 in the original tests (noop writer vs io.nopWriter)
* wrap takes a long time ...

//	errorLoggerTestStruct = newTestStruct(true, "", nil, nil, nil)
//	wrapTestStruct        = newTestStruct(true, "", fakeSysCallError, nil, nil)
//	messageTestStruct     = newTestStruct(true, "fake test message", nil, nil, nil)

errorLoggerTestStruct noErr_(noop) 			427324900	         2.792 ns/op	       0 B/op	       0 allocs/op
wrapTestStruct noErr_(noop) 				427969722	         2.803 ns/op	       0 B/op	       0 allocs/op
messageTestStruct noErr_(noop) 				431078481	         2.785 ns/op	       0 B/op	       0 allocs/op

errorLoggerTestStruct yesErr_(real_error) 	   583864	      	  1875 ns/op	     456 B/op	      15 allocs/op
wrapTestStruct yesErr_(real_error) 			   331940	      	  3544 ns/op	     995 B/op	      23 allocs/op
messageTestStruct yesErr_(real_error) 		   617725	      	  1876 ns/op	     456 B/op	      15 allocs/op
*/

var testLoggers = []struct {
	name string
	log  ErrorLogger
}{
	{"yesnologger", yesnologger},
	{"nillogger", nil},
	{"disabled", NewWithOptions(false, "", nil, nil, nil)},
	{"wrap", NewWithOptions(true, "", nil, errFake, nil)},
	{"message", NewWithOptions(true, "message", nil, nil, nil)},
	{"wrap&message", NewWithOptions(true, "wrap&message", nil, errFake, nil)},
}

func Benchmark_errorLogger_noErr_yesErr(b *testing.B) {
	yesNoTests := []struct {
		name    string
		errName string
		input   error
		want    error
	}{
		{"yesnologger", "fake", errFake, errFake}, // return an error unchanged
		{"nillogger", "nil", nil, nil},            // return nil unchanged
	}

	for _, bb := range yesNoTests {
		name := fmt.Sprintf("%s.noErr(err=%s)", bb.name, bb.errName)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fakeOuter = errorLoggerTestStruct.noErr(bb.input)
			}
		})

		name = fmt.Sprintf("%s.yesErr(err=%s)", bb.name, bb.errName)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fakeOuter = errorLoggerTestStruct.yesErr(bb.input)
			}
		})
	}
}

func Benchmark_errorLogger_options(b *testing.B) {
	yesNoTests := []struct {
		name    string
		errName string
		input   error
		want    error
	}{
		{"yesnologger", "fake", errFake, errFake}, // return an error unchanged
		{"nillogger", "nil", nil, nil},            // return nil unchanged
	}
	for _, bb := range yesNoTests {
		name := fmt.Sprintf("%s.noErr(err=%s)", bb.name, bb.errName)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fakeOuter = errorLoggerTestStruct.noErr(bb.input)
			}
		})

		name = fmt.Sprintf("%s.yesErr(err=%s)", bb.name, bb.errName)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fakeOuter = errorLoggerTestStruct.yesErr(bb.input)
			}
		})
	}
}

func Test_errorLogger_noErr_yesErr(t *testing.T) {
	yesNoTests := []struct {
		name    string
		errName string
		input   error
		want    error
	}{
		{"yesnologger", "fake", errFake, errFake}, // return an error unchanged
		{"nillogger", "nil", nil, nil},            // return nil unchanged
	}

	for _, tt := range yesNoTests {
		t.Run(tt.name+".noErr", func(_ *testing.T) {
			fakeOuter = errorLoggerTestStruct.noErr(tt.input)
		})

		t.Run(tt.name+".yesErr", func(_ *testing.T) {
			fakeOuter = errorLoggerTestStruct.yesErr(tt.input)
		})

		t.Run(tt.name+"Err", func(_ *testing.T) {
			fakeOuter = errorLoggerTestStruct.Err(tt.input)
		})
	}
}

func Test_nopWriter_Write(t *testing.T) {
	tests := []struct {
		name    string
		n       io.Writer
		b       []byte
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := nopWriter
			gotN, err := n.Write(tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("nopWriter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("nopWriter.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
