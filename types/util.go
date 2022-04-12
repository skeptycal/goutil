package types

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/mattn/go-isatty"
)

// ReplacementChar is the recognized unicode replacement
// character for malformed unicode or errors in
// encoding.
//
// It is also found in unicode.ReplacementChar
const ReplacementChar rune = '\uFFFD'

const (
	UPPER    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LOWER    = "abcdefghijklmnopqrstuvwxyz"
	DIGITS   = "0123456789"
	ALPHA    = LOWER + UPPER
	ALPHANUM = ALPHA + DIGITS
)

var (

	// NoColor defines if the output is colorized or not. It's dynamically set to
	// false or true based on the stdout's file descriptor referring to a terminal
	// or not. It's also set to true if the NO_COLOR environment variable is
	// set (regardless of its value). This is a global option and affects all
	// colors. For more control over each color block use the methods
	// DisableColor() individually.
	//
	// Reference: color.NoColor from https://github.com/fatih/color
	noColor bool = noColorEnvExists() || os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))
	// NoColor = color.NoColor
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Intn(start, end int) int { return rand.Intn(start-end) - start }

func RandomString(n int) string {
	sb := strings.Builder{}
	defer sb.Reset()

	for i := 0; i < n; i++ {
		pos := rand.Intn(len(ALPHANUM) - 1)
		sb.WriteByte(ALPHANUM[pos])
	}

	return sb.String()
}

// noColorEnvExists returns true if the environment variable NO_COLOR exists.
//
// Reference: color.noColorExists from https://github.com/fatih/color
func noColorEnvExists() bool {
	_, exists := os.LookupEnv("NO_COLOR")
	return exists
}

func chr(c byte) string {
	return fmt.Sprintf("%c", c)
}

// WithLock runs fn while holding lk.
func WithLock(lk Locker, fn func()) {
	defer lk.Unlock() // in case fn panics
	lk.Lock()
	fn()
}

// IsComparable returns true if the underlying value
// is of a type that is capable of comparisions, e.g.
// equal, not equal
//
// Bools, strings and most numeric values are comparable.
//
// Next, types that have a Len() method are considered
// comparable by this function based on their length and
// item type alone. This is different from the standard
// library approach.
func IsComparable(a Any) bool {

	v := ValueOf(a)
	if v.Kind() == reflect.Interface {
		v = ValueOf(Interface(v))
	}
	k := v.Kind()

	return kindMaps[k].IsComparable()
}

// IsOrdered returns true if the underlying value is ordered.
// This means that it is capable of order based comparisons, e.g.
// less than, greater than
//
// Strings and most numeric values are ordered.
//
//
func IsOrdered(v Any) bool { return new_any(v).IsOrdered() }

// IsDeepComparable returns true if the underlying value
// is of a type that is capable of DeepEqual, the Go
// standard library approach to rigorous comparisons.
func IsDeepComparable(v Any) bool { return new_any(v).IsDeepComparable() }

// IsIterable returns true if the underlying value is
// made up of smaller units that can be read out one by
// one.
//
// Maps, strings, and slices naturally come to mind, but this
// package also adds functionality to iterate over most numeric
// values and structs.
func IsIterable(v Any) bool { return new_any(v).IsIterable() }

// HasAlternate returns true if the underlying value has
// alternate methods in addition to the Go standard library
// operations.
func HasAlternate(v Any) bool { return new_any(v).HasAlternate() }

// Contains returns true if the underlying iterable
// sequence (haystack) contains the search term
// (needle) in at least one position.
func Contains(needle Any, haystack []Any) bool {
	for _, x := range haystack {
		if reflect.DeepEqual(needle, x) {
			return true
		}
	}
	return false
}

// Count returns the number of times the search term
// (needle) occurs in the underlying iterable
// sequence (haystack).
func Count(needle Any, haystack []Any) int {
	retval := 0
	for _, x := range haystack {
		if reflect.DeepEqual(needle, x) {
			retval += 1
		}
	}
	return retval
}

// ToString converts the given argument to the
// standard string representation. If a implements
// fmt.Stringer, it is used, otherwise the slower
// fmt.Sprintf is used as a backup.
func ToString(a Any) string {
	if v, ok := a.(fmt.Stringer); ok {
		return v.String()
	}
	return fmt.Sprintf("%v", a)
}

func ToValues(list []Any) []reflect.Value {
	retval := make([]reflect.Value, 0, len(list))
	for _, item := range list {
		retval = append(retval, reflect.ValueOf(item))
	}
	return retval
}
