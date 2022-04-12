package generic

import (
	"fmt"
	"testing"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1, false)
		AssertEqual(t, 1, 2, true)
		AssertNotEqual(t, 1, 2, false)
		AssertNotEqual(t, 1, 1, true)

	})
}
func AssertEqual[T comparable](t *testing.T, got, want T, wantErr bool) {
	t.Helper()
	if got != want != wantErr {
		t.Errorf("AssertEqual failed: got %v(%T), want %v(%T)", got, got, want, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T, wantErr bool) {
	t.Helper()
	if got == want != wantErr {
		t.Errorf("AssertNotEqual failed: got %v(%T), want %v(%T)", got, got, want, want)
	}
}
func ExampleListOrderedKinds() {
	fmt.Println("List of ordered Go reflect.Kind types.")
	fmt.Println(ListOrderedKinds())

	// TODO: list order varies ... so test will fail most of the time
	// TODO: (put an empty line below this one to reactivate this test)
	// Output:
	// List of ordered Go reflect.Kind types.
	// [int32 int16 uint8 uintptr float32 uint string float64 int int8 int64 uint16 uint32 uint64]
}

func ExampleExampleAddOne() {
	var i int64 = 42
	var f float32 = 42.0

	ExampleAddOne(i)
	ExampleAddOne(f)

	// Output:
	// 43
	// 43
}

func ExampleScale() {

	list := []int{0, 7, 6, 2, 8, 8, 8, 2, 8, 1, 5, 3, 2, 8, 9, 1, 2, 6, 7, 0}
	// list := SmallIntList(20)

	scaled := Scale(list, 2)

	fmt.Println(list)
	fmt.Println(scaled)

	// Output:
	// [0 7 6 2 8 8 8 2 8 1 5 3 2 8 9 1 2 6 7 0]
	// [0 14 12 4 16 16 16 4 16 2 10 6 4 16 18 2 4 12 14 0]
}

func TestScale(t *testing.T) {

	n := 20
	multiplier := 3
	list := SmallIntList(n)
	scaled := Scale(list, multiplier)

	t.Run("Scale(ints)", func(t *testing.T) {
		for i := 0; i < len(list); i++ {
			if list[i]*multiplier != scaled[i] {
				t.Errorf("Scaled item not equal to list item: got %v(%T), want %v(%T)", scaled[i], scaled[i], list[i]*multiplier, list[i]*multiplier)
			}
		}
	})

	// var gens = RandomNumberList(n)
	// genMultiplier := N(multiplier)
	// genscaled := Scale(gens, genMultiplier)

	// t.Run("Scale(generic Numbers)", func(t *testing.T) {
	// 	for i := 0; i < len(list); i++ {
	// 		if gens[i]*genMultiplier != genscaled[i] {
	// 			t.Errorf("Scaled item not equal to list item: got %v(%T), want %v(%T)", genscaled[i], genscaled[i], gens[i]*genMultiplier, gens[i]*genMultiplier)
	// 		}
	// 	}
	// })
}

// func GetOrderedPairs[T Ordered](t *testing.T) (a T, b T, want T) {
// 	t.Helper()
// tests := []struct {
// 	name string
// 	a    T
// 	b    T
// 	want T
// }{
// 	{"1,2", 1, 2, 2},
// 	{"2,1", 2, 1, 2},
// 	{"2,1", "2", "1", "2"},
// }
// a = int(1)
// b = int8(2)

// return a, b, a + b
// }

// func TestMax(t *testing.T) {
// 	a, b, want := GetOrderedPairs(t)

// 	// for _, test := range tests {
// 	// for _, typ := range types {
// 	t.Run(test.name, func(t *testing.T) {
// 		if got := Max(test.a, test.b); got != test.want {
// 			t.Errorf("Max(%v, %v) = %v, want %v", test.a, test.b, got, test.want)
// 		}
// 	})
// 	// }
// 	// }
// }
