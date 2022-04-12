package list

import (
	"fmt"
	"unsafe"

	constraints "github.com/skeptycal/goutil/generic/generictypes/constraints"
)

type (
	Number     constraints.Number
	Ordered    constraints.Ordered
	Stringable any
)

// TODO ...
func MakeRandomString[O Ordered](list List[O], n int) List[O] {
	return list
}

// ToNumber recasts a variable of any type in the
// constraint Number to a specific instantiated
// type of Number.
func ToNumber[N1, N2 Number](n N1) N2 {
	return *(*N2)(unsafe.Pointer(&n))
}

// ToInt recasts a variable of any type in the
// constraint Number to int.
func ToInt[N Number](n N) int {
	return *(*int)(unsafe.Pointer(&n))
}

func toInt2[N Number](n N) int {
	return ToNumber[N, int](n)
}

// func toString[S Stringable](n S) string {
// 	return string(n)
// }
func ToString[S Stringable](s []S) string {
	length := len(s)

	if length == 0 {
		return ""
	}

	return fmt.Sprintf("%v", s)

	/*

		size := int(unsafe.Sizeof(s[0]))

		alloc := length * size

		buf := make([]byte, alloc)

		for _, v := range s {
			c := []byte{}
			for j := 0; j < size; j++ {
				c = append(c, byte(v))
			}
			buf = append(buf, c...)
		}

		return string(buf)
		// return *(*string)(unsafe.Pointer(&n))
	*/
}

// ToByte recasts a variable of any type in the
// constraint Number to byte.
func ToByte[N Number](n N) byte {
	return *(*byte)(unsafe.Pointer(&n))
}
