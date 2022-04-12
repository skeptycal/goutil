package ghshell

import (
	"errors"
	"os"
	"strings"
)

func IsDir(path string) bool {
	if fi, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) && fi.IsDir() {
		return true
	}
	return false
}

func Exists(path string) bool {
	if fi, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) && fi.Mode().IsRegular() {
		return true
	}
	return false
}

// ToFields converts any number of string
// arguments into a slice of fields, as
// defined by
//
//  strings.Fields()
//
// from the standard library, and returns
// a slice of substrings of s or an empty
// slice if s contains only white space.
func ToFields(in ...string) []string {
	s := strings.Join(in, " ")
	return strings.Fields(s)
}

// NormalizeWhitespace removes any unicode
// whitespace characters from s and returns
// a string with all words separate by
// single spaces (0x20). As a convenience,
// there is an option to ignore newlines
// (0xA) and leave them intact.
//
// Whitespace is a space character as defined
// by Unicode's White Space property; in the
// Latin-1 space this is
//
//	'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP)
//
// Other definitions of spacing characters
// are set by category Z and property
// Pattern_White_Space.
func NormalizeWhitespace(s string, saveNewLines bool) string {
	if saveNewLines {
		list := strings.Split(s, "\n")
		new := make([]string, 0, len(s))
		for _, line := range list {
			new = append(new, NormalizeWhitespace(line, false))
		}
		return strings.Join(new, " ")
	}
	// uses unicode.IsSpace(r rune)
	return strings.Join(ToFields(s), " ")
}
