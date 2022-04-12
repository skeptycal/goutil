package main

import "testing"

func TestBreakIt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"short", args{s: "short"}, "short"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BreakIt(tt.args.s); got != tt.want {
				t.Errorf("BreakIt() = %v, want %v", got, tt.want)
			}
		})
	}
}
