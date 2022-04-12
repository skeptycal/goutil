package stats

import (
	"math"
	"time"
)

type Any = interface{}

type (
	StatMap interface {
		Keys() []string
		Values() []int
		Min() int
		Max() int
		Mean() float64
		StDev() float64

		String() string
	}

	dataPoint struct {
		t0   time.Time
		t1   time.Time
		fn   func() Any
		data Any
	}

	DataPoint interface {
		Start()
		Stop()
		Collect()
	}
)

func dev(x, m float64) float64 {
	return (x - m) * (x - m)
}

func Mean(list []int) float64 {
	var sum int
	for _, v := range list {
		sum += v
	}
	return float64(sum) / float64(len(list))
}

func StDev(list []int) float64 {
	mean := Mean(list)
	var sum float64
	for _, v := range list {
		x := float64(v)
		sum += (x - mean) * (x - mean)
	}
	return math.Sqrt(sum / float64(len(list)))

}

func Less(i, j int) bool {
	return i < j
}

func (d dataPoint) Start() {
	d.t0 = time.Now()
}

func (d dataPoint) Stop() {
	d.t1 = time.Now()
}

func (d dataPoint) Collect() {
	d.Start()
	d.data = d.fn()
	d.Stop()
}

func GetData() DataPoint {
	d := dataPoint{}
	d.Collect()
	return d
}
