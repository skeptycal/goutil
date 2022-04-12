package main

import (
	"testing"

	"github.com/skeptycal/goutil/data/pooler"
)

/*
results on macOS arm test machine:

	original design:
	numberOfJobs = 20:
	BenchmarkConcurrent-8      1	 3003465833 ns/op	9680 B/op 105 allocs/op
	BenchmarkNonconcurrent-8   1	20018518667 ns/op	2120 B/op  49 allocs/op

	goos: darwin
	goarch: arm64
	pkg:

	numberOfJobs = 5:
	BenchmarkConcurrent-8    100	  500511354 ns/op	 790 B/op  30 allocs/op
	BenchmarkNonconcurrent-8   1	 5004593250 ns/op	 848 B/op  40 allocs/op

	numberOfJobs = 20:
	BenchmarkConcurrent-8      1	 3006381083 ns/op	9400 B/op 157 allocs/op
	BenchmarkNonconcurrent-8   1	20019600375 ns/op	2880 B/op 117 allocs/op

	numberOfJobs = 50:
	BenchmarkConcurrent-8      1	 9008377500 ns/op	3360 B/op 314 allocs/op
	BenchmarkNonconcurrent-8   1	50053344791 ns/op	6496 B/op 269 allocs/op
*/

const numberOfJobs = 20

func BenchmarkConcurrent(b *testing.B) {
	collector := pooler.StartDispatcher(WORKER_COUNT) // start up worker pool

	for n := 0; n < b.N; n++ {
		for i, job := range pooler.CreateExampleJobs(numberOfJobs) {
			collector.Work <- pooler.Work{Job: job, ID: i}
		}
	}
}

func BenchmarkNonconcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, work := range pooler.CreateExampleJobs(numberOfJobs) {
			// pooler.ExampleWork(work, 1)
			pooler.ExampleWork(pooler.Params{"job": work, "id": 1})
		}
	}
}
