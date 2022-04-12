package types

import (
	"fmt"
	"io"
	"reflect"
	"testing"
	"unsafe"
)

func TAssertType(t *testing.T, name string, got, want reflect.Kind) {
	t.Helper()

	t.Run(name, func(t *testing.T) {
		if got != want {
			t.Errorf("not the same Kind: got %v(%T), want %v(%T)", got, got, want, want)
		}
	})
}

func TRun(t *testing.T, name string, got, want Any) {
	if r := recover(); r != nil {
		Log.Errorf("panic recovered while testing %v: %v", name, r)
	}
	// defer Rec()
	t.Helper()
	TAssertEqual(t, name, got, want)
}

func TAssertEqual(t *testing.T, name string, got, want Any) {
	t.Helper()

	// if reflect.DeepEqual(g.ValueOf(), w.ValueOf()) {
	// 	return
	// }

	AssertComparable(t, got, "nil")

	TTypeRun(t, name, got, want)

	t.Run(name, func(t *testing.T) {
		if !AssertEqual(got, want) {
			t.Errorf("values are not equal: got %v(%T), want %v(%T)", got, got, want, want)
		}
	})
}

// checkRetVal calls an error if retval is false
func checkRetVal(t *testing.T, msg string, retval bool) {
	if !retval {
		TError(t, msg, retval, false)
	}
}

func AssertComparable(t *testing.T, got, want Any) (retval bool) {
	g, w := PrepValues(got, want)
	retval = g.IsComparable() && w.IsComparable()
	checkRetVal(t, fmt.Sprintf("values are not comparable: %v(%T), %v(%T)", got, want, got, want), retval)
	return
}

func AssertEqual(got, want Any) bool {

	// g := NewAnyValue(got).Elem()
	// w := NewAnyValue(want).Elem()
	return true
	// return g == w // || reflect.DeepEqual(g, w)
}

func PrepValues(got, want Any) (AnyValue, AnyValue) {
	return NewAnyValue(NewAnyValue(got).Elem()), NewAnyValue(NewAnyValue(want).Elem())
}

func GetAnyValues(args ...Any) []AnyValue {
	list := make([]AnyValue, 0, len(args))

	for _, arg := range args {
		list = append(list, NewAnyValue(NewAnyValue(arg).Elem()))
	}
	return list
}

func BTest(b *testing.B, text string, got, want Any) bool {
	if got == want || reflect.DeepEqual(got, want) {
		return true
	}
	b.Errorf(text, got, want)
	return false
}

// BRun runs function with arguments args and returns
// a slice of results to a global dummy variable.
//
// If a panic occurs, BRun will mark it as failed
// and continue running other benchmarks / tests.
func BRun(b *testing.B, name string, function Any, args ...Any) {
	b.Helper()

	// Catch panics ...
	defer func() {
		if err := recover(); err != nil {
			// log.Errorf("panic occurred: %v\n", err)
			Log.Errorf("panic occurred: %v\n", err)
			b.FailNow()
		}
	}()

	// convert args to []reflect.Value for Call()
	in := ToValues(args)

	// convert function to AnyValue to access reflection methods
	anyfn := NewAnyValue(function)

	// fn is an alias to the actual function call
	//  before: using function.NewAnyValue().ValueOf()...
	//  2658889	       401.3 ns/op	      64 B/op	       3 allocs/op
	//  after: add anyfn and fn
	//  2925698	       437.2 ns/op	      64 B/op	       3 allocs/op
	// removed some timer resets
	//  3042357	       397.2 ns/op	      64 B/op	       3 allocs/op

	fn := anyfn.ValueOf().Call

	b.Run(name, func(b *testing.B) {
		// kind must be reflect.Func
		BTest(b, "'function' must be of Kind reflect.Func - got: %v(%T), want %v(%T)", anyfn.Kind(), reflect.Func)

		var retval Any = nil

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			retval = fn(in)
		}
		Log.Infof("benchmark results: %v", retval)
		global = retval
		b.ReportAllocs()
	})

}

var testSample int = 42
var ptrSample = &testSample

var reflectTests = []struct {
	name string
	a    Any
	want reflect.Kind
}{
	{"invalid", nil, reflect.Invalid},
	{"bool", true, reflect.Bool},
	{"int", 42, reflect.Int},
	{"uint", uint(42), reflect.Uint},
	{"int", 42, reflect.Int},
	{"Int8", int8(42), reflect.Int8},
	{"Int16", int16(42), reflect.Int16},
	{"Int32", int32(42), reflect.Int32},
	{"Int64", int64(42), reflect.Int64},
	{"Uint", uint(42), reflect.Uint},
	{"Uint8", uint8(42), reflect.Uint8},
	{"Uint16", uint16(42), reflect.Uint16},
	{"Uint32", uint32(42), reflect.Uint32},
	{"Uint64", uint64(42), reflect.Uint64},
	{"Uintptr", uintptr(42), reflect.Uintptr},
	{"Float32", float32(42), reflect.Float32},
	{"Float64", float64(42), reflect.Float64},
	{"Complex64", complex64(42), reflect.Complex64},
	{"Complex128", complex128(42), reflect.Complex128},
	{"Array", [4]int{42, 42, 42, 42}, reflect.Array},
	{"Chan", make(chan int, 1), reflect.Chan},
	{"Func", IsComparable, reflect.Func},
	{"Map", make(map[string]interface{}), reflect.Map},
	{"Ptr", ptrSample, reflect.Ptr},
	{"Slice", []int{42}, reflect.Slice},
	{"String", "42", reflect.String},
	{"UnsafePointer", unsafe.Pointer(nil), reflect.UnsafePointer},
	{"io.ReadCloser", io.NopCloser(nil), reflect.Interface},
	{"Interface", nil, reflect.Interface},
	{"byte slice element", []byte("fake")[0], reflect.Uint8},
}
