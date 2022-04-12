package main

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

var c = Config{testUseGoRoutines: true}

type Any = interface{}

type Test interface {
	Run(t *testing.T)
	Name() string
	Input() string
	Got() string
	Want() string
	WantErr() bool
}

type TestSet interface {
	Run(t *testing.T)
	Name()
	t_UseGoRoutines() bool
}

type testBasicSet struct {
	list              []struct{} // filler .. must be replaced in specific Sets
	name              string
	testUseGoRoutines bool `default="true"`
}

func (ts testBasicSet) SetUseGoRoutines(v bool) { ts.testUseGoRoutines = v }
func (ts testBasicSet) t_UseGoRoutines() bool   { return ts.testUseGoRoutines }
func (ts testBasicSet) Name() string            { return ts.name }

type testSet struct {
	list []Test
	testBasicSet
}

func (ts testSet) Run(t *testing.T) {
	if ts.testUseGoRoutines {
		for _, tt := range ts.list {
			go tt.Run(t)
		}
	} else {
		for _, tt := range ts.list {
			tt.Run(t)
		}
	}
}

type testSetMap struct {
	list map[int]Test
	testBasicSet
}

func (ts testSetMap) KeysSorted() bool {
	last := math.MinInt64
	for k, _ := range ts.list {
		if k > last {
			return false
		}
	}
	return true
}

func (ts testSetMap) intKeysSorted() bool {
	return sort.IntsAreSorted(ts.Keys())
}

func (ts testSetMap) Keys() []int {
	keys := make([]int, len(ts.list))
	var i int = 0
	for k := range ts.list {
		keys[i] = k
		i++
	}
	return keys
}

func (ts testSetMap) Run(t *testing.T) {
	if ts.testUseGoRoutines {
		for _, tt := range ts.list {
			go tt.Run(t)
		}
	} else {
		for _, tt := range ts.list {
			tt.Run(t)
		}
	}
}

type testString struct {
	name    string
	input   Any
	got     Any
	want    Any
	wantErr bool
}

func (tst testString) Run(t *testing.T) { tRunAssertEqualWithError(t, tst) }
func (tst testString) Name() string     { return tst.name }
func (tst testString) WantErr() bool    { return tst.wantErr }
func (tst testString) Input() string    { return fmt.Sprintf("%v", tst.input) }
func (tst testString) Got() string      { return fmt.Sprintf("%v", tst.got) }
func (tst testString) Want() string     { return fmt.Sprintf("%v", tst.want) }

func TestConfig(t *testing.T) {
	ts := testSet{
		list: []Test{
			// only works on my machine!!
			testString{"gopath", "whoami", c.Sh("whoami"), "skeptycal", false},
			testString{"gopath", "echo $GOPATH", `echo ` + c.Env("$GOPATH"), "/Users/skeptycal/go", false},
		},
	}

	ts.Run(t)
}

func tRunAssertEqualWithError(t *testing.T, tt Test) {
	t.Run(tt.Name(), func(t *testing.T) {
		if tt.Got() != tt.Want() != tt.WantErr() {
			t.Errorf("%s = %v, want %v", tt.Name(), tt.Got(), tt.Want())
		}
	})
}
func NewTestSet(t *testing.T, name string, t_UseGoRoutines bool, tests ...Test) TestSet {
	ts := testSet{
		testUseGoRoutines: t_UseGoRoutines,
		list:              tests,
	}

	return &ts
}

func TestEnv(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		lenCheck int
		want     string
	}{
		{"gopath", `$GOPATH`, 7, `/Users/`},
		{"gopath", c.GoPath(), 7, `/Users/`},
		{"gopath", `$PATH`, 10, `/Users/ske`},
		{"gopath", c.Path(), 10, `/Users/ske`},
	}

}
