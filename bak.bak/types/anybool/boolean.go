package anybool

import (
	"math"
	"reflect"
	"strconv"

	"github.com/skeptycal/goutil/types"
)

type (
	Any      = types.Any
	Stringer = types.Stringer
	Enabler  = types.Enabler
	Cosa     = types.Cosa
)

func IsTrue(v Any) bool  { return AnyBooler(v).AsBool() }
func IsFalse(v Any) bool { return !AnyBooler(v).AsBool() }

// AnyBooler returns a new anyBool value that implements Booler.
// If the default value is a bool, the returned value will be
// a NewBooler that is much more efficient than AnyBooler.
func AnyBooler(defaultValue Any) Booler {
	switch v := defaultValue.(type) {
	case bool:
		return &boolean{v}
	case Booler:
		return v
	default:
		return &anyBool{v}
	}
}

// NewBooler returns a new boolean value that implements Booler.
// When the default value is always a bool, this is approx.
// 1 - 2 orders of magnitude faster than AnyBooler
func NewBooler(v bool) Booler { return &boolean{bool: v} }

// Booler represents a boolean value that implements Enabler
// and Stringer interfaces.
type Booler interface {
	Enabler
	Stringer
	AsBool() bool
}

// boolean represents a boolean value that implements Booler.
// It is the bool value only version of anyBool and is
// approx 1 - 2 orders of magnitude faster than anyBool.
type boolean struct{ bool }

func (b *boolean) Enable()      { b.bool = true }
func (b *boolean) Disable()     { b.bool = false }
func (b *boolean) AsBool() bool { return b.bool }
func (b *boolean) String() string {
	if b.bool {
		return "true"
	}
	return "false"
}

// anyBool represents any type of value that can be converted
// to a boolean to implement the Booler interfaces.
type anyBool struct{ any Any }

// Enable sets the underlying value to "a true value" as compatible
// with the original type as possible.
func (b *anyBool) Enable() {
	switch b.any.(type) {
	case int, uint:
		b.any = 1
	case float32, float64:
		b.any = 1.0
	case complex64, complex128:
		b.any = complex(math.Pi, -1)
	case string:
		b.any = "true"
	case []byte:
		b.any = []byte("true")
	case bool:
		// this case is not normally used ... the constructor assigns
		// natural bools to *boolean instead of *anyBool
		b.any = true
	default:
		b.any = true
	}
}

// Disable sets the underlying value to "a false value" as compatible
// with the original type as possible.
func (b *anyBool) Disable() {
	switch b.any.(type) {
	case bool:
		b.any = false
	case int, uint:
		b.any = 0
	case float32, float64:
		b.any = 0.0
	case complex64, complex128:
		b.any = complex(0, 0)
	case string:
		b.any = ""
	case []byte:
		b.any = []byte{}
	default:
		b.any = nil
	}
}

func (b *anyBool) String() (s string) {
	if b.AsBool() {
		return "true"
	}
	return "false"
}

func (b *anyBool) AsBool() bool {
	if b.any == nil {
		return false
	}

	v := reflect.ValueOf(b.any)
	k := v.Kind()

	if k == reflect.Invalid {
		return false
	}

	if k == reflect.Ptr {
		return !v.IsNil()

	}

	if k == reflect.Bool {
		return v.Bool()
	}

	if k == reflect.UnsafePointer {
		return false
	}

	if v.IsZero() {
		return false
	}

	if k == reflect.String {
		if ok, err := strconv.ParseBool(v.String()); err != nil {
			return ok
		}
		if s := b.any.(string); s == "" || s == "false" || s == "0" || s == "False" || s == "no" {
			return false
		}
		return true
	}

	// Kinds 2 - 6 are ints
	if k > 1 && k < 7 {
		return v.Int() != 0
	}

	// Kinds 7 - 11 are uints
	if k > 6 && k < 12 {
		return v.Uint() != 0
	}

	if k == reflect.Float32 || k == reflect.Float64 {
		return v.Float() != 0.0
	}

	if k == reflect.Complex64 || k == reflect.Complex128 {
		return v.Complex() != 0.0
	}

	if k == reflect.Array || k == reflect.Map || k == reflect.Chan || k == reflect.Slice || k == reflect.String {
		return v.Len() != 0
	}

	if k == reflect.Func {
		return !v.IsNil()
	}

	if k == reflect.Struct {
		return v.NumField() != 0
	}

	return false
}
