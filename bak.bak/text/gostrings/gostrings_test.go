package gostrings

import (
	"errors"
	"testing"
)

var errFake = errors.New("fake test error")

func testFakeRev(s string) (string, error) {
	return Reverse(s), errFake
}

func testFakeNoop(s string) (string, error) {
	return s, errFake
}

func TestStrErrorConvert(t *testing.T) {

	tests := []struct {
		name    string
		fn      func(string) (string, error)
		in      string
		errIn   error
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"rev 12345", testFakeRev, "12345", errFake, "54321", false},
		{"rev 12345", testFakeNoop, "12345", errFake, "12345", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StrErrorConvert(tt.fn(tt.in))
			log.Info("got: ", got)
			if got != tt.want == tt.wantErr {
				t.Errorf("StrErrorConvert(fn) = %q, want %q (wantErr: %v)", got, tt.want, tt.wantErr)
			}
		})
	}
}
