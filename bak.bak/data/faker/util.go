package faker

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

const (
	maxArgChoices int     = 10
	one           uint64  = 1<<64 - 1
	halfBool      int64   = 1<<63 - 1
	halfHalfBool  int64   = 1<<62 - 1
	ratio         float64 = float64(one) / float64(halfBool)
	halfRatio     float64 = float64(halfBool) / float64(halfHalfBool)
)

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
