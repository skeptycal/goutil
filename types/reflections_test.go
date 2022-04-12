package types_test

import (
	"io"
	"reflect"
	"testing"
	"unsafe"

	"github.com/fatih/structs"
	. "
)

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

func Test_ValueOf(t *testing.T) {
	for _, tt := range reflectTests {
		name := TName(tt.name, "ValueOf()", tt.a)
		TRun(t, name, NewAnyValue(tt.a).ValueOf(), ValueOf(tt.a))
	}
}

func Test_KindOf(t *testing.T) {
	for _, tt := range reflectTests {
		name := TName(tt.name, "KindOf()", tt.a)
		TRun(t, name, NewAnyValue(tt.a).Kind(), KindOf(tt.a))
	}
}

func Test_TypeOf(t *testing.T) {
	for _, tt := range reflectTests {
		name := TName(tt.name, "TypeOf()", tt.a)
		TRun(t, name, NewAnyValue(tt.a).TypeOf(), TypeOf(tt.a))
	}
}

func Test_Indirect(t *testing.T) {
	// LimitResult = true
	for _, tt := range reflectTests {
		want := ValueOf(tt.a)
		got := NewAnyValue(tt.a).Indirect().ValueOf()
		name := TName(tt.name, "Indirect", tt.a)
		TRun(t, name, got, want)
	}
	// LimitResult = false
}

func Test_Addr(t *testing.T) {
	for _, tt := range reflectTests {
		want := ValueOf(tt.a)
		got := Addr(want)

		if !want.CanAddr() {
			continue
		}

		name := TName(tt.name, "Addr()", tt.a)
		TRun(t, name, got, want.Addr())
	}
}

func Test_Interface(t *testing.T) {
	// defer leaktest.AfterTest(t)()
	for _, tt := range reflectTests {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("TestBench: recovered from panic while testing interface %v: %v", tt.name, r)
			}
		}()

		v := NewAnyValue(tt.a)
		if !v.ValueOf().IsValid() {
			continue
		}

		got := v.Interface()
		want := Interface(v.ValueOf())

		name := TName(tt.name, "Interface", tt.a)
		TRun(t, name, got, want)
	}
}

func Test_Elem(t *testing.T) {
	for _, tt := range reflectTests {
		want := ValueOf(tt.a)
		got := Elem(want)

		if want.Kind() != reflect.Ptr || want.Kind() != reflect.Interface {
			continue
		}
		name := TName(tt.name, "Elem()", tt.a)
		TRun(t, name, got, want.Elem())
	}
}

func Test_Convert(t *testing.T) {

	wantType := ValueOf(int(42)).Type()

	for _, tt := range reflectTests {
		v := ValueOf(tt.a)
		_ = Interface(Convert(v, wantType))

		if !v.IsValid() {
			continue
		}

		want := tt.a
		got := tt.a

		if v.CanConvert(wantType) {
			got = Convert(v, wantType).Interface()
			want = v.Convert(wantType).Interface()
		}

		name := TName(tt.name, "CanConvert()", tt.a)
		TRun(t, name, got, want)
	}
}

func TestNewStruct(t *testing.T) {
	tests := []struct {
		name string
		v    Any
		want *structs.Struct
	}{
		{"Cosa", Cosa{}, structs.New(Cosa{})},
		{"int 42", nil, NilStruct},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStruct(tt.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GuardReflectType(t *testing.T) {
	for _, tt := range reflectTests {
		t.Run(tt.name, func(t *testing.T) {
			// v := ValueOf(tt.a)
			want := TypeOf(tt.a)
			got := GuardReflectType(ValueOf(tt.a))
			if !reflect.DeepEqual(got, want) {
				t.Errorf("guardReflectType() = %v, want %v", got, want)
			}
		})
	}
}
