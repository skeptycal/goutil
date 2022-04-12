package shpath

import (
	"testing"
)

var fakeString = ""

var stringFuncs = []struct {
	name string
	fn   func(string) string
}{
	{"(cache loading ... ignore this one)", NormalizeNL},
	{"for loop", normalizeNLForLoop},
	{"bytes.Replace", normalizeNewlinesBytesWrapper},
	{"strings.Replace", normalizeNewlinesString},
	{"aliased", NormalizeNL},
	{"wrapped", normalizeWrapped},
	// {"global strings.Builder", normalizeGlobalStringsBuilder},
	{"use strings.Builder", normalizeStringsBuilder},
}

var stringModifyFuncs = []struct {
	name string
	fn   func(string, string) string
}{
	{"DropDupes", DropDupes},
}

// BenchmarkNormalize/for_loop-8         	 2511970	       529.6 ns/op	     336 B/op	       3 allocs/op
// BenchmarkNormalize/bytes.Replace-8    	 2226400	       532.8 ns/op	     448 B/op	       4 allocs/op
// BenchmarkNormalize/strings.Replace-8  	 3018738	       497.9 ns/op	     224 B/op	       2 allocs/op
func BenchmarkNormalize(b *testing.B) {
	arg := "asdlfkn2;leja-9cv8yh	-2piouej4b-	2u9hnasdj;lasdjflkasnvj8q92nn2den\rasdfklw\r\nl;jkqw;cijhpoiqjwd\n\njl-9c8vn-	wd"

	for _, bb := range stringFuncs {
		b.Run(bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fakeString = bb.fn(arg)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1234567890", args{"1234567890"}, "1234567890"},
		{"\n1234567890", args{"\n1234567890"}, "\n1234567890"},
		{"\r1234567890", args{"\r1234567890"}, "\n1234567890"},
		{"\r\n1234567890", args{"\r\n1234567890"}, "\n1234567890"},
		{"12345\n67890", args{"12345\n67890"}, "12345\n67890"},
		{"12345\r67890", args{"12345\r67890"}, "12345\n67890"},
		{"12345\r\n67890", args{"12345\r\n67890"}, "12345\n67890"},
		{"fake", args{"fake"}, "fake"},
		{"fa\rke", args{"fa\rke"}, "fa\nke"},
		{"fa\r\nke", args{"fa\r\nke"}, "fa\nke"},
	}
	for _, tt := range tests {
		for _, ff := range stringFuncs {
			name := ff.name + "(" + tt.name + ")"
			t.Run(name, func(t *testing.T) {
				if got := ff.fn(tt.args.s); got != tt.want {
					t.Errorf("%q = %q, want %q", name, got, tt.want)
				}
			})
		}
	}
}

func TestDropDupes(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1234567890", args{"1234567890", "5"}, "1234567890"},
		{"fake", args{"fake", "e"}, "fake"},
		{"newlines", args{"new\n\nline\n", "\n"}, "new\nline\n"},
		{"eee's", args{"slender feet", "e"}, "slender fet"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DropDupes(tt.args.s, tt.args.sep); got != tt.want {
				t.Errorf("DropDupes() = %v, want %v", got, tt.want)
			}
		})
	}
}
