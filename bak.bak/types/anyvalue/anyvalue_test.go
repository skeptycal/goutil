package anyvalue

import (
	"reflect"
	"testing"
)

func TestNewAnyValue(t *testing.T) {
	type args struct {
		a Any
	}
	tests := []struct {
		name string
		args args
		want AnyValue
	}{
		// TODO: Add test cases.
		{"42", args{"42"}, NewAnyValue("42")},
		{"42", args{NewAnyValue("42")}, NewAnyValue("42")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnyValue(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnyValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
