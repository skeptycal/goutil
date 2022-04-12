package testutils

import (
	"fmt"
	"io"
	"sync"

	"
	"
)

const (
	NL    = types.NL    // "\n"
	TAB   = types.TAB   // "\t"
	SPACE = types.SPACE // " "
)

// AnyBooler returns a new anyBool value that implements Booler. If the default value is a bool, the returned value will be a NewBooler that is much more efficient than AnyBooler.
var AnyBooler = anybool.AnyBooler

type (

	// Any represents a object that may contain any
	// valid type.
	//  type Any interface{}
	Any = types.Any

	// An Enabler represents an object that can be enabled or disabled.
	// 	type Enabler interface {
	// 		Enable()
	// 		Disable()
	// 	}
	Enabler = types.Enabler

	// A GetSetter represents an object that can be accessed using
	// Get and Set methods to access an underlying map.
	// 	type GetSetter interface {
	// 		Get(key Any) (Any, error)
	// 		Set(key Any, value Any) error
	// 	}
	GetSetter = types.GetSetter

	// A Printer implements common printing functions.
	//  type Printer interface {
	// 		Print(args ...interface{}) (n int, err error)
	// 		Println(args ...interface{}) (n int, err error)
	// 		Printf(format string, args ...interface{}) (n int, err error)
	//  }
	Printer = types.Printer

	Booler = anybool.Booler

	//********* Interfaces from standard library for reference.

	// Stringer is implemented by any value that has a String method,
	// which defines the ``native'' format for that value.
	// The String method is used to print values passed as an operand
	// to any format that accepts a string or to an unformatted printer
	// such as Print.
	//
	// Ref: fmt package
	//  type Stringer interface {
	// 		String() string
	//  }
	Stringer = fmt.Stringer

	// State represents the printer state passed to custom formatters.
	// It provides access to the io.Writer interface plus information about
	// the flags and options for the operand's format specifier.
	//
	// Ref: fmt package
	// 	type State interface {
	// 		// Write is the function to call to emit formatted output to be printed.
	// 		Write(b []byte) (n int, err error)
	// 		// Width returns the value of the width option and whether it has been set.
	// 		Width() (wid int, ok bool)
	// 		// Precision returns the value of the precision option and whether it has been set.
	// 		Precision() (prec int, ok bool)
	// 		// Flag reports whether the flag c, a character, has been set.
	// 		Flag(c int) bool
	// 	}
	State = fmt.State

	// Formatter is implemented by any value that has a Format method.
	// The implementation controls how State and rune are interpreted,
	// and may call Sprint(f) or Fprint(f) etc. to generate its output.
	//
	// Ref: fmt package
	//	 type Formatter interface {
	//	 	Format(f State, verb rune)
	//	 }
	Formatter = fmt.Formatter

	// A Locker represents an object that can be locked and unlocked.
	//
	// Ref: sync package
	//	 type Locker interface {
	//	 	Lock()
	//	 	Unlock()
	//	 }
	Locker = sync.Locker

	// Writer is the interface that wraps the basic Write method.
	//
	// Write writes len(p) bytes from p to the underlying data stream.
	// It returns the number of bytes written from p (0 <= n <= len(p))
	// and any error encountered that caused the write to stop early.
	// Write must return a non-nil error if it returns n < len(p). Write
	// must not modify the slice data, even temporarily.
	//
	// Implementations must not retain p.
	//
	// Ref: io package
	// 	type Writer interface {
	// 		Write(p []byte) (n int, err error)
	// 	}
	Writer = io.Writer

	// StringWriter is the interface that wraps the WriteString method.
	//
	// WriteString writes the contents of the string s to w, which accepts a slice of bytes.
	// If w implements StringWriter, its WriteString method is invoked directly.
	// Otherwise, w.Write is called exactly once.
	//
	// Ref: io package
	// 	type StringWriter interface {
	// 		WriteString(s string) (n int, err error)
	// 	}
	StringWriter = io.StringWriter
)
