package osargsutils

import "testing"

// Benchmark Results
/*
/// It does not matter ... in fact, simply calling the argzero function to get the initial
/// path is what takes up all of the time.

/// Will any of this matter? Not at all ... this function is likely to be called only once
/// during the runtime of the program. a few hundred nanoseconds will not matter ...

/// however, if 100,000 of these are started up in a scaling architecture, run quickly,
/// and exit nearly immediately ... well then the startup code has a huge effect.

BenchmarkHereMe2-8            	   23202	     55725 ns/op	    3696 B/op	      42 allocs/op
BenchmarkHereMe-8             	   24092	     48869 ns/op	    3696 B/op	      42 allocs/op
BenchmarkZeroOsExecutable-8   	   25843	     46830 ns/op	    3696 B/op	      42 allocs/op
BenchmarkZeroOsArgs-8         	   25575	     47170 ns/op	    3696 B/op	      42 allocs/op
BenchmarkRawOsArgsZero-8      	   25155	     48109 ns/op	    3695 B/op	      42 allocs/op
*/

var (
	pathName, baseName string
	err                error
)

func BenchmarkHereMe2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pathName, baseName, err = hereMe2()
	}
}

func BenchmarkHereMe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pathName, baseName, err = hereMe()
	}
}

func BenchmarkZeroOsExecutable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pathName, err = zeroOsExecutable()
	}
}

func BenchmarkZeroOsArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pathName, err = zeroOsArgs()
	}
}

func BenchmarkRawOsArgsZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pathName, err = rawOsArgsZero()
	}
}
