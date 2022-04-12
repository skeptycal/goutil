package main

import (
	"reflect"
	"testing"
)

var (
	retval interface{}

	digitTests = []struct {
		name string
		b    []byte
		want []byte
	}{
		{"''", []byte(""), []byte("")},
		{"empty", []byte{}, []byte{}},
		{"<nil>", nil, nil},
		{"1234", []byte("1234"), []byte("1234")},
		{"12 34", []byte("12 34"), []byte("12")},
		{"1 234", []byte("1 234"), []byte("1")},
		{"1234 1234", []byte("1234 1234"), []byte("1234")},
		{"stuff1234", []byte("stuff1234"), []byte("1234")},
		{"stuff 1234", []byte("stuff 1234"), []byte("1234")},
		{"df483asoe1234", []byte("df483asoe1234"), []byte("483")},
		{"1234", []byte("1234"), []byte("1234")},
	}

	filterTests = []struct {
		name  string
		ff    filterfunc
		input []byte
		want  []byte
	}{
		{"digit", filterDigit, []byte("1234"), []byte("1234")},
		{"unicode digit", filterUnicodeDigit, []byte("1234"), []byte("1234")},
		{"digit", filterDigit, []byte("1234"), []byte("1234")},
		{"unicode digit", filterUnicodeDigit, []byte("1234"), []byte("1234")},
	}
	filterlist = []filterfunc{
		filterDigit,
		filterUnicodeDigit,
		filterNoSpace,
		filterUnicodeNumeric,
	}
)

func BenchmarkFilter(b *testing.B) {
	for _, bb := range filterTests {
		b.Run(bb.name, func(b *testing.B) {
			retval = SubSequence(bb.input, bb.ff)
			for i := 0; i < b.N; i++ {
				retval = SubSequence(bb.input, bb.ff)
			}
		})
	}
}

func TestFirstNumber(t *testing.T) {

	for _, tt := range digitTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubSequence(tt.b, filterDigit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
