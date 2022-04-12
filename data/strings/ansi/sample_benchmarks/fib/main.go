// Copyright (c) 2021 Michael Treanor
// https://
// MIT License

package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"

	"
)

func main() {
	const defaultCount = 7

	var n int = defaultCount

	if len(os.Args) > 1 {
		sc, err := strconv.Atoi(os.Args[1])
		if err == nil {
			n = sc
		}
	}

	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, ...

	var a *big.Int

	var s string
	var ss string

	fmt.Println()

	for i := 0; i < n; i++ {
		a = fib.BigFib(1 << i)
		s = fmt.Sprintf("%v", a)
		w := len(s)
		if w > 6 {
			ss = fmt.Sprintf("%c.%s E %d", s[0], s[1:5], w-1)
		} else {
			ss = fmt.Sprintf("%v", a)

		}
		fmt.Printf("%3d: Fib(%6d) = %20v\n", i, 1<<i, ss)
	}

	fmt.Println()

	for i := 0; i < n; i++ {
		a = fib.BigFib(i)
		fmt.Printf("%3d: Fib(%6d) = %20v\n", i, i, a)
	}

	fmt.Println()

	fib.TestMaxInt64()
}
