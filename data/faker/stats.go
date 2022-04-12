package faker

import (
	"math"
	"time"
)

const aRune = 'A'

type (
	// StatMap implements statistical funcionality on
	// a data structure. The basis is an underlying
	// map that allows accessing data by key or by
	// index.
	//
	// The keys are assumed to be strings and the
	// values are float64 values.
	StatMap interface {

		// Keys returns a list of keys from
		// the underlying data structure.
		Keys() []string

		// Values returns a list of values from
		// the underlying data structure.
		Values() []float64

		// Min returns the minimum value.
		Min() float64

		// Max returns the maximum value.
		Max() float64

		// Mean returns the mean of the data set.
		Mean() float64

		// StDev returns the standard deviation
		// of the data set.
		StDev() float64

		String() string
	}

	dataPoint struct {
		t0   time.Time
		t1   time.Time
		fn   func() Any
		data Any
	}

	// DataPoint implements a data collection
	// device capable of recording one piece
	// of data and the start and stop times
	// required for the operation.
	//
	// It is based on a data collection function
	// that collects the data and records the
	// time values as required.
	DataPoint interface {
		Start()
		Stop()
		Collect()
	}
)

func dev(x, m float64) float64 {
	return (x - m) * (x - m)
}

func Mean(list []float64) float64 {
	var sum float64
	for _, v := range list {
		sum += v
	}
	return sum / float64(len(list))
}

func meanWithCount(list []float64) float64 {
	var v, sum float64
	var i int
	for i, v = range list {
		sum += v
	}
	return sum / float64(i)
}

func StDev(list []float64) float64 {
	mean := Mean(list)
	var sum float64
	for _, v := range list {
		x := float64(v)
		sum += (x - mean) * (x - mean)
	}
	return math.Sqrt(sum / float64(len(list)))
}

func stDevWithCounter(list []float64) float64 {
	mean := Mean(list)
	var v, sum float64
	var i int
	for i, v = range list {
		x := float64(v)
		sum += (x - mean) * (x - mean)
	}
	return math.Sqrt(sum / float64(i))
}

func Less(i, j int) bool {
	return i < j
}
