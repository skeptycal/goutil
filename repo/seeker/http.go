package seeker

import (
	"net/http"
	"net/url"
)

// The following are type aliases for the Go
// standard library net/http package types.

type (
	URL      = url.URL
	File     = http.File
	Request  = http.Request
	Response = http.Response
	Header   = http.Header
)

type (
	// A Handler responds to an HTTP request.
	//
	// ServeHTTP should write reply headers and data to the ResponseWriter
	// and then return. Returning signals that the request is finished; it
	// is not valid to use the ResponseWriter or read from the
	// Request.Body after or concurrently with the completion of the
	// ServeHTTP call.
	//
	// Depending on the HTTP client software, HTTP protocol version, and
	// any intermediaries between the client and the Go server, it may not
	// be possible to read from the Request.Body after writing to the
	// ResponseWriter. Cautious handlers should read the Request.Body
	// first, and then reply.
	//
	// Except for reading the body, handlers should not modify the
	// provided Request.
	//
	// If ServeHTTP panics, the server (the caller of ServeHTTP) assumes
	// that the effect of the panic was isolated to the active request.
	// It recovers the panic, logs a stack trace to the server error log,
	// and either closes the network connection or sends an HTTP/2
	// RST_STREAM, depending on the HTTP protocol. To abort a handler so
	// the client sees an interrupted response but the server doesn't log
	// an error, panic with the value ErrAbortHandler.
	//
	//  type Handler interface {
	// 		ServeHTTP(ResponseWriter, *Request)
	// 	}
	Handler = http.Handler

	// A ResponseWriter interface is used by an HTTP handler to
	// construct an HTTP response.
	//
	// A ResponseWriter may not be used after the Handler.ServeHTTP method
	// has returned.
	//
	// 	type ResponseWriter interface {
	// 		// Header returns the header map that will be sent by
	// 		// WriteHeader. The Header map also is the mechanism with which
	// 		// Handlers can set HTTP trailers.
	// 		//
	// 		// Changing the header map after a call to WriteHeader (or
	// 		// Write) has no effect unless the modified headers are
	// 		// trailers.
	// 		//
	// 		// There are two ways to set Trailers. The preferred way is to
	// 		// predeclare in the headers which trailers you will later
	// 		// send by setting the "Trailer" header to the names of the
	// 		// trailer keys which will come later. In this case, those
	// 		// keys of the Header map are treated as if they were
	// 		// trailers. See the example. The second way, for trailer
	// 		// keys not known to the Handler until after the first Write,
	// 		// is to prefix the Header map keys with the TrailerPrefix
	// 		// constant value. See TrailerPrefix.
	// 		//
	// 		// To suppress automatic response headers (such as "Date"), set
	// 		// their value to nil.
	// 		Header() Header
	//
	// 		// Write writes the data to the connection as part of an HTTP reply.
	// 		//
	// 		// If WriteHeader has not yet been called, Write calls
	// 		// WriteHeader(http.StatusOK) before writing the data. If the Header
	// 		// does not contain a Content-Type line, Write adds a Content-Type set
	// 		// to the result of passing the initial 512 bytes of written data to
	// 		// DetectContentType. Additionally, if the total size of all written
	// 		// data is under a few KB and there are no Flush calls, the
	// 		// Content-Length header is added automatically.
	// 		//
	// 		// Depending on the HTTP protocol version and the client, calling
	// 		// Write or WriteHeader may prevent future reads on the
	// 		// Request.Body. For HTTP/1.x requests, handlers should read any
	// 		// needed request body data before writing the response. Once the
	// 		// headers have been flushed (due to either an explicit Flusher.Flush
	// 		// call or writing enough data to trigger a flush), the request body
	// 		// may be unavailable. For HTTP/2 requests, the Go HTTP server permits
	// 		// handlers to continue to read the request body while concurrently
	// 		// writing the response. However, such behavior may not be supported
	// 		// by all HTTP/2 clients. Handlers should read before writing if
	// 		// possible to maximize compatibility.
	// 		Write([]byte) (int, error)
	//
	// 		// WriteHeader sends an HTTP response header with the provided
	// 		// status code.
	// 		//
	// 		// If WriteHeader is not called explicitly, the first call to Write
	// 		// will trigger an implicit WriteHeader(http.StatusOK).
	// 		// Thus explicit calls to WriteHeader are mainly used to
	// 		// send error codes.
	// 		//
	// 		// The provided code must be a valid HTTP 1xx-5xx status code.
	// 		// Only one header may be written. Go does not currently
	// 		// support sending user-defined 1xx informational headers,
	// 		// with the exception of 100-continue response header that the
	// 		// Server sends automatically when the Request.Body is read.
	// 		WriteHeader(statusCode int)
	// 	}
	ResponseWriter = http.ResponseWriter
	Flusher        = http.Flusher
	Hijacker       = http.Hijacker
	CloseNotifier  = http.CloseNotifier
	Pusher         = http.Pusher
)
