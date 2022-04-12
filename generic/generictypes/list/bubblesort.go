package list

import (
	"fmt"
	"math/rand"
	"unsafe"

	"github.com/skeptycal/goutil/repo/errorlogger"
)

var log = errorlogger.New()

var defaultRandomNumberRange int = 100

// MakeRandomList generates a list of random numbers
// of type N, constrained as Number instead of Ordered
// because we cannot easily create an analogous
// process with strings.
func MakeRandomList[N Number](list []N, n int) List[N] {
	// TODO do something different for strings
	// if t := reflect.TypeOf(list.list[0]); t == stringType {
	// 	return MakeRandomString(list, n)
	// }
	// var i int8 = -1 // -1 binary representation: 11111111
	// var k uint8 = *(*uint8)(unsafe.Pointer(&i))
	// println(k) // 255 is the uint8 value for the binary 11111111

	var imin int = 12
	var imax int = 42

	if imin == imax {
		log.Errorf("min and max cannot be equal: %v and %v", imin, imax)
		imin = 0
		imax = defaultRandomNumberRange
	}
	if imin > imax {
		imin, imax = imax, imin
	}

	var min N = ToNumber[int, N](imin) // *(*N)(unsafe.Pointer(&imin))
	var max N = ToNumber[int, N](imax) // *(*N)(unsafe.Pointer(&imax))
	var rng N = max - min

	_ = rng

	slc := make([]N, n)

	// for i := 0; i < n; i++ {
	// 	slc[i] = *new(N)
	// }

	SetRandomValues(slc, rng)
	return List[N]{slc}
}

func SetRandomValues[N Number](list []N, rng N) []N {
	buf := make([]N, len(list))

	size := unsafe.Sizeof(rng) // size of a single instantiated N
	// intRange := ToNumber[N, int](rng)
	tinybuf := make([]byte, size)

	for i := 0; i < len(list); i++ {
		_, err := rand.Read(tinybuf)
		if err != nil {
			log.Infof("rand.Read(tinybuf) error: %v", err)
		}
		// n := rand.Intn(rng)

		var b byte
		for j := 0; j < int(size); j++ {
			b &= tinybuf[j] << (j * 8)
		}
		// x = ToNumber[byte, N](tinybuf[0 : size-1])
		buf[i] = ToNumber[byte, N](b)
		// tinybuf = []byte{}
	}

	return buf
}

// BubbleSort sorts a slice of Ordered items in place.
func BubbleSort[O Ordered](input []O) []O {
	n := len(input)
	swapped := true
	for swapped {
		// set swapped to false
		swapped = false
		// iterate through all of the elements in our list
		for i := 0; i < n-1; i++ {
			// if the current element is greater than the next
			// element, swap them
			if input[i] > input[i+1] {
				// log that we are swapping values for posterity
				fmt.Println("Swapping")
				// swap values using Go's tuple assignment
				input[i], input[i+1] = input[i+1], input[i]
				// set swapped to true - this is important
				// if the loop ends and swapped is still equal
				// to false, our algorithm will assume the list is
				// fully sorted.
				swapped = true
			}
		}
	}
	return input
}
