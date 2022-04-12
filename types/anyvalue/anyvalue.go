package anyvalue

import (
	"fmt"
	"reflect"
)

// Any represents a object that may contain any
// valid type.
type Any = interface{}

type any struct {
	v       reflect.Value
	k       reflect.Kind
	t       reflect.Type
	kindmap kindInfo
	i       interface{}
}

// NewAnyValue returns a new AnyValue interface, e.g.
// v := NewAnyValue(uint(42))
//
// AnyValue is a wrapper around the Any interface,
// or interface{}, which may contain any value.
// The original interface{} value is returned by
// v.Interface()
//
// The extra features of this wrapper allow
// value, type, and kind information, as well
// as whether the type is comparable, ordered,
// and/or iterable.
func NewAnyValue(a Any) AnyValue {
	if v, ok := a.(AnyValue); ok {
		return v
	}
	return new_any(a)
}

func new_any(a Any) *any {

	switch v := a.(type) {
	case reflect.Value:
		a = v
	case nil:
		break
	default:
		a = ValueOf(v).Interface()

	}

	return &any{
		v:       ValueOf(a),
		kindmap: NewKindInfo(a),
		i:       a,
	}
}

// Kind returns the type of the underlying variable.
// This is cached in a jit struct field upon request.
func (a *any) TypeOf() reflect.Type {

	if a.t == nil {
		a.t = TypeOf(a.i)
	}

	return a.t
}

// Kind returns the kind of type of the underlying variable.
// This is cached in a jit struct field upon request.
func (a *any) Kind() reflect.Kind {
	if a.k == 0 {
		a.k = reflect.ValueOf(a.i).Kind()
	}
	return a.k
}

// Indirect returns the value that a pointer
// points to. If the underlying object is a nil
// pointer, Indirect returns a zero Value.
// If the underlying is not a pointer, Indirect
// returns the orignal AnyValue.
func (a *any) Indirect() AnyValue {
	if a.Kind() == reflect.Ptr {
		return NewAnyValue(a)
	}
	return a
}

// Elem returns the value that the interface
// contains or that the pointer poin111111111111111111111111111111111111111111111111111111111111111111111w                                                                                                                                                                        wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwts to.
// If the kind of the AnyValue is not Interface
// or Ptr, the original AnyValue is returned.
// It returns the zero Value if the underlying is nil.
func (a *any) Elem() reflect.Value {
	return Elem(a.ValueOf())
	// v := a.ValueOf()

	// if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
	// 	return v.Elem()
	// }
	// return v
}

func (a *any) ValueOf() reflect.Value { return ValueOf(a.i) }
func (a *any) KindInfo() KindInfo     { return a.kindmap }
func (a *any) Interface() Any {
	if a.i == nil {
		v := a.ValueOf()
		if !v.IsValid() {
			return nil
		}

		if a.ValueOf().CanInterface() {
			a.i = Interface(v)
		}

		a.i = Interface(Elem(v))
	}
	return a.i
}
func (a *any) IsComparable() bool     { return a.KindInfo().IsComparable() }
func (a *any) IsOrdered() bool        { return a.KindInfo().IsOrdered() }
func (a *any) IsDeepComparable() bool { return a.KindInfo().IsDeepComparable() }
func (a *any) IsIterable() bool       { return a.KindInfo().IsIterable() }
func (a *any) HasAlternate() bool     { return a.KindInfo().HasAlternate() }
func (a *any) String() string         { return fmt.Sprintf("%v", a.i) }

// AnyValue is a wrapper around the Any interface,
// or interface{}, which may contain any value.
// The extra features of this wrapper allow
// value, type, and kind information, as well
// as whether the type is comparable, ordered,
// and/or iterable.
type AnyValue interface {

	// ValueOf returns a new Value initialized to the
	// concrete value stored in the interface i.
	// ValueOf(nil) returns the zero Value.
	ValueOf() reflect.Value

	// TypeOf returns the object's type.
	TypeOf() reflect.Type

	// Kind returns v's Kind. If v is the zero Value
	// (IsValid returns false), Kind returns Invalid.
	Kind() reflect.Kind

	// Interface returns the original underlying interface.
	Interface() Any

	// Indirect returns the value pointed to by a pointer.
	// If the AnyValue is not a pointer, indirect returns the
	// AnyValue unchanged.
	Indirect() AnyValue

	// Elem returns the value that the interface contains
	// or that the pointer points to. If the kind of the
	// AnyValue is not Interface or Ptr, the original
	// AnyValue is returned.
	Elem() reflect.Value

	String() string

	KindInfo
}
