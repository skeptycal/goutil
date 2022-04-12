package defaults

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	// . "github.com/logrusorgru/aurora"
)

// DebugWriter is the default io.Writer used for debug output.
// The default is os.Stderr but may be set to any io.Writer,
// such as log output or a file path.
var DebugWriter DbWriter = NewDebugWriter(dbEchoPrefixBytes, defaultDbWriter)

func NewDebugWriter(p []byte, w io.Writer) DbWriter {
	return &dbWriter{
		prefix: p,
		w:      w,
	}
}

var (
	// dbEchoPrefixFormatString contains the ANSI escape codes to colorize
	// debug output
	dbEchoPrefix      = "\x1b[97;41m"
	dbEchoPrefixBytes = []byte(dbEchoPrefix)

	// reset contains the ANSI code to reset all colors and effects
	reset      = "\x1b[0m"
	resetBytes = []byte(reset)

	// dbWriter is the default writer used for debug output
	defaultDbWriter io.Writer = os.Stderr

	// errDebugInactive returns an error message if color output is disabled
	errDebugInactive error = errors.New("cannot write debug messages when Defaults.Debug is false")
)

// DbEcho is the shortcut to write ANSI wrapped text lines
// output using the default debug io.Writer
func DbEcho(args ...interface{}) (n int, err error) {
	return DebugWriter.Println(args...)
}

type (

	// DbWriter is an io.Writer implementation that only writes
	// to output if the defaults.Debug flag is set to true.
	DbWriter interface {
		io.Writer
		io.StringWriter
		Enabler
		Printer

		SetWriter(w io.Writer)
	}

	dbWriter struct {

		// prefix contains the ANSI code used to prefix all output
		prefix []byte

		// fn is the function used to write output. It is set to write
		// when the dbWriter is enabled and noWrite when disabled.
		//
		// This eliminates numerous flag checks which leads to
		// increased performance when many calls are made.
		fn func(b []byte) (n int, err error)

		// w is the io.Writer used for debug output
		w io.Writer
	}
)

func (w *dbWriter) Enable() {
	w.fn = w.write
}

func (w *dbWriter) Disable() {
	w.fn = w.noWrite
}

func (w *dbWriter) SetWriter(writer io.Writer) {
	w.w = writer
}

func (w *dbWriter) Write(p []byte) (n int, err error) {
	return w.fn(p)
}

func (w *dbWriter) noWrite(b []byte) (n int, err error) {
	return 0, errDebugInactive
}

// Write writes bytes to the output stream and returns
// the number of bytes written and any error encountered.
func (w *dbWriter) write(b []byte) (n int, err error) {
	return w.w.Write(b)
}

func (w *dbWriter) WriteString(s string) (n int, err error) {
	return w.Write([]byte(s))
}

func (w *dbWriter) Print(args ...interface{}) (n int, err error) {
	if Defaults.IsDebug() {
		return fmt.Fprint(w.w, args...)
	}
	return 0, errDebugInactive
}

func (w *dbWriter) Println(args ...interface{}) (n int, err error) {
	if Defaults.IsDebug() {
		return fmt.Fprintln(w.w, args...)
	}
	return 0, errDebugInactive
}

func (w *dbWriter) Printf(format string, args ...interface{}) (n int, err error) {
	if Defaults.IsDebug() {
		return fmt.Fprintf(w.w, format, args...)
	}
	return 0, errDebugInactive
}

func (w *dbWriter) print(args ...interface{}) (n int, err error) {
	var l = len(args)

	sb := strings.Builder{}
	defer sb.Reset()

	for i, arg := range args {
		switch t := arg.(type) {
		case []byte:
			sb.Write(t)
		case string:
			sb.WriteString(t)
		default:
			sb.WriteString(fmt.Sprintf("%v", arg))
		}
		if i < l {
			sb.Write([]byte(" "))
		}
	}

	w.Write(dbEchoPrefixBytes)
	n, err = w.WriteString(sb.String())
	w.Write(resetBytes)
	return
}

func (w *dbWriter) sprintf(args ...interface{}) (n int, err error) {
	w.Write(dbEchoPrefixBytes)
	n, err = w.WriteString(fmt.Sprint(args...))
	w.Write(resetBytes)
	return
}

func (w *dbWriter) println(args ...interface{}) (n int, err error) {
	n, err = w.Print(args...)
	w.WriteString("\n")
	return n + 1, err
}
