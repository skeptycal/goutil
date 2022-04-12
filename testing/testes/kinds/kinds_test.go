package kinds

import (
	"math/rand"
	"testing"
)

var tests = []struct {
	name string
	n    int
	want int
}{
	// TODO: Add test cases.
	{"1", 1, 1},
	{"1", 2, 4},
	{"1", 3, 9},
	{"1", 4, 16},
	{"1", 5, 25},
}

var global interface{}

func BenchmarkSquare(b *testing.B) {
	b.Run("square", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := rand.Intn(len(tests))
			bb := tests[n]
			global = square(bb.n)
		}
	})

	b.Run("square2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := rand.Intn(len(tests))
			bb := tests[n]
			global = square2(bb.n)
		}
	})

	b.Run("square3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := rand.Intn(len(tests))
			bb := tests[n]
			global = square3(bb.n)
		}
	})

	b.Run("square4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := rand.Intn(len(tests))
			bb := tests[n]
			global = square4(float64(bb.n))
		}
	})

	b.Run("square5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := rand.Intn(len(tests))
			bb := tests[n]
			global = square5(bb.n)
		}
	})
}

func Test_square(t *testing.T) {

	var intSquaredFuncs = []struct {
		name string
		fn   func(n int) int
	}{
		{"square", square},
		{"square2", square2},
		{"square3", square3},
		// {"square4", square4},
		// {"square5", square5},
		// {"square6", square6},
	}

	for _, ff := range intSquaredFuncs {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ff.fn(tt.n); got != tt.want {
					t.Errorf("%v() = %v, want %v", ff.name, got, tt.want)
				}
			})
		}
	}
}

func Test_square2(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		// TODO: Add test cases.
		{"1", 1, 1},
		{"1", 2, 4},
		{"1", 3, 9},
		{"1", 4, 16},
		{"1", 5, 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := square2(tt.n); got != tt.want {
				t.Errorf("square2() = %v, want %v", got, tt.want)
			}
		})
	}
}
