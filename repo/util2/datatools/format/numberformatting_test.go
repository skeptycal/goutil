// Package format contains functions that format numeric values.
package format

import (
	"fmt"
	"testing"
)

const (
	input = "The quick brown 狐 jumped over the lazy 犬"
)

func ExampleReverseRune() {
    fmt.Println(input)
    fmt.Println(ReverseRune(input))
    // Output:
    // The quick brown 狐 jumped over the lazy 犬
    // 犬 yzal eht revo depmuj 狐 nworb kciuq ehT
}

type Any interface{}

type AnyFunc func(args ...Any) Any

type FuncType func(s string) string

type BenchMark struct {
    name string
    f FuncType
}

func (b *BenchMark) Name() string {
    return fmt.Sprintf("Benchmark: %s: %v", b.name, b.f)
}

// func (b *BenchMark) Run(args ...Any) Any {

// }

type BenchMarkSet []*BenchMark

func BenchmarkReverse(b *testing.B) {
    benchmarksReverse := BenchMarkSet{
        {"",Reverse},
    }

    for _, bb := range benchmarksReverse {
        for i := 0; i < b.N; i++ {
            bb.f(bb.name)
        }
    }
}



// func TestNumSpace(t *testing.T) {
// 	type args struct {
// 		n string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 		{"12345.54321e42", args{"12345.54321e42"}, "12345.54321e42"},
// 		{"1", args{"1"}, "1"},
// 		{"-1", args{"-1"}, "-1"},
// 		{"0.123", args{"0.123"}, "0.123"},
// 		{"-43.3234e-105", args{"-43.3234e-105"}, "-43.3234e-105"},
// 		{"1234567890.09876543210", args{"1234567890.09876543210"}, "1234567890.09876543210"},
// 		{input, args{input}, Reverse(input) + "..."},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NumSpace(tt.args.n); got != tt.want {
// 				t.Errorf("NumSpace() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// using []byte variable and string conversion
// BenchmarkReverse-8   	18962943	        60.7 ns/op	      16 B/op	       2 allocs/op
// using strings.Builder
// BenchmarkReverse-8     	29489194	        38.7 ns/op	       8 B/op	       1 allocs/op
// func BenchmarkReverse(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Reverse("12345")
// 		Reverse(input)
// 	}
// }

// BenchmarkReverse2-8   	13703583	        83.6 ns/op	       8 B/op	       1 allocs/op

// func BenchmarkReverse2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Reverse2("12345")
// 		Reverse2(input)
// 	}
// }

// // BenchmarkReverse3-8   	10324681	       108 ns/op	      40 B/op	       2 allocs/op

// func BenchmarkReverse3(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Reverse3("12345")
// 		Reverse3(input)
// 	}
// }

// // BenchmarkReverse4-8    	13986646	        82.5 ns/op	       8 B/op	       1 allocs/op

// func BenchmarkReverse4(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Reverse4("12345")
// 		Reverse4(input)
// 	}
// }

// func BenchmarkReverse5(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Reverse5("12345")
// 		Reverse5(input)
// 	}
// }

// func BenchmarkReverse8(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Reverse8("12345")
// 		Reverse8(input)
// 	}
// }

// func BenchmarkReverseRune(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		ReverseRune("12345")
// 		ReverseRune(input)
// 	}
// }

// // BenchmarkNumSpaces-8   	 2261641	       520 ns/op	      80 B/op	      12 allocs/op

// func BenchmarkNumSpaces(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		NumSpace("12345.54321e-42")
// 	}
// }

// func TestReverse8(t *testing.T) {

// }

// func TestReverse(t *testing.T) {
// 	type args struct {
// 		s string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 		{"hello", args{"hello"}, "olleh"},
// 		// {"12345", args{"12345"}, "54321"},
// 		// {"dot.net", args{"dot.net"}, "ten.tod"},
// 		{input, args{input}, Reverse4(input)},
// 	}
// 	for _, tt := range tests {
// 		// t.Run(tt.name, func(t *testing.T) {
// 		// 	if got := Reverse8(tt.args.s); got != tt.want {
// 		// 		t.Errorf("Reverse8() = %v, want %v", got, tt.want)
// 		// 	}
// 		// })

// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Reverse2(tt.args.s); got != tt.want {
// 				t.Errorf("Reverse2() = %v, want %v", got, tt.want)
// 			}
// 		})

// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Reverse3(tt.args.s); got != tt.want {
// 				t.Errorf("Reverse3() = %v, want %v", got, tt.want)
// 			}
// 		})

// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Reverse4(tt.args.s); got != tt.want {
// 				t.Errorf("Reverse4() = %v, want %v", got, tt.want)
// 			}
// 		})
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := ReverseRune(tt.args.s); got != tt.want {
// 				t.Errorf("ReverseRune() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
