package testes

import (
	"reflect"
	"testing"
)

func TRun(t *testing.T, name string, got, want Any) {
	if NewAnyValue(got).IsComparable() && NewAnyValue(want).IsComparable() {
		t.Run(name, func(t *testing.T) {
			if got != want {
				if !reflect.DeepEqual(got, want) {
					TError(t, name, got, want)
				}
			}
		})
	}
}

func TTypeRun(t *testing.T, name string, got, want Any, wantErr bool) {
	if NewAnyValue(got).IsComparable() && NewAnyValue(want).IsComparable() {
		t.Run(name, func(t *testing.T) {
			if got != want != wantErr {
				if !reflect.DeepEqual(got, want) {
					TTypeError(t, name, got, want)
				}
			}
		})
	}
}

func TRunTest(t *testing.T, tt *test) {
	if NewAnyValue(tt.got).IsComparable() && NewAnyValue(tt.want).IsComparable() {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want != tt.wantErr {
				if reflect.DeepEqual(tt.got, tt.want) == tt.wantErr {
					TError(t, tt.name, tt.got, tt.want)
				}
			}
		})
	}
}
