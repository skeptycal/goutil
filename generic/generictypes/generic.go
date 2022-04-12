package generic

import (
	"math"
	"math/rand"
	"reflect"
	"time"

	"github.com/skeptycal/goutil/repo/errorlogger"
)

var log = errorlogger.New()

var stringType = reflect.TypeOf(" ")

func AddOne(ex int) int { return ex + 1 }

func ListOrderedKinds() []string {
	list := make([]string, 0)
	for k, v := range kindMaps {
		if v.Ordered {
			list = append(list, k.String())
		}
	}
	return list
}

func MinMax[T Ordered](a, b T) (T, T) {
	if a < b {
		return a, b
	}
	return b, a
}

func Max[T Ordered](a, b T) T {
	if a < b {
		return b
	}
	return a
}

func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func RandomNumber[T Number]() T {
	switch rand.Intn(2) {
	case 0:
		return T(rand.Int63())
	case 1:
		return T(rand.Uint64())
	case 2:
		return T(rand.Float64() * float64(rand.Int63()))
	}
	return T(math.NaN())
}

// func RandomNumberChan[T Number](n int) chan T {
// 	c := make(chan T)
// 	return c
// }

func RandomNumberList[T IntType](n int) []T {
	list := make([]T, n)
	for i := 0; i < n; i++ {
		switch rand.Intn(2) {
		case 0:
			list[i] = T(rand.Int63())
		case 1:
			list[i] = T(rand.Uint64())
		case 2:
			list[i] = T(rand.Float64() * float64(rand.Int63()))
		default:
			list[i] = T(math.NaN())
		}
	}
	return list
}

func SmallIntList(n int) []int {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = int(rand.Intn(9))
	}
	return list
}

// func SmallNumberList[N Number](n int) []N {
// 	list := make([]N, n)
// 	for i := 0; i < n; i++ {
// 		var v any = RandomNumber()
// 		list[i] = RandomNumber()
// 	}
// }

func Scale[E IntType](s []E, c E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

// An implementation of Interface can be sorted by the routines in this package.
// The methods refer to elements of the underlying collection by integer index.
type Sort interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	//
	// If both Less(i, j) and Less(j, i) are false,
	// then the elements at index i and j are considered equal.
	// Sort may place equal elements in any order in the final result,
	// while Stable preserves the original input order of equal elements.
	//
	// Less must describe a transitive ordering:
	//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
	//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
	//
	// Note that floating-point comparison (the < operator on float32 or float64 values)
	// is not a transitive ordering when not-a-number (NaN) values are involved.
	// See Float64Slice.Less for a correct implementation for floating-point values.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
