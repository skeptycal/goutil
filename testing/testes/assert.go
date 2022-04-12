package testes

import (
	"fmt"
	"reflect"
	"testing"

	"
)

var NewAnyValue = types.NewAnyValue

const assertFmtSuffix = "(%v): got %v, want %v"

type assertFMT string

const (
	assertEqual     assertFMT = "AssertEqual"
	assertNotEqual  assertFMT = "AssertNotEqual"
	assertDeepEqual assertFMT = "AssertDeepEqual"
	assertSameType  assertFMT = "AssertSameType"
	assertSameKind  assertFMT = "AssertSameKind"
	assertSameFunc  assertFMT = "AssertSameFunc"
)

func (a assertFMT) String() string {
	return string(a) + assertFmtSuffix
}

func AssertErrorf(t *testing.T, formatString assertFMT, name string, got, want Any) {
	if formatString == "" {
		formatString = "%v " + assertFmtSuffix // "%v = %v(%T), want %v(%T)"
	}
	t.Errorf(formatString.String(), name, limitTestResultLength(got), got, limitTestResultLength(want), want)
}

func IsKindEqual(got, want Any) bool {
	if gv, ok := got.(reflect.Kind); ok {
		if wv, ok := want.(reflect.Kind); ok {
			return gv == wv
		}
	}
	return false
}

func AssertKindEqual(t *testing.T, name string, got, want Any) bool {
	if IsKindEqual(got, want) {
		return true
	}
	AssertErrorf(t, assertSameKind, name, got, want)
	return false
}

func GetFuncResult(t *testing.T, name string, fn Any, args ...reflect.Value) ([]reflect.Value, error) {
	if fn == nil {
		return nil, fmt.Errorf("fn must be provided: %v", fn)
	}

	f := NewAnyValue(fn)

	if !IsKindEqual(f.Kind(), reflect.Func) {
		return nil, fmt.Errorf("fn must be a function: %v(%v)", fn, f.Kind())
	}

	return nil, nil
}

// AssertSameFunc returns true if got and want are
// both functions that return the same value when
// called with args... as input.
func AssertSameFunc(t *testing.T, name string, got, want Any, args ...reflect.Value) bool {

	g := NewAnyValue(got)

	// If got is a pointer, run again with object
	if IsKindEqual(g.Kind(), reflect.Ptr) {
		return AssertSameFunc(t, name, reflect.Indirect(g.ValueOf()), want)
	}

	// If want is a pointer, run again with object
	w := NewAnyValue(want)
	if w.Kind() == reflect.Ptr {
		return AssertSameFunc(t, name, got, reflect.Indirect(w.ValueOf()))
	}
	if g.Kind() != reflect.Func {
		AssertErrorf(t, assertSameFunc, name, g.Kind(), reflect.Func)
		return false
	}

	if g.Kind() != reflect.Func {
		AssertErrorf(t, assertSameFunc, name, g.Kind(), reflect.Func)
		return false
	}

	gf := g.ValueOf().Call(args)
	wf := w.ValueOf().Call(args)

	return AssertEqual(t, name, gf, wf)

}

func AssertEqual(t *testing.T, name string, got, want Any) bool {
	if got == want {
		return true
	}
	AssertErrorf(t, assertEqual, name, got, want)
	return false
}

func AssertNotEqual(t *testing.T, name string, got, want Any) bool {
	if got != want {
		return true
	}
	AssertErrorf(t, assertNotEqual, name, got, want)
	return false
}

func AssertDeepEqual(t *testing.T, name string, got, want Any) bool {
	if reflect.DeepEqual(got, want) {
		return true
	}
	AssertErrorf(t, assertDeepEqual, name, got, want)
	return false
}

func AssertSameType(t *testing.T, name string, got, want Any) bool {
	g := NewAnyValue(got).TypeOf()
	w := NewAnyValue(want).TypeOf()

	if g == w {
		return true
	}
	AssertErrorf(t, assertSameType, name, g, w)
	return false
}

func AssertSameKind(t *testing.T, name string, got, want Any) bool {
	g := NewAnyValue(got).Kind()
	w := NewAnyValue(want).Kind()

	if g == w {
		return true
	}
	AssertErrorf(t, assertSameKind, name, g, w)
	return false
}
