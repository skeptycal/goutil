package types

import (
	"math"
	"testing"
	"unicode"
)

func Benchmark_test_bRun(b *testing.B) {
	fn := unicode.ToUpper
	BRun(b, "bRun test", fn, 'f')
}

func Test_Panic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			Log.Errorf("panic occurred: %v", err)
			// t.Errorf("panic occurred: %v\n", err)
			return
		} else {
			t.Errorf("panic did not occur as planned: %v\n", nil)
		}
	}()

	// var fakeList = make([]reflect.Value, 0, 5)
	// log.Printf("test bad slice index: %v", fakeList[6])

	panic("panic() executed")
}

func Test_Primes(t *testing.T) {

	answerList := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	p := primeNumbers(t, 100)

	// TRun(t, "primeNumbers(100)", p, []int{0, 1})
	TRun(t, "primeNumbers(100)", p, answerList)
	TRun(t, "primeNumbers(100)", p, nil)
}

func primeNumbers(t *testing.T, max int) []int {
	t.Helper()

	var primes []int

	for i := 2; i < max; i++ {
		isPrime := true

		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}

func Test_any_ValueOf(t *testing.T) {
	t.Parallel()
	for _, tt := range reflectTests {
		a := new_any(tt.a)
		A := NewAnyValue(tt.a)

		TRun(t, tt.name, a.ValueOf(), ValueOf(tt.a))

		TRun(t, tt.name, A.ValueOf(), ValueOf(tt.a))
	}
}

func Benchmark_any_IsComparable(b *testing.B) {
	for _, bb := range reflectTests {
		a := new_any(bb.a)
		AA := NewAnyValue(bb.a)

		b.Run("struct method", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				globalReturn = a.IsComparable()
			}
		})

		b.Run("interface method", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				globalReturn = AA.IsComparable()
			}
		})

		b.Run("global function", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				globalReturn = IsComparable(bb.a)
			}
		})

	}
}

func Benchmark_any_IsOrdered(b *testing.B) {
	for _, bb := range reflectTests {
		a := new_any(bb.a)
		AA := NewAnyValue(bb.a)

		b.Run("struct method", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				globalReturn = a.IsOrdered()
			}
		})

		b.Run("interface method", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				globalReturn = AA.IsOrdered()
			}
		})

		b.Run("global function", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				globalReturn = IsOrdered(bb.a)
			}
		})

	}
}
