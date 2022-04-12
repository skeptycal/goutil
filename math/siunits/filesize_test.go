package siunits

import (
	"testing"
)

func BenchmarkAllIntLen(b *testing.B) {
	for _, fn := range intLenFuncList {
		for i := 0; i < b.N; i++ {
			b.Run(fn.name, func(b *testing.B) {
				retval = fn.fn(i)
				retval = "nothing"
				retval = fn.fn(i)
			})
		}
	}
}

func TestAllIntLen(t *testing.T) {
	for _, tt := range IntLenTests() {
		for _, fn := range intLenFuncList {
			name := fn.name + "(" + tt.name + ")"
			t.Run(name, func(t *testing.T) {
				if got := fn.fn(tt.n); got != tt.want {
					t.Errorf("%v = %v, want %v", name, got, tt.want)
				}
			})
		}
	}
}
