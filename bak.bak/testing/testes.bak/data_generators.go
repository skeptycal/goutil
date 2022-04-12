package testes

import (
	"math"
	"math/cmplx"
	"math/rand"
	"reflect"
	"strings"
	"time"
	"unsafe"

	"github.com/skeptycal/goutil/testing/testes/kinds"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func RandomData(knd int, useInvalid bool) Any {
	var k reflect.Kind

	if 0 > knd || knd > 26 {
		k = kinds.RandomKind(useInvalid)
	} else {
		k = reflect.Kind(knd)
	}

	switch k {
	case 0: // Invalid
		if useInvalid {
			return reflect.Value{}
		}
		return RandomData(knd, false)
	case 1: // Bool
		return randBool()
	case 2: // Int
		return rand.Int()
	case 3: // Int8
		return int8(rand.Intn(b8))
	case 4: // Int16
		return int16(rand.Intn(b16))
	case 5: // Int32
		return int32(rand.Uint32())
	case 6: // Int64
		return int64(rand.Uint64())
	case 7: // Uint
		return uint(rand.Uint64())
	case 8: // Uint8
		return rand.Uint64() >> 56
	case 9: // Uint16
		return rand.Uint64() >> 48
	case 10: // Uint32
		return rand.Uint32()
	case 11: // Uint64
		return rand.Uint64()
	case 12: // Uintptr
	case 13: // Float32
		return rand.Float32()
	case 14: // Float64
		return rand.Float64()
	case 15: // Complex64
		return complex(rand.Float32(), rand.Float32())
	case 16: // Complex128
		return complex(rand.Float64(), rand.Float64())
	case 17: // Array
		a := [16]int{}
		for i := range a {
			a[i] = rng(10) + 5
		}
		return a
	case 18: // Chan
		return make(chan int, rng(10)+5)
	case 19: // Func
		return func() string { return "fake function" }
	case 20: // Interface
		return NewAnyValue(rng(42))
	case 21: // Map
		m := make(map[int]bool, rng(10)+5)
		for i := 0; i < len(m); i++ {
			m[i] = randBool()
		}
		return m
	case 22: // Ptr
		value := int(rand.Uint64())
		return &value
	case 23: // Slice
		s := make([]bool, 0, rng(10)+5)
		for i := 0; i < len(s); i++ {
			s[i] = randBool()
		}
		return s
	case 24: // String
		return RandomString(rng(10) + 5)
	case 25: // Struct
		return struct {
			name  string
			value int
		}{RandomString(rng(10) + 5), rng(10) + 5}
	case 26: // UnsafePointer
		value := int(rand.Uint64())
		return unsafe.Pointer(&value)
	}
	return nil
}

func makeRandomNumbers(numInts int, ch chan int) {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	for i := 0; i < numInts; i++ {
		ch <- generator.Intn(numInts * 100)
	}
}

func rng(n int) int { return rand.Intn(n) }

func Conj(c complex128) complex128 {
	return cmplx.Conj(c)
}

func randBool() bool {
	return rand.Int63() >= halfHalfBool
}

func noop(any interface{}) []AnyValue { return nil } // noop function

func stdev(accepted float64, list []float64) float64 {
	w := len(list)
	sum := 0.0

	ac2 := accepted * accepted
	for i := 0; i < w; i++ {

		diff := math.Abs(ac2 - math.Pow(list[i], 2))
		sum += math.Sqrt(diff)
	}

	return sum / float64(w)
}

func boolRatio(n int) float64 {
	m := boolSet(n)
	return float64(m[false]) / float64(m[true])
}

func boolSet(n int) map[bool]int {
	m := make(map[bool]int, 2)
	for i := 0; i < n; i++ {
		m[randBool()]++
	}
	return m
}

func boolSetGroup(n int) []float64 {
	list := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		list = append(list, boolRatio(n))
	}
	return list
}

func RandomString(n int) string {
	sb := strings.Builder{}
	defer sb.Reset()

	for i := 0; i < n; i++ {
		pos := rand.Intn(len(AllAlphanumeric) - 1)
		sb.WriteByte(AllAlphanumeric[pos])
	}

	return sb.String()
}
