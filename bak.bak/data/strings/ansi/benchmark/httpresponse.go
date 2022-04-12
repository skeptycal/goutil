package benchmark

import (
	"net/http/httptest"
)

// httptest.ResponseRecorder.Body is a bytes.Buffer
// Reference: https://stackoverflow.com/a/13766062

type ResponseRecorder httptest.ResponseRecorder

// Write appends the contents of p to the buffer, growing the buffer as
// needed. The return value n is the length of p; err is always nil. If the
// buffer becomes too large, Write will panic with ErrTooLarge.
func (r *ResponseRecorder) Write(p []byte) (int, error) {
	return r.Body.Write(p)
}

// WriteString appends the contents of s to the buffer, growing the buffer as
// needed. The return value n is the length of s; err is always nil. If the
// buffer becomes too large, WriteString will panic with ErrTooLarge.
func (r *ResponseRecorder) WriteString(s string) (n int, err error) {
	return r.Body.WriteString(s)
}

// Read reads the next len(p) bytes from the buffer or until the buffer
// is drained. The return value n is the number of bytes read. If the
// buffer has no data to return, err is io.EOF (unless len(p) is zero);
// otherwise it is nil.
func (r *ResponseRecorder) Read(p []byte) (n int, err error) {
	return r.Body.Read(p)
}

// ReadString reads until the first occurrence of delim in the input,
// returning a string containing the data up to and including the delimiter.
// If ReadString encounters an error before finding a delimiter,
// it returns the data read before the error and the error itself (often io.EOF).
// ReadString returns err != nil if and only if the returned data does not end
// in delim.
func (r *ResponseRecorder) ReadString(delim byte) (line string, err error) {
	return r.Body.ReadString(delim)
}

// Bytes returns a slice of length b.Len() holding the unread portion of the buffer.
// If the buffer is empty, nil is returned.
//
// The slice is valid for use only until the next buffer modification (that is,
// only until the next call to a method like Read, Write, Reset, or Truncate).
// The slice aliases the buffer content at least until the next buffer modification,
// so immediate changes to the slice will affect the result of future reads.
func (r *ResponseRecorder) Bytes() []byte {
	if r.Body.Len() < 1 {
		return nil
	}
	return r.Body.Bytes()
}

// String returns the contents of the unread portion of the buffer
// as a string. If the Buffer is a nil pointer, it returns "<nil>".
//
// To build strings more efficiently, see the strings.Builder type.
func (r *ResponseRecorder) String() string {
	return r.Body.String()
}
