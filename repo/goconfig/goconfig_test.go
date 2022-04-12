package goconfig

import "testing"

func Test_tes(t *testing.T) {
	tests := []struct {
		name string
		s    string
		got  *string
		want string
	}{
		// TODO: Add test cases.
		{"fake", "fake", "fake", "fake"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tes(tt.s, tt.got, tt.want)
			if tt.got != tt.want {
				t.Errorf("%s: got %s, want %s", tt.name, tt.got, tt.want)
			}
		})
	}
}
