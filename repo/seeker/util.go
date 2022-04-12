package seeker

import "net/http"

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
//
//  type HandlerFunc func(ResponseWriter, *Request)
type HandlerFunc = http.HandlerFunc

// DetectContentType implements the algorithm described
// at https://mimesniff.spec.whatwg.org/ to determine the
// Content-Type of the given data. It considers at most the
// first 512 bytes of data.
//
// DetectContentType always returns a valid MIME type: if
// it cannot determine a more specific one, it
// returns "application/octet-stream".
//
// http.DetectContentType
//  DetectContentType = func(data []byte) string
func DetectContentType(data []byte) string {
	return http.DetectContentType(data)
}

// StatusText returns a text for the HTTP status code. It returns the empty string if the code is unknown.
func StatusText(code int) string {
	return http.StatusText(code)
}
