package fib

import (
	"fmt"
	"math/big"
	"testing"
)

var fibTests = []struct {
	n        int   // input
	expected int64 // expected result
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}

func TestFib(t *testing.T) {
	for _, tt := range fibTests {
		actual := Fib(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

func TestFibLoop(t *testing.T) {
	for _, tt := range fibTests {
		actual := FibLoop(tt.n)
		if actual != tt.expected {
			t.Errorf("FibLoop(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

func TestFibLoopU(t *testing.T) {
	for _, tt := range fibTests {
		actual := FibLoopU(tt.n)
		if actual != uint64(tt.expected) {
			t.Errorf("FibLoop(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

func TestBigFib(t *testing.T) {
	for _, tt := range fibTests {
		got := BigFib(tt.n)
		want := big.NewInt(tt.expected)
		if got.Cmp(want) != 0 {
			t.Errorf("FibLoop(%d): expected %d, actual %d", tt.n, want, got)
		}
	}
}

var fibFunc = FibLoop

// Comparison
/*
/fibFunc = Fib
BenchmarkFib1-8          	462941701	         2.518 ns/op
BenchmarkFib2-8          	163377795	         7.904 ns/op
BenchmarkFib3-8          	94761864	        12.37 ns/op
BenchmarkFib10-8         	 2723514	       445.4 ns/op
BenchmarkFib20-8         	   20402	     55050 ns/op
BenchmarkFib40-8         	       2	 779317892 ns/op
BenchmarkFibComplete-8   	 2874825	       429.8 ns/op

/fibFunc = FibLoop
BenchmarkFib1-8          	404492092	         2.793 ns/op
BenchmarkFib2-8          	316388378	         3.713 ns/op
BenchmarkFib3-8          	278436595	         4.307 ns/op
BenchmarkFib10-8         	142531582	         8.438 ns/op
BenchmarkFib20-8         	82595132	        14.58 ns/op
BenchmarkFib40-8         	38278879	        31.35 ns/op
BenchmarkFibComplete-8   	140743995	         8.431 ns/op
*/

var n = 20

func BenchmarkSeries(b *testing.B) {
	for c := 0; c < n; c++ {
		name := fmt.Sprintf("Fib(%d)", c)
		b.Run(name, func(b *testing.B) {
			var r int64
			for i := 0; i < b.N; i++ {
				r = fibFunc(1 << c)
			}
			result = r
		})
	}
}

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibFunc(i)
	}
}

func BenchmarkFib1(b *testing.B)     { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)     { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)     { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B)    { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B)    { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B)    { benchmarkFib(40, b) }
func BenchmarkFib200(b *testing.B)   { benchmarkFib(200, b) }
func BenchmarkFib10000(b *testing.B) { benchmarkFib(10000, b) }

var result int64

// func BenchmarkFibComplete(b *testing.B) {
func benchmarkFibComplete(b *testing.B) {
	var r int64
	for n := 0; n < b.N; n++ {
		// always record the result of fibFunc to prevent
		// the compiler eliminating the function call.
		r = fibFunc(10)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

var resultBI = big.NewInt(0)

func BenchmarkBigFib(b *testing.B) {

	const iter = 40
	var r *big.Int
	for n := 0; n < b.N; n++ {
		// always record the result of fibFunc to prevent
		// the compiler eliminating the function call.
		r = BigFib(iter)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	resultBI = r
	s := fmt.Sprintf("%v", resultBI)
	fmt.Printf("Fib(%d): %v\n", iter, s)
	fmt.Printf("number of digits in Fib(%d): %v\n", iter, len(s))
}

func BenchmarkUintCompare(b *testing.B) {
	// using Uint64 is marginally faster than Int64
	// ~0.1 to 1%
	b.Run("Int64", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = FibLoop(33)
		}
	})

	b.Run("Uint64", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = FibLoopU(33)
		}
	})
}
