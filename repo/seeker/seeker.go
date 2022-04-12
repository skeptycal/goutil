// package seeker is a simplified web browser
// implementation used to search specifically
// for text and numerical data.
//
// This implementation is based on a "read-only"
// style of search where the majority of user
// input, animations, and other functionality
// is disabled.
//
// The intention is to provide:
// - a simple interface for reading data in a
// web browser with minimal distractions
//
// - saving data to a file, database, or other
// storage in convenient formats.
//
// - directly saving "data-type" files
//
// - uses Google Search to provide results
//
// - linked to Google Cloud Storage and related
// services (scripting to manipulate data in the
// cloud not yet supported.)
//
// - maintain hyperlinks by default.
//
// - simple option to turn on image loading.
//
// - options to control the style sheet used
// to display all pages.
//
// - command line options and configuration
//
// NOTE: Scripting and automation features of web
// pages are disabled by default. Most functionality
// of web pages is unavailable be design.
package seeker

import (
	"fmt"
	"net/http"
	"time"
)

type (
	style struct {
		name string
	}
	Styler interface {
		Parse([]byte) ([]byte, error)
	}
	page struct {
		url        string
		lastAccess time.Time
		customCSS  Styler // nil means use default
	}
)

var defaultCSS Styler = nil
var loc time.Location

func Config(portString string, handlers HandlerMap) *Settings {
	return &Settings{
		portString: portString,
		handlers:   handlers,
	}
}

func Seek(portString string) error {
	fmt.Printf("portString: %v\n", portString)
	return http.ListenAndServe(portString, nil)
}
