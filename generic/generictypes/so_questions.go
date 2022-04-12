package generic

import "fmt"

type (
	IExample[T any] interface {
		ExampleFunc(ex T) T
	}

	IExampleMap[T any] map[string]interface{ ExampleFunc(ex T) T }

	Example[T any] func(ex T) T

	ExampleMap[T any] map[string]func(ex T) T
)

func TryMapping1() {

	// var mmm1 map[string]IExample[any]

	// mapping maps string names to examples. The examples are instantiated with [any] and thus are filled in with interface objects that implement ExampleFunc(ex any) any
	var mapping = map[string]IExample[any]{
		"any1": anything{any: "stuff"},
		"any2": anything{any: "other stuff"},
	}

	// examples:
	// var mappingFloat map[string]IExample[float64]
	// var mapping3 map[string]IExample[any]

	mapping["any3"] = anything{any: "more different stuff"}

	// mapping["float64"] = float64thing{42.0} // compile time error: InvalidIfaceAssign

	// If the "float64" instance is passed, the following compile time error occurs:
	/* cannot use (float64thing literal) (value of type float64thing) as IExample[any] value in map literal: float64thing does not implement IExample[any] (wrong type for method ExampleFunc)
	// 		have ExampleFunc(ex float64) float64
	// 		want ExampleFunc(ex any) any m
	*/

	// "int": intthing{}, // similar compile time error: InvalidIfaceAssign
}

func PrintExample[T any](ex T) {
	fmt.Println(ex)
}

func MyFuncAny(ex any) any { return "stuff" }
func MyFuncInt(ex int) int { return 42 }

// Mapping2 maps string names to examples. The examples are stored as functions with the signature
// 		func(ex T) T
// which is instantiated to
// 		func(ex any) any
// in this example
var Mapping2 = ExampleMap[any]{
	"any":  MyFuncAny,
	"a":    anything{"stuff"}.ExampleFunc,
	"fake": func(ex any) any { return "fake" },
	// "int": myFuncInt, // compile time error: IncompatibleAssign

	// cannot use myFuncInt (value of type func(ex int) int)
	// as func(ex any) any value in map literal
}

// type IMap map[string]IExample

type (
	anything struct {
		any
	}
	intthing struct {
		int
	}
	float64thing struct {
		float64
	}
	stringthing struct{}
)

func (s stringthing) String() string {
	return "Not Implemented"
}

func (t anything) ExampleFunc(ex any) any             { return t.any }
func (t intthing) ExampleFunc(ex int) int             { return t.int }
func (t float64thing) ExampleFunc(ex float64) float64 { return t.float64 }

// func TryMap() {
// 	mapping := IExample[any]{
// 		"stuff": func(x any) any { return "stuff" },
// 	}
// }

// func (d *Dict[K, V]) IsEmpty() bool {
// 	return d.Len() == 0
// }
