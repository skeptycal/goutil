// Package polynomial provides functions that support polynomial arithmetic.
package polynomial

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

var DevNull DevNullReadWriter

type DevNullReadWriter struct{ io.ReadWriteCloser }

func (w DevNullReadWriter) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)
	defer buf.Reset()
	return buf.Len(), nil
}

func (w DevNullReadWriter) Read(p []byte) (n int, err error) {
	buf := bytes.NewBuffer(p)
	defer buf.Reset()
	return buf.Len(), nil
}

func init() {
	DevNull, err := os.OpenFile(os.DevNull, os.O_RDWR, 0777)
	if err != nil {
		fmt.Fprint(os.Stderr, "failed to initialize DevNull")
	}
	fmt.Fprint(DevNull, "DevNull")
}


/* Benchmark results

------------------------------------------------------
----- with Free()
------------------------------------------------------
BenchmarkList_String/00000-8         	 6760014	       154 ns/op	      32 B/op	       2 allocs/op
BenchmarkList_String/00001-8         	 6041767	       204 ns/op	      40 B/op	       3 allocs/op
BenchmarkList_String/___-1-8         	 5969319	       206 ns/op	      40 B/op	       3 allocs/op
BenchmarkList_String/12345-8         	 5660700	       217 ns/op	      40 B/op	       3 allocs/op

------------------------------------------------------
----- with set to nil instead of Free()
------------------------------------------------------
BenchmarkList_String/00000-8         	 7704391	       168 ns/op	      32 B/op	       2 allocs/op
BenchmarkList_String/00001-8         	 5056119	       228 ns/op	      40 B/op	       3 allocs/op
BenchmarkList_String/___-1-8         	 5677894	       217 ns/op	      40 B/op	       3 allocs/op
BenchmarkList_String/12345-8         	 5503621	       216 ns/op	      40 B/op	       3 allocs/op

------------------------------------------------------
----- with no Free() or nil
------------------------------------------------------
BenchmarkList_String/00000-8         	 6933988	       158 ns/op	      32 B/op	       2 allocs/op
BenchmarkList_String/00001-8         	 5340334	       227 ns/op	      40 B/op	       3 allocs/op
BenchmarkList_String/___-1-8         	 5442990	       228 ns/op	      40 B/op	       3 allocs/op
BenchmarkList_String/12345-8         	 5457669	       247 ns/op	      40 B/op	       3 allocs/op

------------------------------------------------------
----- using LoadInt() instead of New()
------------------------------------------------------
BenchmarkList_String/00000-8         	 7447820	       157 ns/op	      32 B/op	       2 allocs/op
BenchmarkList_String/00001-8         	 7851932	       158 ns/op	      32 B/op	       2 allocs/op
BenchmarkList_String/___-1-8         	 7911214	       157 ns/op	      32 B/op	       2 allocs/op
BenchmarkList_String/12345-8         	 7652253	       159 ns/op	      32 B/op	       2 allocs/op

------------------------------------------------------
----- (global list using LoadInt) with no io.ReadWriter output (with Free)
------------------------------------------------------
BenchmarkList_String/00000-8         	360265500	         3.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkList_String/00001-8         	367157193	         3.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkList_String/___-1-8         	374736079	         3.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkList_String/12345-8         	367849651	         3.19 ns/op	       0 B/op	       0 allocs/op

------------------------------------------------------
----- (new local list using New) with no io.ReadWriter output (with Free)
------------------------------------------------------
BenchmarkList_String/00000-8         	349358112	         3.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkList_String/00001-8         	31779576	        37.4 ns/op	       8 B/op	       1 allocs/op
BenchmarkList_String/___-1-8         	31469778	        38.9 ns/op	       8 B/op	       1 allocs/op
BenchmarkList_String/12345-8         	21438880	        56.8 ns/op	       8 B/op	       1 allocs/op
*/
func BenchmarkList_String(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
		// fun  func() string
	}{
		{"00000", 0},
		{"00001", 1},
		{"   -1", -1},
		{"12345", 12345},
    }
    list := New(0)
	for _, bb := range benchmarks {
		list.LoadInt(bb.n)
		// defer list.Free()

		b.Run(bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
                // fmt.Fprint(DevNull, list.String())
                _ = list.String()
			}
        })

	}
}

func TestList_String(t *testing.T) {
	tests := []struct {
		name string
        n    int
        want string
	}{
		{"00000", 0, "0"},
		{"00001", 1, "1"},
		{"   -1", -1, "-1"},
		{"12345", 12345, "54321"},
    }
    list := New(0)
    var got string
	for _, tt := range tests {
		// defer list.Free()

		t.Run(tt.name, func(t *testing.T) {
            list = New(tt.n)
            got = list.LoadInt(tt.n)
            got = list.String()
            if tt.want != got {
                t.Errorf("String() = %v, want %v", got, tt.want)
            }
        })

	}
}

func TestList_String2(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		// TODO: Add test cases.
		{"0", 0, "0"},
		{"12345", 12345, "12345"},
		{"1", 1, "1"},
		{"-1", -1, "-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.n).String()
			if got != tt.want {
				t.Errorf("List.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"12345", args{12345}, "54321"},
		{"0", args{0}, "0"},
		{"503050300", args{503050300}, "003050305"},
		{"-12345", args{-12345}, "-54321"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringDigits(tt.args.n); got != tt.want {
				t.Errorf("StringDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want List
	}{
		// TODO: Add test cases.
		{"15", args{15}, List{false, &ListNode{1, &ListNode{5, nil}}, nil}},
		{"25", args{25}, List{false, &ListNode{2, &ListNode{5, nil}}, nil}},
		{"1234567890", args{1234567890}, List{false, &ListNode{1, &ListNode{2, nil}}, nil}},
		{"0", args{0}, List{false, &ListNode{0, nil}, nil}},
		{"-12345", args{-12345}, List{true, &ListNode{1, &ListNode{2, nil}}, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO - these tests are not thorough and exhaustive ...
			// TODO - only bool and first 2 ints are tested
			got := *New(tt.args.n)
			want := tt.want
			if got.SignBit != want.SignBit {
				t.Errorf("New() sign bit = %v, want %v", got.SignBit, want.SignBit)
			}
			if got.First.Val != want.First.Val {
				t.Errorf("New() first value = %v, want %v", got.First.Val, want.First.Val)
			}
			if got.First.Next != nil && want.First.Next != nil {
				if got.First.Next.Val != want.First.Next.Val {
					t.Errorf("New() second value = %v, want %v", got.First.Next.Val, want.First.Next.Val)
				}
			}
		})
	}
}



func TestListDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want *List
	}{
		// TODO: Add test cases.
		{"1", args{1}, &List{false, &ListNode{1, nil}, nil}},
		{"-1", args{-1}, &List{true, &ListNode{1, nil}, nil}},
		{"0", args{0}, &List{false, &ListNode{0, nil}, nil}},
		{"12345", args{12345}, &List{false, &ListNode{5, nil}, nil}},
		{"-54321", args{-54321}, &List{true, &ListNode{1, nil}, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO - these tests are not thorough and exhaustive ...
			// TODO - only bool and first 2 ints are tested
			got := *ListDigits(tt.args.n)
			want := tt.want
			if got.SignBit != want.SignBit {
				t.Errorf("New() sign bit = %v, want %v", got.SignBit, want.SignBit)
			}
			if got.First.Val != want.First.Val {
				t.Errorf("New() first value = %v, want %v", got.First.Val, want.First.Val)
			}
			if got.First.Next != nil && want.First.Next != nil {
				if got.First.Next.Val != want.First.Next.Val {
					t.Errorf("New() second value = %v, want %v", got.First.Next.Val, want.First.Next.Val)
				}
			}
		})
	}
}

func TestList_LoadInt(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		// TODO: Add test cases.
		{"1", 1, "1"},
		{"-1", -1, "-1"},
		{"0", 0, "0"},
		{"12345", 12345, "54321"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := List{}
			if got := l.LoadInt(tt.n); got != tt.want {
				t.Errorf("List.LoadInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
