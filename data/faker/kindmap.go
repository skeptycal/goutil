package faker

import (
	"fmt"
	"math"
	"sort"
	"strings"
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
	min := 1<<63 - 1
	for _, v := range m {
		if v < min {
			min = v
		}
	}
	return min
}

func (m KindMap) Max() int {
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

func (m KindMap) Mean() float64 {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return float64(sum) / float64(len(m))
}

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
