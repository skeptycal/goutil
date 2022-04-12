package benchmark

import (
	"bufio"
	"bytes"
	"io"
	"net/http/httptest"
	"os"
)

var (
	// defaultBenchmarkOutput is os.Stderr
	defaultBenchmarkOutput io.Writer = os.Stderr

	// benchmarkBuffer implements io.Writer and is
	// used to store output of benchmark until all
	// steps have finished.
	//
	// This eliminates the processing or lag that
	// may occur due to the writing of data.
	//
	// Upon successful benchmark completion, the
	// buffer contents are sent to the output.
	benchmarkBuffer *bytes.Buffer = new(bytes.Buffer)

	benchmarkResponse = httptest.NewRecorder()
)

type (
	Any = interface{}
	any = struct{}
)

// TODO - CLI mode

// RunBenchmark initializes the benchmark and processes
// the given tests. Command line arguments may be passed
// when used in CLI mode.
func RunBenchmark(b Benchmarks) {

	w := NewWriter(b.config.output)

	n, err := w.WriteString("test")

	verbosef("RunBenchmark return value: %v, err: %v", n, err)

}

// NewWriter returns a valid buffered StringWriter (bufio.Writer)
// based on w. The bufio.Writer interface implements Writer,
// StringWriter, ReadFrom, WriteByte, WriteRune, etc.
//
// If w is nil, w is reassigned defaultBenchmarkOutput (os.Stdout).
//
// This io.Writer is wrapped in a bufio.Writer and returned:
//  bufio.NewWriter(w)
//
// (from bufio.Writer)
//
// Writer implements buffering for an io.Writer object.
// If an error occurs writing to a Writer, no more data will be
// accepted and all subsequent writes, and Flush, will return the error.
// After all data has been written, the client should call the
// Flush method to guarantee all data has been forwarded to
// the underlying io.Writer.
func NewWriter(w io.Writer) *bufio.Writer {
	if w == nil {
		return bufio.NewWriter(defaultBenchmarkOutput)
	}
	return bufio.NewWriter(w)
}
