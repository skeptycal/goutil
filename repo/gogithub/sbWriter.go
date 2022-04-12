package gogithub

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"sync"
)

type (
	swimmer sbWriter
	sbPool  map[*swimmer]bool
)

func (p sbPool) release(s *swimmer) {
	p[s] = false
}

func (p sbPool) new() *swimmer {
	return nil
}

type (
	sbWriter struct {
		sb *strings.Builder
		mu *sync.Mutex
	}

	SBWriter interface {
		Println(a ...interface{}) (n int, err error)
		Reset()
		io.Writer
		fmt.Stringer
	}
)

// NewSBWriter returns a new SBWriter that implements
// io.Writer to write to a strings.Builder. These are
// used for building temporary strings and will be
// truncated when the final string is returned.
//
// The SBWriter retains its allocated capacity and can
// be reused. A sync.Mutex provides locks and protects
// against concurrent writes and it is safe for concurrent
// use.
func NewSBWriter() SBWriter {
	sbw := sbWriter{}
	sbw.mu.Lock()
	return &sbw
}

// String returns the string representation of the
// buffer and resets the Builder to zero length.
func (w *sbWriter) String() string {
	defer w.sb.Reset()
	return w.sb.String()
}

// Reset resets the Builder to be empty and
// releases the mutex lock.
func (w *sbWriter) Reset() {
	defer w.mu.Unlock()
	w.sb.Reset()
}

func (w *sbWriter) Write(p []byte) (n int, err error) {
	return w.sb.Write(p)
}

func (w *sbWriter) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, a...)
}

func (w *sbWriter) WriteLines(p []byte) (n int, err error) {
	// this is kinda redundant and silly ... but ...
	// it is a placeholder for another parsing function
	lines := bytes.Split(p, NL)
	for _, line := range lines {
		n, err := w.Write(line)
		if err != nil {
			return n, err
		}
		_, err = w.Write(NL)
		if err != nil {
			return n, err
		}
		n += 1
	}
	if len(p) != n {
		return n, io.ErrShortWrite
	}
	return n, nil
}
