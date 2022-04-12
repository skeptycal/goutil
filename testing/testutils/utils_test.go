package testutils

import (
	"reflect"
	"testing"
)

func TestGetType(t *testing.T) {
	type args struct {
		any Any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetType(tt.args.any); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetKind(t *testing.T) {
	tests := []struct {
		name string
		any  Any
		want string
	}{
		// TODO: Add test cases.
		{"int", 42, "int"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetKind(tt.any); got != tt.want {
				t.Errorf("GetKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_remove(t *testing.T) {
	type args struct {
		slice []int
		i     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"12345, 3", args{[]int{1, 2, 3, 4, 5}, 3}, []int{1, 2, 3, 5}},
		{"12345, 4", args{[]int{1, 2, 3, 4, 5}, 4}, []int{1, 2, 3, 4}},
		{"12345, 0", args{[]int{1, 2, 3, 4, 5}, 0}, []int{2, 3, 4, 5}},
		{"index too large", args{[]int{1, 2, 3, 4, 5}, 10}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := remove(tt.args.slice, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckType(t *testing.T) {
	type args struct {
		any Any
		typ string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"string", args{"string1", "string"}, true},
		{"not string", args{"string1", "int"}, false},
		{"nil", args{nil, "<nil>"}, true},
		{"not nil", args{nil, "int"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckType(tt.args.any, tt.args.typ); got != tt.want {
				t.Errorf("CheckType(%v,%v) = %v, want %v", tt.args.any, tt.args.typ, got, tt.want)
			}
		})
	}
}
