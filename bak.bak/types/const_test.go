package types

import (
	"testing"
)

func TestConst(t *testing.T) {
	var (
		esc     byte = '\x1b'
		nl      byte = '\n'
		tab     byte = '\t'
		space   byte = ' '
		bracket byte = '['
		zero    byte = '0'
		m       byte = 'm'

		ansiReset = []byte{esc, bracket, zero, m}
	)

	tests := []struct {
		want    string
		got     string
		wantErr bool
	}{
		{string(esc), ESC, false},
		{string(nl), NL, false},
		{string(tab), TAB, false},
		{string(space), SPACE, false},
		{string(ansiReset), AnsiReset, false},
		{"fail", "fake", true},
	}

	for _, tt := range tests {
		t.Run("Constant("+tt.want+")", func(t *testing.T) {
			if tt.got != tt.want != tt.wantErr {
				t.Errorf("incorrect constant: got %q, want %q", tt.got, tt.want)
			}
		})
	}
}
