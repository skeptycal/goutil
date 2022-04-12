package types

import (
	"reflect"

	"github.com/fatih/structs"
)

var NilStruct = &structs.Struct{}

// ValueOf returns a new Value initialized to the concrete value stored in the interface i. ValueOf(nil) returns the zero Value.
func ValueOf(i interface{}) reflect.Value {
	return reflect.ValueOf(i)
}

// Kind returns a's Kind. If a is the zero Value (IsValid returns false), Kind returns Invalid.
func KindOf(a Any) reflect.Kind {
	return ValueOf(a).Kind()
}

// Type returns a's type.
func TypeOf(a Any) reflect.Type {
	if a == nil {
		return nil
	}
	return ValueOf(a).Type()
}

// Indirect returns the value that v points to. If v is a nil pointer, Indirect returns a zero Value. If v is not a pointer, Indirect returns v.
func Indirect(v reflect.Value) reflect.Value {
	return reflect.Indirect(v)
}

// Addr returns a pointer value representing the address of v. If v is not addressable, v is returned unchanged.
//
// Addr is typically used to obtain a pointer to a struct field or slice element in order to call a method that requires a pointer receiver.
func Addr(v reflect.Value) reflect.Value {
	if !v.CanAddr() {
		return v
	}
	return v.Addr()
}

// Interface returns v's current value as an interface{}.
// It is equivalent to:
//	var i interface{} = (v's underlying value)
// It panics if the Value was obtained by accessing
// unexported struct fields.
func Interface(v reflect.Value) Any {
	if !v.IsValid() {
		return v
	}

	if v.CanInterface() {
		return v.Interface()
	}
	return v
}

// Elem returns the value that the interface
// contains or that the pointer points to.
// If the kind of a is not Interface or Ptr,
// the v is returned.
// It returns the zero Value if the underlying is nil.
func Elem(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Interface:
		return Elem(ValueOf(Interface(v)))
	case reflect.Ptr:
		return Indirect(v)
	case reflect.Invalid:
		return ValueOf(nil)
	default:
		return v
	}
}

// Convert returns the value v converted to type t. If the
// usual Go conversion rules do not allow conversion of the
// value v to type t, or if converting v to type t would
// panic, v is returned.
func Convert(v reflect.Value, typ reflect.Type) reflect.Value {
	if !v.IsValid() {
		return v
	}

	if v.CanConvert(typ) {
		return v.Convert(typ)
	}
	return v
}

func NewStruct(v Any) *structs.Struct {
	if KindOf(v) == reflect.Struct {
		return structs.New(v)
	}
	return NilStruct
}

func GuardReflectType(v reflect.Value) reflect.Type {
	if v.Kind() == reflect.Invalid {
		return nil
	}

	return v.Type()
}
