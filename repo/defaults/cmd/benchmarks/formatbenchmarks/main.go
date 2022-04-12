package main

import (
	"bytes"
	"log"
	"os"
	"unicode"

	"github.com/skeptycal/goutil/repo/defaults"
)

const (
	defaultInputFilename = "in.txt"
	NL                   = "\n"
)

var Defaults = defaults.Defaults

type TextFile interface {
	ReadFile() error
	WriteFile() error
}

type textFile struct {
	filename string
}

func main() {

	var (
		filename string = defaultInputFilename
		nl              = []byte(NL)
		tab             = []byte("\t")
	)

	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := bytes.Split(b, nl)

	// parse fields
	for _, line := range lines {

		fields := bytes.Split(line, tab)

		// TrimSpace
		for i, field := range fields {
			fields[i] = bytes.TrimSpace(field)
		}

		// Find Numbers

	}

}

func Remove(slice []byte, i int) []byte {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func TrimSpace(b []byte) []byte {
	return bytes.TrimSpace(b)
}

type filterfunc func(c byte) bool

var filterNoSpace = func(c byte) bool { return c != ' ' }
var filterUnicodeDigit filterfunc = func(c byte) bool { return unicode.IsDigit(rune(c)) }
var filterDigit filterfunc = func(c byte) bool { return c > '0' && c < '9' }
var filterUnicodeNumeric filterfunc = func(c byte) bool { return unicode.IsDigit(rune(c)) || c == ',' || c == '.' }

// SubSequence returns the first slice containing only elements
// determined by the filter function.
func SubSequence(b []byte, filter func(c byte) bool) []byte {
	start := 0
	end := len(b)

	if filter == nil {
		filter = filterNoSpace
	}

	for i := 0; i < len(b); i++ {
		if filter(b[i]) {
			start = i
			break
		}
	}

	for i := start; i < len(b); i++ {
		if !filter(b[i]) {
			end = i
			break
		}
	}

	return b[start:end]
}
