package shpath

import (
	"bytes"
	"strings"
)

// NormalizeNL normalizes newlines to the popular
// standard '\n' convention.
//
// Any '\r\n' (a Windows convention) or '\r' (an
// older Apple convention) sequences are replaced
// with a single '\n'.
//
// Several methods are profiled and the most efficient
// one is aliased to this function.
var NormalizeNL = normalizeNewlinesString

func normalizeWrapped(s string) string {
	return normalizeNewlinesString(s)
}

// DropDupes replaces consecutive duplicates of only specific 'sep' strings with a single 'sep'
func DropDupes(s string, sep string) string {
	return strings.Replace(s, sep+sep, sep, -1)
}

/* ----------------------------------------------------------------
BenchmarkNormalize/for_loop-8                   2278214	       522.9 ns/op	     336 B/op	       3 allocs/op
BenchmarkNormalize/bytes.Replace-8              2078160	       574.4 ns/op	     448 B/op	       4 allocs/op
BenchmarkNormalize/strings.Replace-8            4335457	       276.2 ns/op	     224 B/op	       2 allocs/op
BenchmarkNormalize/wrapped-8                    4339246	       279.2 ns/op	     224 B/op	       2 allocs/op
PASS
*/

func normalizeStringsBuilder(s string) string {
	const old = "\r"
	const old2 = "\r\n"
	const new = "\n"
	var n = -1

	// if old == s || old2 == s {
	// 	return s // avoid allocation
	// }

	// Compute number of replacements.
	if m := strings.Count(s, old); m == 0 {
		if m = strings.Count(s, old2); m == 0 {
			return s // avoid allocation
		}
		// return s // avoid allocation
		// } else if n < 0 || m < n {
	} else if true || m < n {
		n = m
	}

	// Apply replacements to buffer.
	var b strings.Builder

	// in this use case  len(new) is always <= len(old)
	// b.Grow(len(s) + n*(len(new)-len(old)))
	b.Grow(len(s))

	start := 0
	for i := 0; i < n; i++ {
		j := start
		// in this use case, len(old) is always 1 or 2
		// if len(old) == 0 {
		// 	if i > 0 {
		// 		_, wid := utf8.DecodeRuneInString(s[start:])
		// 		j += wid
		// 	}
		// } else {
		j += strings.Index(s[start:], old)
		if j == -1 {
			strings.Index(s[start:], old2)
		}
		// }
		b.WriteString(s[start:j])
		b.WriteString(new)
		start = j + len(old)
	}
	b.WriteString(s[start:])
	return b.String()
}

var global_sb = strings.Builder{}

// reuse global string builder
func normalizeGlobalStringsBuilder(s string) string {
	global_sb.Reset()
	global_sb.Grow(len(s))

	var old int = 0

	for {
		if i := strings.Index(s[old:], "\r"); i > 0 {
			global_sb.WriteString(s[old:i])

			old = i + 1
			if s[i+1] == '\n' {
				old += 1
			}
		} else {
			break
		}
	}

	return global_sb.String()
}

// normalizeNewlines normalizes \r\n (windows) and \r (mac)
// into \n (unix)
//
// Reference: https://www.programming-books.io/essential/go/normalize-newlines-1d3abcf6f17c4186bb9617fa14074e48
func normalizeNewlines(d []byte) []byte {
	// replace CR LF \r\n (windows) with LF \n (unix)
	d = bytes.Replace(d, []byte{13, 10}, []byte{10}, -1)
	// replace CF \r (mac) with LF \n (unix)
	d = bytes.Replace(d, []byte{13}, []byte{10}, -1)
	return d
}

// normalizeNewlinesString uses strings.Replace to compare normalization efficiency
func normalizeNewlinesString(d string) string {
	// replace CR LF \r\n (windows) with LF \n (unix)
	d = strings.Replace(d, "\r\n", "\n", -1)
	// replace CF \r (mac) with LF \n (unix)
	d = strings.Replace(d, "\r", "\n", -1)
	return d
}

func normalizeNewlinesBytesWrapper(d string) string {
	return string(normalizeNewlines([]byte(d)))
}

func normalizeNLForLoop(s string) string {

	old := []byte(s)
	new := make([]byte, 0, len(old))

	for i, c := range old {
		switch c {
		case '\r':
			if old[i+1] == '\n' {
				continue
			}
			new = append(new, '\n')
		case '\n':
			new = append(new, '\n')
		default:
			new = append(new, c)

		}
	}
	return string(new)
}
