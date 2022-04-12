package types

import (
	"reflect"
	"testing"
)

type (
	testFunc = func(t *testing.T)

	test struct {

		// required inputs
		t         *testing.T
		name      string
		got, want Any
		wantErr   bool

		// jit values
		g, w AnyValue

		// which assertion to use ... default is assertEqual
		assert testFunc
	}

	funcTest struct {
		test
		fn Any
		in Any
	}

	// Tester implements an individual test. It may
	// be implemented by traditional tests,
	// asserts, random inputs, or custom code.
	Tester interface {
		// Run runs an individual test.
		Run()
	}

	TestRunner interface {
		// Run runs all tests in the set.
		Tester
	}

	FunctionCaller interface {
		// Run runs a function and compares the results
		Tester
	}
)

func (s *test) WantErr() bool { return s.wantErr }

func NewTest(t *testing.T, name string, got, want Any, wantErr bool, assert testFunc) Tester {

	if got == nil && want == nil {
		return nilTest{}
	}

	s := test{
		t:       t,
		name:    name,
		got:     got,
		want:    want,
		wantErr: wantErr,
		assert:  assert,
	}

	if s.assert == nil {
		s.assert = s.assertEqual
	}

	return &s
}

// Run runs the individual test
func (s *test) Run() {
	s.t.Run(s.name, s.assertEqual)
}

func (s *test) Errorf(format string, args ...interface{}) {
	s.t.Errorf(format, args...)
}

func (s *test) LogType(msg string, g, w Any) {
	s.t.Logf("%v, got %v(%T), want %v(%T)", msg, g, g, w, w)
}
func (s *test) Logf(msg string, g, w Any) {
	s.t.Logf("%v, got %v, want %v", msg, g, w)
}

func (s *test) GV() reflect.Value {
	return s.Got().ValueOf()
}

func (s *test) WV() reflect.Value {
	return s.Want().ValueOf()
}

func (s *test) Got() AnyValue {
	if s.g == nil {
		s.g = NewAnyValue(NewAnyValue(s.got).Elem())
	}
	return s.g
}
func (s *test) Want() AnyValue {
	if s.w == nil {
		s.w = NewAnyValue(NewAnyValue(s.got).Elem())
	}
	return s.w
}

func tRunTest(t *testing.T, tt *test) {
	if NewAnyValue(tt.got).IsComparable() && NewAnyValue(tt.want).IsComparable() {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want != tt.wantErr {
				if reflect.DeepEqual(tt.got, tt.want) == tt.wantErr {
					TError(t, tt.name, tt.got, tt.want)
				}
			}
		})
	}
}

type nilTest struct{}

func (nilTest) Run() {}
