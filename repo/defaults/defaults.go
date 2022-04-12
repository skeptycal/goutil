package defaults

import (
	"errors"
	"fmt"
	"strings"
)

const (
	defaultDebugState = true
	defaultTraceState = true
)

// Defaults provides a global set of defaults.
var Defaults DefaultMapper = NewDefaults(defaultDebugState, defaultTraceState)

func NewDefaults(debugState, traceState bool) DefaultMapper {
	m := defaultMap{
		m:          make(map[string]Setting, 2),
		debugState: debugState,
		traceState: traceState,
	}
	m.Set("debugState", debugState)
	m.Set("traceState", traceState)
	return &m
}

type (
	// defaultMap is the main storage for a DefaultMapper
	defaultMap struct {
		debugState bool
		traceState bool
		m          map[string]Setting
	}

	// DefaultMapper contains the defaults and settings for
	// an application. It has several defaults set ...
	// by default and any number of settings of any type
	// can be stored and retrieved.
	DefaultMapper interface {
		GetSetter
		Stringer
		IsDebug() bool
		IsTrace() bool
	}
)

func (d defaultMap) Set(key Any, value Any) error {
	switch v := key.(type) {
	case string:
		d.m[v] = NewSetting(v, value)
		return nil
	default:
		return fmt.Errorf("key type not string: %v", GetType(key))
	}
}

func (d defaultMap) Get(key Any) (Any, error) {

	switch v := key.(type) {
	case string:
		if v, ok := d.m[v]; ok {
			return v, nil
		}
		return nil, errors.New("key not found: " + v)
	default:
		return nil, fmt.Errorf("key type not string: %v", GetType(key))

	}
}

func (d defaultMap) String() string {

	const format = "%-20s = %-20v\n"
	sb := &strings.Builder{}
	defer sb.Reset()

	fmt.Fprint(sb, "Default Settings Map:\n")
	fmt.Fprintf(sb, format, "Key", "Value")

	for key, value := range d.m {
		fmt.Fprintf(sb, format, key, value)
	}

	return sb.String()
}

func (d defaultMap) IsDebug() bool { return d.m["debugState"].AsBool() }
func (d defaultMap) IsTrace() bool { return d.m["traceState"].AsBool() }
