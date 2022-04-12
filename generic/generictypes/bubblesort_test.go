package generic

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

// N returns a random number between 0 and n
// orders of magnitude.
func N(oom int) int { return 1 << rand.Intn(4) }

func TestToNumber(t *testing.T) {
	var tint int = 42
	var tuint uint = 42
	var tfloat float64 = 42.0

	t.Errorf("actual type of 42(int): %v(%T)", tint, tint)
	t.Errorf("actual type of 42(uint): %v(%T)", tuint, tuint)
	t.Errorf("actual type of 42(float64): %v(%T)", tfloat, tfloat)

}

func TestToString(t *testing.T) {
	var tint = fmt.Sprintf("%v", int(42))
	var tint2 = fmt.Sprintf("%v", uint(42))
	var tint3 = fmt.Sprintf("%f", float64(42.0))

	t.Errorf("actual type of 42(int) converted to int: %v(%T)", tint, tint)
	t.Errorf("actual type of 42(uint) converted to int: %v(%T)", tint2, tint2)
	t.Errorf("actual type of 42(float64) converted to int: %v(%T)", tint3, tint3)
}

func TestToInt(t *testing.T) {
	var tint = ToInt(42)
	var tint2 = ToInt(uint(42))
	var tint3 = ToInt(float64(42))

	t.Errorf("actual type of 42(int) converted to int: %v(%T)", tint, tint)
	t.Errorf("actual type of 42(uint) converted to int: %v(%T)", tint2, tint2)
	t.Errorf("actual type of 42(float64) converted to int: %v(%T)", tint3, tint3)
}
func TestListerIsSorted(t *testing.T) {
	list := MakeRandomList(make([]int, 42), N(5))
	got := list.IsSorted()
	tRun(t, "Lister.IsSorted", got, false, false)
	t.Log()
}

type Tester interface {
	Run() error
}

type oneTest[N Number] struct {
	name    string
	input   []N
	want    []N
	wantErr bool
}

type testNumberSlice[N Number] []oneTest[N]

func tRun[T any](t *testing.T, name string, got, want T, wantErr bool) {
	t.Run(name, func(t *testing.T) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s = %v(%T), want %v(%T)", name, got, got, want, want)
		}
	})
}
