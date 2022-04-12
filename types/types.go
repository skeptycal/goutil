package types

import (
	"reflect"

	"
	"
)

var (
	Log = errorlogger.Log
)

type (
	Any = anyvalue.Any

	// Type is the representation of a Go type.
	//
	// Not all methods apply to all kinds of types. Restrictions,
	// if any, are noted in the documentation for each method.
	// Use the Kind method to find out the kind of type before
	// calling kind-specific methods. Calling a method
	// inappropriate to the kind of type causes a run-time panic.
	//
	// Type values are comparable, such as with the == operator,
	// so they can be used as map keys.
	// Two Type values are equal if they represent identical types.
	//
	// Synopsis of methods:
	//
	// 	type Type interface {
	// 		Align() int
	// 		FieldAlign() int
	// 		Method(int) Method
	// 		MethodByName(string) (Method, bool)
	// 		NumMethod() int
	// 		Name() string
	// 		PkgPath() string
	// 		Size() uintptr
	// 		String() string
	// 		Kind() Kind
	// 		Implements(u Type) bool
	// 		AssignableTo(u Type) bool
	// 		ConvertibleTo(u Type) bool
	// 		Comparable() bool
	// 		Bits() int
	// 		ChanDir() ChanDir
	// 		IsVariadic() bool
	// 		Elem() Type
	// 		Field(i int) StructField
	// 		FieldByIndex(index []int) StructField
	// 		FieldByName(name string) (StructField, bool)
	// 		FieldByNameFunc(match func(string) bool) (StructField, bool)
	// 		In(i int) Type
	// 		Key() Type
	// 		Len() int
	// 		NumField() int
	// 		NumIn() int
	// 		NumOut() int
	// 		Out(i int) Type
	//
	//	 	common() *rtype
	// 		uncommon() *uncommonType
	// 	}
	Type = reflect.Type

	// Cosa is a generic 'thing.' It is an empty stuct
	// used for 'preallocating' zero resource objects.
	Cosa struct{}
)
