// Copyright (c) 2020 Michael Treanor
// MIT License

// Package fib contains sample benchmarks for Fibonacci numbers.
//
package fib

import (
	"fmt"
	"math"
	"math/big"
)

// Fib returns the nth number in the Fibonacci series
// using recursion.
//
// A good example of the use of benchmarks to increase efficiency
// and performance is found by benchmarking two versions of code
// that produce the Fibonacci numbers.
//
// As shown below, the common recursion method slows down very quickly
// and by the 40th iteration is nearly worthless.
//
// In contrast, a looping method using three variables to carry over
// numbers from each iteration maintains adequate efficiency over
// a much larger range of input numbers.
//
// For even larger numbers, the looping method is combined with the
// Go standard library 'big' package that can handle very large numbers.
// Although it loses efficiency faster then the looping method, it is
// able to calculate the 100,000th fibonacci number without much effort
// on a six year old laptop.
//
//  number of digits in Fib(100000): 20899
/*

fibFunc = Fib

	BenchmarkFib1-8          	462941701	         2.518 ns/op
	BenchmarkFib2-8          	163377795	         7.904 ns/op
	BenchmarkFib3-8          	94761864	         12.37 ns/op
	BenchmarkFib10-8         	 2723514	         445.4 ns/op
	BenchmarkFib20-8         	   20402	         55050 ns/op
	BenchmarkFib40-8         	       2	     779317892 ns/op

fibFunc = FibLoop

	BenchmarkFib1-8          	404492092	         2.793 ns/op
	BenchmarkFib2-8          	316388378	         3.713 ns/op
	BenchmarkFib3-8          	278436595	         4.307 ns/op
	BenchmarkFib10-8         	142531582	         8.438 ns/op
	BenchmarkFib20-8         	82595132	         14.58 ns/op
	BenchmarkFib40-8         	38278879	         31.35 ns/op
	BenchmarkFib10000-8   	  	  203696	          5667 ns/op

BigFib

	BenchmarkBigFib10-8   	      1847473	         596.2 ns/op
	BenchmarkBigFib100-8   	       349316	          3062 ns/op
	BenchmarkBigFib1000-8   	    24037	         43525 ns/op
	BenchmarkBigFib10000-8   	     1377	        808891 ns/op
	BenchmarkBigFib-8   	           21	      47725780 ns/op
*/
func Fib(n int) int64 {
	if n < 2 {
		return int64(n)
	}
	return Fib(n-1) + Fib(n-2)
}

// FibLoop returns the nth number in the Fibonacci series
// using a loop instead of recursion. Uses Int64.
func FibLoop(n int) int64 {
	if n < 2 {
		return int64(n)
	}

	var nMinus1 int64 = 0
	var nMinus2 int64 = 1
	var sum int64 = 0

	for i := 0; i < n; i++ {
		sum = nMinus1 + nMinus2
		nMinus2 = nMinus1
		nMinus1 = sum
	}
	return sum
}

// FibLoopU returns the nth number in the Fibonacci series
// using a loop instead of recursion. Uses Uint64.
func FibLoopU(n int) uint64 {
	if n < 2 {
		return uint64(n)
	}

	var nMinus1 uint64 = 0
	var nMinus2 uint64 = 1
	var sum uint64 = 0

	for i := 0; i < n; i++ {
		sum = nMinus1 + nMinus2
		nMinus2 = nMinus1
		nMinus1 = sum
	}
	return sum
}

// BigFib returns the nth number in the Fibonacci series
// using a loop instead of recursion. It uses the Go standard
// library 'big' package to handle large numbers.
func BigFib(n int) *big.Int {
	maxN := int64(n)
	if n < 2 {
		return big.NewInt(maxN)
	}

	var nMinus1 = big.NewInt(0)
	var nMinus2 = big.NewInt(1)
	var sum = big.NewInt(0)

	for i := 0; i < n; i++ {
		sum.Add(nMinus1, nMinus2)

		nMinus2.Set(nMinus1)
		nMinus1.Set(sum)

	}
	return sum
}

func TestMaxInt64() {
	var n int = 100000
	var max int64 = math.MaxInt64
	var r int64
	var r1 int64
	var c int

	for i := 0; i < n; i++ {
		c = i
		r = FibLoop(i)
		if r > max || r < 0 {
			break
		}
		r1 = r
	}
	fmt.Printf("Max Int64:\n %v\n", max)
	fmt.Printf("Maximum Fibonacci number as %T: Fib(%d) =\n %v\n", r1, c, r1)
}
