// Copyright (c) 2020 Michael Treanor
// MIT License

package ansi

import (
	"fmt"
	"testing"
)

var (
	testsNewAnsi = []struct {
		name string
		args []string
		want string
	}{
		// {"no bytes(return default value)", []string{}, "\x1b[0;38;5;7;48;5;0m"},
		// {"1 byte (Green)", []string{"2"}, "\x1b[38;5;2m"},
		// {"2 bytes (fg/bg)", []string{"2", "4"}, "\x1b[38;5;2;48;5;4m"},
		{"3 bytes (fg/bg/effect)", []string{"2", "0", "1"}, "\x1b[1;38;5;2;48;5;0m"},
		// {">3 bytes (fg/bg + extra effects)", []string{"2", "4", "1", "3"}, "\x1b[1;3;38;5;2;48;5;4m"},
		{"bold black on red", []string{"0", "1", "1"}, "\x1b[1;38;5;0;48;5;1m"},
		{"italic blue on red", []string{"4", "1", "3"}, "\x1b[3;38;5;4;48;5;1m"},
	}

	funcsNewAnsi = []struct {
		name string
		fn   func(fg, bg, ef string) string
	}{
		// {"NewColor", NewColor}, // this is simply an export function for the most efficient algorithm
		{"newColorConcat", newColorConcat},
		{"newColorSprintf", newColorSprintf},
		{"newColorJoin", newColorJoin},
		{"newColorSB", newColorSB},
		{"newColorBB", newColorBB},
	}
)

func Test_newColorString(t *testing.T) {
	for _, tt := range testsNewAnsi {
		for _, fn := range funcsNewAnsi {
			name := fmt.Sprintf("%s(%s)", tt.name, fn.name)
			t.Run(name, func(t *testing.T) {
				if got := fn.fn(tt.args[0], tt.args[1], tt.args[2]); got != tt.want {
					t.Errorf("newColorString() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

// this is a noop function
// the three arguments are present to match the test function signature
// pass the expected return value to 'foreground'
// background and effect are ignored
func newColorControl(foreground, background, effect string) string {
	return foreground
}

func BenchmarkNewColor(b *testing.B) {
	// initial results
	/*
		Conclusions:
		* For small strings (the output values are on the order of 18 bytes long),
		* Concat is faster than any other method, even with repetition and varied data

		Single test (testNewColor[0])

		All Tests (looping through tests inside timed loop - all tests are within timed loop)

		All Tests (looping through tests outside timed loop - only one test is in each time loop)


		BenchmarkNewColor/NewColor(newColorControl)-8         	28835474	        40.19 ns/op	      16 B/op	       1 allocs/op
		BenchmarkNewColor/NewColor(newColorControl)#01-8      	 9267547	       135.5 ns/op	      48 B/op	       3 allocs/op
		BenchmarkNewColor/NewColor(newColorControl)#02-8      	10482974	       117.2 ns/op	      48 B/op	       3 allocs/op

		BenchmarkNewColor/NewColor(newColorConcat)-8          	 9474012	       123.6 ns/op	      40 B/op	       2 allocs/op
		BenchmarkNewColor/NewColor(newColorConcat)#01-8       	 3061848	       388.5 ns/op	     120 B/op	       6 allocs/op
		BenchmarkNewColor/NewColor(newColorConcat)#02-8       	 3252994	       369.5 ns/op	     120 B/op	       6 allocs/op

		BenchmarkNewColor/NewColor(newColorStringer)-8        	 3800276	       317.0 ns/op	      88 B/op	       5 allocs/op
		BenchmarkNewColor/NewColor(newColorStringer)#01-8     	 1240056	       977.8 ns/op	     264 B/op	      15 allocs/op
		BenchmarkNewColor/NewColor(newColorStringer)#02-8     	 1265025	       954.4 ns/op	     264 B/op	      15 allocs/op

		BenchmarkNewColor/NewColor(newColorSprintf)-8         	 3862953	       321.5 ns/op	      88 B/op	       5 allocs/op
		BenchmarkNewColor/NewColor(newColorSprintf)#01-8      	 1251404	      1001 ns/op	     264 B/op	      15 allocs/op
		BenchmarkNewColor/NewColor(newColorSprintf)#02-8      	 1259708	       956.6 ns/op	     264 B/op	      15 allocs/op

		BenchmarkNewColor/NewColor(newColorJoin)-8            	 5132965	       227.9 ns/op	      56 B/op	       5 allocs/op
		BenchmarkNewColor/NewColor(newColorJoin)#01-8         	 1683765	       715.5 ns/op	     168 B/op	      15 allocs/op
		BenchmarkNewColor/NewColor(newColorJoin)#02-8         	 1738910	       691.9 ns/op	     168 B/op	      15 allocs/op

		BenchmarkNewColor/NewColor(newColorSB)-8              	 7133187	       166.8 ns/op	      72 B/op	       4 allocs/op
		BenchmarkNewColor/NewColor(newColorSB)#01-8           	 2287580	       518.8 ns/op	     216 B/op	      12 allocs/op
		BenchmarkNewColor/NewColor(newColorSB)#02-8           	 2370159	       510.8 ns/op	     216 B/op	      12 allocs/op

		BenchmarkNewColor/NewColor(newColorBB)-8              	 6663630	       181.6 ns/op	     104 B/op	       3 allocs/op
		BenchmarkNewColor/NewColor(newColorBB)#01-8           	 2151378	       557.7 ns/op	     312 B/op	       9 allocs/op
		BenchmarkNewColor/NewColor(newColorBB)#02-8           	 2203754	       548.1 ns/op	     312 B/op	       9 allocs/op
	*/
	var result Any
	var r Any
	var bFuncTests = []struct {
		name string
		fn   func(fg, bg, ef string) string
	}{
		{"newColorControl", newColorControl},
	}
	bFuncTests = append(bFuncTests, funcsNewAnsi...)

	for _, fn := range bFuncTests {
		name := fmt.Sprintf("%s(%s)", "NewColor", fn.name)

		// single test
		b.Run(name, func(b *testing.B) {
			bb := testsNewAnsi[0]
			for i := 0; i < b.N; i++ {
				r = fn.fn(bb.args[0], bb.args[1], bb.args[2])
			}
		})

		// with inner loop
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, bb := range testsNewAnsi {
					r = fn.fn(bb.args[0], bb.args[1], bb.args[2])
				}
			}
		})

		// with outer loop
		b.Run(name, func(b *testing.B) {
			for _, bb := range testsNewAnsi {
				for i := 0; i < b.N; i++ {
					r = fn.fn(bb.args[0], bb.args[1], bb.args[2])
				}
			}
		})
	}
	result = r
	_ = result
}
