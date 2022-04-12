package siunits

import (
	"crypto/rand"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unsafe"

	"log"
)

const (
	MaxInt = math.MaxInt
)

var (
	MaxIntLen  = intLenString(MaxInt)
	MaxIntOnes = oneString(MaxIntLen)
)

// oneString creates a string of 1's with length n.
// If this would lead to an int that causes an
// overflow, the length is reduced until it is safe.
func oneString(n int) string {
	if n > MaxIntLen {
		n = MaxIntLen
	}
	return strings.Repeat("1", n)
}

// randString creates a string of random integers
// of length n.
// If this would lead to an int that causes an
// overflow, the length is reduced until it is safe.
func randString(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return string(b)
}

func IntLen(n int) int {
	return intLenString(n)
}

func intLenUnsafe(n int) int {
	return len(*(*string)(unsafe.Pointer(&n)))
}

func intLenLog(n int) int {
	// int n = 1000;
	return (int(math.Log10(float64(n)) + 1))
}

func intLenString(n int) int {
	return len(strconv.Itoa(n))
}

func intLenSprint(n int) int {
	s := fmt.Sprintf("%d", n)
	return len(s)
}

func intLenLoop(n int) int {

	//* 11111 11111 11111 11111

	r := 10
	i := 1
	for r <= n {
		i += 1
		r *= 10
	}
	return i
}

type maxIntListType map[int]int

var maxIntList maxIntListType = makeIntList(MaxIntLen)

func makeIntList(n int) maxIntListType {
	if n > MaxIntLen {
		n = MaxIntLen
	}

	m := make(map[int]int, n)
	for i := 0; i < MaxIntLen; i++ {
		s := "1" + strings.Repeat("0", MaxIntLen-1)
		r, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatalf("error parsing integer: %s", s)
		}
		m[i] = int(r)
	}
	return m
}

// var maxIntList maxIntListType = maxIntListType{
// 	1:  1,
// 	2:  10,
// 	3:  100,
// 	4:  1000,
// 	5:  10000,
// 	6:  100000,
// 	7:  1000000,
// 	8:  10000000,
// 	9:  100000000,
// 	10: 1000000000,
// 	11: 10000000000,
// 	12: 100000000000,
// 	13: 1000000000000,
// 	14: 10000000000000,
// 	15: 100000000000000,
// 	16: 1000000000000000,
// 	17: 10000000000000000,
// 	18: 100000000000000000,
// 	19: 1000000000000000000,
// }

func intLenCase(n int) int {

	if n < 10 {
		return 1
	}

	for i := MaxIntLen; i > 0; i-- {
		if n >= maxIntList[i] {
			return i
		}
	}
	return -1
}

func intLenSimple(n int) int {
	if n < 100000 {
		// 5 or less
		if n < 100 {
			// 1 or 2
			if n < 10 {
				return 1
			} else {
				return 2
			}
		} else {
			// 3 or 4 or 5
			if n < 1000 {
				return 3
			} else {
				// 4 or 5
				if n < 10000 {
					return 4
				} else {
					return 5
				}
			}
		}
	}
	// 6 or more
	if n < 10000000 {
		// 6 or 7
		if n < 1000000 {
			return 6
		} else {
			return 7
		}
	}

	// 8 to 10
	if n < 100000000 {
		return 8
	} else {
		// 9 or 10
		if n < 1000000000 {
			return 9
		} else {
			return 10
		}
	}

	// 11 to 12
	// if n < 1000000000000 {
	// 	if n < 100000000000 {
	// 		return 11
	// 	} else {
	// 		return 12
	// 	}
	// }
	// return -1
}

var (
	retval Any

	intLenFuncList = []struct {
		name string
		fn   func(n int) int
	}{
		// {"dict", intLenCase},
		// {"itoa", intLenString},
		// {"loop", intLenLoop},
		// {"unsafe", intLenUnsafe},
		// {"simple", intLenSimple},
		// {"log10", intLenLog},
		// {"Sprint", intLenSprint},
		{"Public", IntLen},
	}
)

type lenTestStruct struct {
	name string
	n    int
	want int
}

func pint(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return n
}

func IntLenTests() []lenTestStruct {

	m := []lenTestStruct{}

	for i := 0; i < MaxIntLen; i++ {
		one := oneString(i)
		test := lenTestStruct{
			name: "lenTest(" + one + ")",
			n:    pint(one),
			want: i,
		}
		m = append(m, test)
	}

	return m

	/*return []lenTestStruct{
		{"1", 1, 1},
		{"11", 11, 2},
		{"111", 111, 3},
		{"1111", 1111, 4},
		{"11111 ", 11111, 5},
		{"11111 1", 111111, 6},
		{"11111 11", 1111111, 7},
		{"11111 111", 11111111, 8},
		{"11111 1111", 111111111, 9},
		{"11111 11111", 1111111111, 10},
		{"11111 11111 1", 11111111111, 11},
		{"11111 11111 11", 111111111111, 12},
		{"11111 11111 111", 1111111111111, 13},
		{"11111 11111 11111", 11111111111111, 14},
		{"11111 11111 11111 ", 111111111111111, 15},
		{"11111 11111 11111 1", 1111111111111111, 16},
		{"11111 11111 11111 11", 11111111111111111, 17},
		{"11111 11111 11111 111", 111111111111111111, 18},
		{"11111 11111 11111 1111", 1111111111111111111, 19},
	}
	*/
}
