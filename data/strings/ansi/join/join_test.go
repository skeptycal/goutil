package join

import (
	"bytes"
	"math/rand"
	"strings"
	"testing"
	"time"
)

// benchmark results
/*	/// Conclusions
	/// - For tiny strings, never use Sprintf if performance is any issue at all.
	/// - Once strings read size 2048, Sprintf catches up to all but concat
	/// - Instead, use it to setup constants and prefixes, then use other methods in loops
	/// - Concatenate is much better than most people think .. from what I can see online
*/

// n=5, size=1 (original args and loops)
/*
	reference: https://gist.github.com/dtjm/c6ebc86abe7515c988ec
	/// if these are concrete (constant) strings, the results of concat are much better
	/// apparently not Join, Builder, Buffer outclass other options
	BenchmarkConcat-8            	13583527	        77.42 ns/op	       0 B/op	       0 allocs/op
	BenchmarkJoin-8              	13136424	        83.52 ns/op	      16 B/op	       1 allocs/op
	BenchmarkStringsBuilder-8    	13841370	        87.40 ns/op	      24 B/op	       2 allocs/op
	BenchmarkBuffer-8            	12019822	        96.20 ns/op	      64 B/op	       1 allocs/op
	BenchmarkSprintf-8           	 2813542	        442.7 ns/op	      96 B/op	       6 allocs/op
*/

// /// 5 single bytes ... simple, right?
/*
	/// apparently not Join, Builder, Buffer outclass other options
	BenchmarkConcat-8           	 2926448	       410.0 ns/op	      80 B/op	       9 allocs/op
	BenchmarkJoin-8             	13808566	        82.92 ns/op	      16 B/op	       1 allocs/op
	BenchmarkStringsBuilder-8   	12270021	        95.41 ns/op	      24 B/op	       2 allocs/op
	BenchmarkBuffer-8           	11209058	       102.4 ns/op	      64 B/op	       1 allocs/op
	BenchmarkSprintf-8          	 1275196	       939.3 ns/op	     192 B/op	      14 allocs/op
*/

// 5, 1 (updated for variable test data - single line concat)
/*
	/// concat is only better when doing all work on a single line ... loops destroy it
	/// even a single extra concat step (one for variable and one for separator) doubled the time
	BenchmarkConcat-8           	 4606387	       255.7 ns/op	      48 B/op	       5 allocs/op
	BenchmarkJoin-8             	14277806	       88.89 ns/op	      16 B/op	       1 allocs/op
	BenchmarkStringsBuilder-8   	12339235	       99.29 ns/op	      24 B/op	       2 allocs/op
	BenchmarkBuffer-8           	11610976	       103.3 ns/op	      64 B/op	       1 allocs/op
	BenchmarkSprintf-8          	 1280098	       959.6 ns/op	     192 B/op	      14 allocs/op
*/

// 5, 1<<16
/*
	/// for few large strings, Join is much better in every way
	/// concat and sprintf lose ground quickly
	/// bytes.Buffer and strings.Builder remain competitive and similar
	BenchmarkConcat-8           	    3074	    420840 ns/op	 2859014 B/op	       9 allocs/op
	BenchmarkJoin-8             	   13700	     87580 ns/op	  491521 B/op	       1 allocs/op
	BenchmarkStringsBuilder-8   	    5396	    220215 ns/op	 1613828 B/op	       5 allocs/op
	BenchmarkBuffer-8           	    5874	    211659 ns/op	 1302532 B/op	       4 allocs/op
	BenchmarkSprintf-8          	    2029	    625249 ns/op	 4065628 B/op	      31 allocs/op
*/

// 1<<16 , 1
/*
	/// many small strings overwhelms concat and sprintf ...
	/// strings.Builder stays ahead in speed, but Join has much lower allocation
	BenchmarkConcat-8           	       1	1388670367 ns/op	 5593303640 B/op	   65730 allocs/op
	BenchmarkJoin-8             	    1509	    831126 ns/op	  	 172032 B/op	       1 allocs/op
	BenchmarkStringsBuilder-8   	    1789	    667591 ns/op	  	 858107 B/op	      27 allocs/op
	BenchmarkBuffer-8           	    1263	    951043 ns/op	  	 733618 B/op	      14 allocs/op
	BenchmarkSprintf-8          	       1	2486296949 ns/op	10279359488 B/op	  279592 allocs/op
*/

// 1<<16-1, 1 /// curious ...
/*
	/// almost the same ...
	/// does join use a strings.Builder? it seems to be very similar, but more efficient at smaller sizes
	/// haha, it does ... with guards for very small arguments
		BenchmarkConcat-8           	       1	1336306879 ns/op	 5593130216 B/op	   65709 allocs/op
		BenchmarkJoin-8             	    1592	    757898 ns/op	     172032 B/op	       1 allocs/op
		BenchmarkStringsBuilder-8   	    1846	    719259 ns/op	     858108 B/op	      27 allocs/op
		BenchmarkBuffer-8           	    1160	   1007859 ns/op	     733616 B/op	      14 allocs/op
		BenchmarkSprintf-8          	       1	2659282056 ns/op		10279512752 B/op	  279565 allocs/op

		BenchmarkAppend-8           	    2883	    420550 ns/op	  	505850 B/op	      25 allocs/op
		BenchmarkJoin-8             	    1466	    816025 ns/op	  	172032 B/op	       1 allocs/op
		BenchmarkStringsBuilder-8   	     979	   1246182 ns/op	  	999424 B/op	       4 allocs/op
		BenchmarkBuffer-8           	    1225	    954970 ns/op	  	733616 B/op	      14 allocs/op
		BenchmarkConcat-8           	       1	1355541871 ns/op	5593129200 B/op	   65703 allocs/op
		BenchmarkSprintf-8          	       1	2593540327 ns/op   10279745960 B/op	  279572 allocs/op
	/// adding those strings.Join "guards" ...
*/

// 1<<16 - 1 , 1024
/* /// this is just ridiculous ... concat and sprintf could not finish ...
BenchmarkAppend-8           	      14	  81130664 ns/op	529470726 B/op	      44 allocs/op
BenchmarkJoin-8             	      45	  29993178 ns/op	100605954 B/op	       1 allocs/op
BenchmarkBuffer-8           	      16	  69543716 ns/op	400996864 B/op	      18 allocs/op
BenchmarkStringsBuilder-8   	       8	 136004574 ns/op	825688124 B/op	       5 allocs/op
* BenchmarkConcat-8           	       1	389421211038 ns/op	3296694767992 B/op	   67138 allocs/op
*/

const (
	n    int    = 1<<8 - 1
	size int    = 1<<8 - 1
	sep  string = ":"
)

// var testData = []string{"a", "b", "c", "d", "e"}
var testData []string

func BenchmarkTestData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testData = generateTestData(n, size)
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b []byte
		for _, line := range testData {
			b = append(b, line...) // s + line + ":"
		}
		_ = b
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string = strings.Join(testData, ":")
		_ = s
	}
}

// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func Join(elems []string, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(elems[0])
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}

func BenchmarkStringsBuilderPrep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch len(testData) {
		case 0:
			_ = ""
		case 1:
			_ = testData[0]
		}

		n := len(sep) * (len(testData) - 1)

		for i := 0; i < len(testData); i++ {
			n += len(testData[i])
		}

		var b strings.Builder
		b.Grow(n)
		b.WriteString(testData[0])
		for _, s := range testData[1:] {
			b.WriteString(sep)
			b.WriteString(s)
		}
		_ = b.String()

		for _, line := range testData {
			b.WriteString(line)
			b.WriteByte(':')
		}
		_ = b.String()
	}
}

func BenchmarkStringsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch len(testData) {
		case 0:
			_ = ""
		case 1:
			_ = testData[0]
		}

		// n := len(sep) * (len(testData) - 1)

		// for i := 0; i < len(testData); i++ {
		// 	n += len(testData[i])
		// }

		var b strings.Builder
		b.Grow(n)
		b.WriteString(testData[0])
		for _, s := range testData[1:] {
			b.WriteString(sep)
			b.WriteString(s)
		}
		_ = b.String()

		for _, line := range testData {
			b.WriteString(line)
			b.WriteByte(':')
		}
		_ = b.String()
	}
}

func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b bytes.Buffer
		for _, line := range testData {
			b.WriteString(line)
			b.WriteByte(':')
		}
		_ = b.String()
	}
}

// func BenchmarkConcat(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		var s string
// 		for _, line := range testData {
// 			s = s + line + ":"
// 		}
// 		_ = s
// 	}
// }

// func BenchmarkSprintf(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		var s string
// 		for _, line := range testData {
// 			s = fmt.Sprintf("%s:%s", s, line)
// 		}
// 		_ = s
// 	}
// }

func init() {
	rand.Seed(time.Now().Unix())
}

func generateRandomString(n int) (retval string) {
	for i := 0; i < n; i++ {
		retval += string(byte(rand.Intn(255)))
	}
	return
}

func generateTestData(n, size int) (retval []string) {
	const (
		r         = 1<<64 - 1
		r1 uint64 = r >> 8
		r2 uint64 = r1 >> 8
		r3 uint64 = r2 >> 8
		r4 uint64 = r3 >> 8
		r5 uint64 = r4 >> 8
		r6 uint64 = r5 >> 8
		r7 uint64 = r6 >> 8
		r8 uint64 = r7 >> 8
	)
	b := []byte{}
	for i := 0; i < n; i++ {

		for j := 0; j < size<<2; j++ {
		}
		r := rand.Uint64()
		for r > 1 {
			rb := r & 8
			b = append(b, byte(rb))
		}

		// retval = append(retval, generateRandomString(size))
	}
	return
}
