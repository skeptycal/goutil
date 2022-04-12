package types

import (
	"reflect"
	"testing"
)

type ass = func() bool

// Assert is a general assertion caller
// that accepts a func() bool {} and a
// message to log if false
func (s *test) Assert(assertion func() bool, msg string) bool {
	if !assertion() {
		s.t.Logf("Assert %v == %v - %v (%T)", s.Got(), s.Want(), msg, assertion)
		s.t.FailNow()
		return false
	}
	return true
}

// AssertNot is a general assertion caller
// that accepts a func() bool {} and a
// message to log if true
func (s *test) AssertNot(assertion func() bool, msg string) bool {
	if assertion() {
		s.t.Logf("Assert %v != %v - %v (%T)", s.Got(), s.Want(), msg, assertion)
		s.t.FailNow()
		return false
	}
	return true
}

/// exported Methods are simple assertions with no test package functionality.

func (s *test) AssertEqual() bool {
	if !s.AssertComparable() {
		return false
	}

	return s.GV() == s.WV() || reflect.DeepEqual(s.GV(), s.WV())
}

func (s *test) AssertComparable() bool {
	return s.Got().IsComparable() && s.Want().IsComparable()
}

func (s *test) AssertOrdered() bool {
	return s.Got().IsOrdered() && s.Want().IsOrdered()
}

func (s *test) AssertIsIterable() bool {
	return s.Got().IsIterable() && s.Want().IsIterable()
}

func (s *test) AssertSameKind() bool {
	return s.GV().Kind() == s.WV().Kind()
}

func (s *test) AssertSameType() bool {
	return s.GV().Type() == s.WV().Type()
}

/// unexported Methods are assertions with test logging functionality.

func (s *test) assertEqual(t *testing.T) {
	if v := s.AssertComparable(); !v {
		s.assertComparable(t)
		return
	}
	if v := s.AssertEqual(); !v {
		s.LogType("values are not equal", s.g, s.w)
	}
}

func (s *test) assertComparable(t *testing.T) {
	// s.Assert(func() bool { return s.Got().IsComparable() && s.Want().IsComparable() }, "values not comparable")

	if !s.AssertComparable() {
		s.LogType("values are not comparable", s.g, s.w)
		s.t.FailNow()
	}
}
