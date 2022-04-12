package main

import (
	"fmt"
	"reflect"
	"unsafe"

	. "
)

func main() {

	var intExample = 42
	tests := []struct {
		name string
		a    Any
	}{
		// {"invalid", nil, nil}, // panic
		// {"Interface", nil},	  // panic

		{"bool", true},
		{"int", 42},
		{"uint", uint(42)},
		{"int", 42},
		{"Int8", int8(42)},
		{"Int16", int16(42)},
		{"Int32", int32(42)},
		{"Int64", int64(42)},
		{"Uint", uint(42)},
		{"Uint8", uint8(42)},
		{"Uint16", uint16(42)},
		{"Uint32", uint32(42)},
		{"Uint64", uint64(42)},
		{"Uintptr", uintptr(42)},
		{"Float32", float32(42)},
		{"Float64", float64(42)},
		{"Complex64", complex64(42)},
		{"Complex128", complex128(42)},
		{"Array", [4]int{42, 42, 42, 42}},
		{"Chan", make(chan int, 1)},
		{"Func", IsComparable},
		{"Map", make(map[string]interface{})},
		{"Ptr", &intExample},
		{"Slice", []int{42}},
		{"String", "42"},
		{"Struct", ValueOf(42)},
		{"UnsafePointer", unsafe.Pointer(nil)},
	}

	var s []string
	for j, tt := range tests {
		t1 := reflect.ValueOf(tt.a)
		t := reflect.ValueOf(t1)

		v := t.MethodByName("Len")

		fmt.Printf("%2d: %-10v (%20v) (len: %v) methods(%v):\n", j, tt.name, t, v, t.NumMethod())

		s = []string{}
		for i := 0; i < t.NumMethod(); i++ {
			s = append(s, t.Type().Method(i).Name)
		}

		_ = s
	}
}
