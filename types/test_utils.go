package types

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	LimitResult            bool
	DefaultTestResultLimit = 15
)

type (

	// Assert implements the Tester interface. It is
	// used for boolean only challenges. In addition
	// to working seamlessly with the standard library
	// testing package, it can return the bool
	// result for use in alternate data collection
	// or CI software.
	Assert interface {
		Tester
		Result() bool
	}

	// Random implements Tester and  creates a random
	// test that can be used to generate many varied
	// tests automatically.
	// After each use, Regenerate() can be called to
	// generate a new test.
	Random interface {
		Tester
		Regenerate()
	}

	// Custom implements Tester and can be used to
	// hook into existing software by passing in
	// the various test arguments with Hook().
	// Calling Hook() also calls Run() automaticaly.
	Custom interface {
		Tester
		Hook(name string, got, want Any, wantErr bool)
	}

	assert struct {
		name   string
		got    Any
		want   Any
		assert Assert
	}
)

// // Reset clears the list of tests
// func (ts *testSet) reset() {
// 	ts.list = []test{}
// }

func limitTestResultLength(v Any) string {
	s := fmt.Sprintf("%v", v)

	if len(s) > DefaultTestResultLimit && LimitResult {
		return s[:DefaultTestResultLimit-3] + "..."
	}

	return s
}

func TName(testname, funcname, argname Any) string {
	if argname == "" {
		return fmt.Sprintf("%v: %v()", testname, funcname)
	}
	return fmt.Sprintf("%v: %v(%v)", testname, funcname, argname)
}

func typeGuardExclude(needle Any, notAllowed []Any) bool {
	return !Contains(needle, notAllowed)
}

func TTypeError(t *testing.T, name string, got, want Any) {
	t.Errorf("%v = %v(%T), want %v(%T)", name, limitTestResultLength(got), got, limitTestResultLength(want), want)
}

func TError(t *testing.T, name string, got, want Any) {
	t.Errorf("%v = %v, want %v", name, limitTestResultLength(got), limitTestResultLength(want))
}
func TTypeRun(t *testing.T, name string, got, want Any) {
	if NewAnyValue(got).IsComparable() && NewAnyValue(want).IsComparable() {
		t.Run(name, func(t *testing.T) {
			if got != want {
				if !reflect.DeepEqual(got, want) {
					TTypeError(t, name, got, want)
				}
			}
		})
	}
}

func tRun(t *testing.T, name string, got, want Any) {
	if NewAnyValue(got).IsComparable() && NewAnyValue(want).IsComparable() {
		t.Run(name, func(t *testing.T) {
			if got != want {
				if !reflect.DeepEqual(got, want) {
					TError(t, name, got, want)
				}
			}
		})
	}
}
