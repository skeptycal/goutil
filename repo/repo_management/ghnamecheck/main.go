package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"unicode"

	. "github.com/skeptycal/goutil/repo/repo_management/config" // _ "config"
)

// IsFile reports whether m describes a regular file.
// That is, it tests that no mode type bits are set.
func IsFile(filename string) bool {
	f, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return f.Mode().IsRegular()
}

// IsDir reports whether m describes a directory.
// That is, it tests for the ModeDir bit being set in m.
func IsDir(filename string) bool {
	f, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return f.Mode().IsDir()
}

var fakelist = []string{
	"fake_list",
	"asdf",
	"asd;fjlk",
	"poiupowne",
	"<<--slauvhwenusyas--slauvhwenusyas--slauvhwenusyas--slauvhwenusyas--slauvhwenusyas--slauvhwenusyas-->>",
}

// namecheck checks whether a string is a valid GitHub
// repository name.
// If invalid characters are found, they are replaced with '-'.
// If the length of the string is greater than allowed,
// the empty string is returned.
//
// GitHub Repository naming conventions:
//
// - Max length: 100 code points
//
// - All code points must be either a hyphen (-), an underscore (_),
// a period (.), or an ASCII alphanumeric code point
//
// - Must be unique per-user and/or per-organization
//
// Note: sequences of invalid code points are automatically replaced by a single hyphen (-)
//
// Note: length checking is performed after replacement
//
// This was verified through checking automatically-generated aliases with repository names.
//
// Reference: https://github.com/isiahmeadows/github-limits
func namecheck(name string) string {

	var retval []rune

	for _, r := range name {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_' || r == '.' {
			// add to string
			retval = append(retval, r)
			continue
		}
		retval = append(retval, '-')
		// add a single hyphen (-)
	}

	return string(retval)
}

type (
	Config interface {
		FileSaver
		SyncMapper
	}
	config struct {
		name     string // name of configuration
		filename string // location of the configuration file

		// sync.Map is a map of configuration settings
		// Map is like a Go map[interface{}]interface{} but is
		// safe for concurrent use by multiple goroutines without
		// additional locking or coordination. Loads, stores, and
		// deletes run in amortized constant time.
		//
		// The Map type is optimized for two common use cases,
		// one of which is when the entry for a given key is
		// only ever written once but read many times...
		//
		// In these cases, use of a Map may significantly reduce
		// lock contention compared to a Go map paired with a
		// separate Mutex or RWMutex.
		sync.Map
	}
)

// NewConfig creates a new configuration object. It is
// stored in the configuration file:
//  ~/.<name>/<name>_config.json
func NewConfig(name string) (Config, error) {
	c := config{name: name}

	// create directory if it doesn't exist'
	directory := filepath.Join(os.Getenv("HOME"), "."+name)
	err := os.MkdirAll(directory, ModeDir)
	if err != nil {
		return nil, err
	}

	// open or create file if it doesn't exist
	filename := filepath.Join(directory, name+"_config.json")
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, ModeFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	c.filename = filename

	return &c, nil
}

func (c *config) Yaml() []byte {
	return nil
}

func (c *config) ReadFile() error {
	b, err := os.ReadFile(c.filename)
	if err != nil {
		return err
	}
	_ = b
	return os.ErrInvalid
}

func (c *config) WriteFile() error {
	return os.WriteFile(c.filename, c.Yaml(), ModeFile)
	// return os.ErrInvalid
}

// BreakIt performs automatic line breaking of text
// based on the config.BreakItMax property.
func BreakIt(s string) string {

	if len(s) < BreakItMax {
		return s
	}

	lastSpace := -1
	b := []byte(s)

	for i, c := range b {
		if c == Space {
			lastSpace = i
		}
		if i%BreakItMax == 0 {
			b[lastSpace] = NL
		}
	}

	return string(b)
}

func main() {
	arglist := os.Args[1:]

	//! TODO - REMOVE AFTER TESTING
	if len(arglist) < 2 {
		arglist = fakelist
	}

	maxWidth := MaxItemLen(arglist)

	if maxWidth > 100 {
		maxWidth = 100
	}

	fmtString := fmt.Sprintf(" %%%d.%ds   %%%d.%ds\n", maxWidth, maxWidth, maxWidth, maxWidth)

	for _, arg := range arglist {

		fmt.Printf(fmtString, arg, namecheck(arg))
	}
}
