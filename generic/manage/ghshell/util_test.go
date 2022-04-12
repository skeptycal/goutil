package ghshell

import (
	"reflect"
	"testing"
)

func TestNormalizeWhitespace(t *testing.T) {
	type args struct {
		s            string
		saveNewLines bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"short", args{s: "", saveNewLines: true}, "this is short"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeWhitespace(tt.args.s, tt.args.saveNewLines); got != tt.want {
				t.Errorf("NormalizeWhitespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toFields(t *testing.T) {
	NUL := string([]byte{0})
	tests := []struct {
		name string
		in   []string
		want []string
	}{
		// TODO: Add test cases.
		{"ls -lah", []string{"ls", "-lah"}, []string{"ls", "-lah"}},
		{"\nls -lah", []string{"\nls", "-lah"}, []string{"ls", "-lah"}},
		{"ls\n\t-lah", []string{"ls\n", "\t-lah"}, []string{"ls", "-lah"}},
		{"\nls " + NUL + "-lah", []string{"ls", NUL + "-lah"}, []string{"ls", "-lah"}},
		{"ls     -lah", []string{"ls", "   -lah"}, []string{"ls", "-lah"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFields(tt.in...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toFields(%q) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
