// Package Parser provides utilities to automate common text
// file parsing and editing tasks. It is designed to be used
// with MemFile, the in memory file system.
//
// The main interface is Parser which implements
// Parse() (string, error):
// 	type Parser interface {
// 		Parse() (string, error)
// 	}
//
// Once all options are set, Parse does the work and returns
// the data as a string. Any other file or data handling
// structures may choose to implement Parser to gain access
// the functionality of this package.
package parser
