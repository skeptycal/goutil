package gostrings

import (
	"
)

var (
	log      = errorlogger.Log
	catchErr = errorlogger.Err
)

// Reverse copies the specified files to the standard output,
// reversing the order of characters in every line.  If no
// files are specified, the standard input is read.
//
// It replicates the macOS 12.1 'rev' utility available
// since  June 9, 1993.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// StrErrorConvert converts a function that returns
//  (string, error)
// to a function that logs the error and returns
//  (string)
func StrErrorConvert(s string, err error) string {
	// s, err := fn()
	catchErr(err)
	trace()
	return s
}
