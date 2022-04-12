package types

import "testing"

type (

	// testSet implements Tester and runs a set of tests
	testSet struct {
		name string
		t    *testing.T
		list []test
	}
)

func NewTestSet(t *testing.T, name string, list []test) TestRunner {
	return &testSet{
		t:    t,
		name: name,
		list: list,
	}
}

// Run runs all tests in the set.
func (ts *testSet) Run() {
	for _, tt := range ts.list {
		tt.Run()
	}
}
