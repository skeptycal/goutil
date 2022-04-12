package anansi

import (
	"os"
	"reflect"
	"testing"
)

func TestNewColor(t *testing.T) {
	type args struct {
		value []Attribute
	}
	tests := []struct {
		name string
		args args
		want *Color
	}{
		{"reset", args{value: []Attribute{ResetCode}}, NewColor(ResetCode)},
		{"redbold", args{value: []Attribute{FgRed, Bold}}, NewColor(FgRed, Bold)},
		{"blueitalic", args{value: []Attribute{FgBlue, Italic}}, NewColor(FgBlue).Add(Italic)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewColor(tt.args.value...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttn(t *testing.T) {
	tests := []struct {
		name   string
		format string
		args   []interface{}
	}{
		// TODO: Add test cases.
		{"ATTN example ", "(yellow bold text on red background)", []interface{}{FgYellow, BgHiRed, Bold}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Attn(tt.format, tt.args...)
		})
	}
}

func ExampleNewColor() {
	in := "(yellow bold log INFO text on red background)"
	NewColor(FgYellow, BgHiRed, Bold).Fprintln(os.Stdout, in)
	// Attn(in)

	// output:
	// (yellow bold log INFO text on red background)
}
