package main

import (
	"fmt"
)

// type Number interface {
// 	constraints.Unsigned | constraints.Signed | constraints.Float
// }

func newGenericFunc[age IntType | FloatType](myAge age) {
	fmt.Println(myAge)
}

func newGenericFuncAny[a any](myAge a) {
	fmt.Println(myAge)
}

func newGenericFuncAdd[N Number](n N) {
	val := ToInt(n) + 1
	fmt.Println(val)
}

func main() {
	fmt.Println("Go Generics Tutorial")
	var testAge int64 = 23
	var testAge2 float64 = 24.5

	newGenericFunc(testAge)
	newGenericFunc(testAge2)

	var testString string = "Elliot"

	newGenericFuncAny(testAge)
	newGenericFuncAny(testAge2)
	newGenericFuncAny(testString)

	newGenericFuncAdd(testAge)
	newGenericFuncAdd(testAge2)

}

type (
	Comparable interface {
		comparable
	}

	// All Ordered types
	// string
	// int, int16, int32, int64, int8
	// uint, uint16, uint32, uint64, uint8
	// float32, float64
	// uintptr
	AllOrdered interface {
		Number | ~string | uintptr
	}

	Ordered interface {
		Number | ~string
	}

	Stringable interface {
		IntType | UintType
	}

	Number interface {
		IntType | UintType | FloatType
	}

	IntType interface {
		int | int8 | int16 | int32 | int64
	}

	UintType interface {
		uint | uint8 | uint16 | uint32 | uint64
	}

	FloatType interface {
		float32 | float64
	}
)
