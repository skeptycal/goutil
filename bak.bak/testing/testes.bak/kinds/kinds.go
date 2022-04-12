package kinds

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"strings"
)

// Constants from Go standard library math/const.go
const (
	intSize = 32 << (^uint(0) >> 63) // 32 or 64
	MaxInt  = 1<<(intSize-1) - 1
	MinInt  = -1 << (intSize - 1)
)

type (
	KindMap map[string]int
)

func (m KindMap) Keys() []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (m KindMap) Values() []int {
	values := make([]int, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	sort.Ints(values)
	return values
}

func (m KindMap) Min() int {
	min := MaxInt
	for _, v := range m {
		if v < min {
			min = v
		}
	}
	return min
}

func (m KindMap) Max() int {
	max := MinInt
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

func (m KindMap) Mean() float64 {
	var sum = 0
	for _, v := range m {
		sum += v
	}
	return float64(sum) / float64(len(m))
}

func square(n int) int { return n * n }

func square2(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += n
	}
	return sum
}

func square3(n int) int {
	return int(math.Pow(float64(n), 2))
}

func square4(n float64) float64 {
	return math.Pow(n, 2)
}

func square5(n int) int {
	return n << 1
}

// StDev returns the standard deviation of the
// integers in the map.
func (m KindMap) StDev() float64 {
	var sum, x, mean float64
	mean = m.Mean()

	for _, v := range m {
		x = float64(v)
		sum += (x - mean) * (x - mean)
		// sum += dev(x, mean)
	}
	return math.Sqrt(sum / float64(len(m)))
}

func (m KindMap) String() string {
	sb := strings.Builder{}
	defer sb.Reset()

	for k, v := range m {
		sb.WriteString(fmt.Sprintf("%v = %v\n", k, v))
	}
	return sb.String()
}

func getEncodedString1(n int) string {
	b := []byte("")
	for i := 0; i < n; i++ {
		b = append(b, byte(RandomKind(true))+65)
	}
	return string(b)
}

func getEncodedString2(n int) string {
	sb := strings.Builder{}
	defer sb.Reset()
	for i := 0; i < n; i++ {
		sb.WriteByte(byte(RandomKind(true) + 65))
	}
	return sb.String()
}

func GetEncodedString(n int) string {
	return getEncodedString1(n)
}

func RandomKind(useInvalid bool) reflect.Kind {
	if useInvalid {
		return reflect.Kind(rand.Intn(26))
	}
	return reflect.Kind(rand.Intn(25) + 1)
}
