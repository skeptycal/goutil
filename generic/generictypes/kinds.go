package generic

import (
	"reflect"
)

// kindMaps is a map of the different kinds of types in Go and
// some useful information about each kind.
//
// It is often useful to know whether types are comparable
// and/or ordered at runtime when different types may be
// handled by the same code block.
//
// Since Go does not support
//
// - Basic data types are always comparable using the == and !=
// operators: integer values, floating-point numbers, complex
// numbers, boolean values, string values, constant values.
//
// - Array values are comparable, if they contain a comparable
// element type
//
// - Pointer values are comparable.
//
// - Channel values are comparable.
//
// - Interface values are comparable. Comparing interface values
// works only if the dynamic type is comparable.
//
// - Function values, Slice values and Map values are not
// comparable, they can only be compared with nil, as a special case.
//
// Some Caveats
//
// it is possible for a value to be unequal to itself:
// - because it is of func type
// - because it is a floating-point NaN value
// - because it is an array, struct, or interface containing
// func or NaN
//
// Pointer values are always equal to themselves, even if
// they point at or contain such problematic values, because
// they compare equal using Go’s == operator, and that is a
// sufficient condition to be deeply equal, regardless of content.
//
// Comparable types
//
// - Boolean values are comparable. Two boolean values are equal
// if they are either both true or both false.
// - Complex values are comparable. Two complex values u and v are
// equal if both real(u) == real(v) and imag(u) == imag(v).
// - Channel values are comparable. Two channel values are
// equal if they were created by the same call to make
// or if both have value nil.
// - Pointer values are comparable. Two pointer values are equal
// if they point to the same variable or if both have value nil.
// Pointers to distinct zero-size variables may or may not be equal.
// - Interface values are comparable. Two interface values are equal
// if they have identical dynamic types and equal dynamic values or
// if both have value nil.
// - Struct values are comparable if all their fields are comparable.
// Two struct values are equal if their corresponding non-blank
// fields are equal.
//
// Deeply Equal Comparisons
//
// - Pointer values are deeply equal if they are equal using
// Go’s == operator or if they point to deeply equal values.
// - Interface values are deeply equal if they hold
// deeply equal concrete values.
// - Struct values are deeply equal if their corresponding
// fields are deeply equal.
//
// Ordered types
//
// - Integer values are comparable and ordered, in the usual way.
// - Floating point values are comparable and ordered, as defined
// by the IEEE-754 standard.
// - String values are comparable and ordered, lexically byte-wise.
//
// Non-Comparable
//
// Function values, Slice values and Map values are not comparable;
// they can only be compared with nil, as a special case.
//
// Functions
//
// Functions are only deeply equal if they are both nil
//
// Opinion: these comparisons are not very useful in most cases.
//   The most useful comparison I have thought of is simply what
//   interface a function may implement.
//
// Maps
//
// Map values are deeply equal when all of the following are true:
//
// - they are both nil or both non-nil
//
// - they have the same length
//
// - they are the same map object or their corresponding keys
//   (matched using Go equality) map to deeply equal values.
//
// Opinion: Map comparisons may be more useful if they are
//   restricted to length and/or element type only.
//   Deep equality comparisons are costly and will rarely be useful.
//   How often do we have 2 copies of the *EXACT* same map and
//   need to compare it?
//
// Slices
//
// Slice values are deeply equal when all of the following are true:
// - they are both nil or both non-nil
// - they have the same length,
// - they point to the same initial entry of the same underlying
//   array (that is, &x[0] == &y[0]) or their corresponding
//   elements (up to length) are deeply equal.
//
// Opinion: similar to maps, slice comparisons may be more useful
//   if they are restricted to length and/or element type.
//
// Special Cases
//
// Arrays
//
// - Arrays require analysis of elements
//
// - Arrays are comparable if element type is comparable
//
// - Array values are deeply equal when their corresponding
// elements are deeply equal.
var kindMaps = kindMap{
	// comparable, ordered, deeplyComparable, alternate, iterable
	reflect.Invalid:       {false, false, false, true, false}, // 0
	reflect.Bool:          {true, false, true, false, false},  // 1
	reflect.Complex64:     {true, false, true, false, false},  // 15
	reflect.Complex128:    {true, false, true, false, false},  // 16
	reflect.Chan:          {true, false, true, false, false},  // 18
	reflect.Ptr:           {true, false, true, false, false},  // 22
	reflect.Interface:     {true, false, true, false, false},  // 20
	reflect.Struct:        {true, false, true, true, true},    // 25
	reflect.Int:           {true, true, true, true, true},     // 2
	reflect.Int8:          {true, true, true, true, true},     // 3
	reflect.Int16:         {true, true, true, true, true},     // 4
	reflect.Int32:         {true, true, true, true, true},     // 5
	reflect.Int64:         {true, true, true, true, true},     // 6
	reflect.Uint:          {true, true, true, true, true},     // 7
	reflect.Uint8:         {true, true, true, true, true},     // 8
	reflect.Uint16:        {true, true, true, true, true},     // 9
	reflect.Uint32:        {true, true, true, true, true},     // 10
	reflect.Uint64:        {true, true, true, true, true},     // 11
	reflect.Uintptr:       {true, true, true, true, true},     // 12
	reflect.Float32:       {true, true, true, true, true},     // 13
	reflect.Float64:       {true, true, true, true, true},     // 14
	reflect.String:        {true, true, true, true, true},     // 24
	reflect.Func:          {false, false, false, true, false}, // 19
	reflect.Map:           {false, false, false, true, true},  // 21
	reflect.Slice:         {false, false, false, true, true},  // 23
	reflect.Array:         {false, false, false, true, true},  // 17
	reflect.UnsafePointer: {false, false, false, true, false}, // 26
}

type kindMap = map[reflect.Kind]KindInfo

type // kindInfo stores information about each kind of type.
KindInfo struct {
	Comparable       bool // if it is useful ...
	Ordered          bool // some alternatives provided here
	DeeplyComparable bool // if it is useful ...
	Alternate        bool // some alternatives provided here
	Iterable         bool // some alternatives provided here
}
