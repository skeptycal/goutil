package find

import (
	"testing"
)

func TestAppArgs(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		wantApp  string
		wantArgs string
	}{
		{"", "the empty string", "", ""},
		{"ls", "ls", "ls", ""},
		{"ls -A", "ls -A", "ls", "ls -A"},
		{"ls -A *.go", "ls -A *.go", "ls", "ls -A *.go"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotApp, gotArgs := AppArgs(tt.in)
			// arglist := strings.Split(tt.wantArgs, " ")

			if gotApp != tt.wantApp {
				t.Errorf("AppArgs(%v) gotApp = %v, want %v", tt.name, gotApp, tt.wantApp)
			}

			// for i, got := range arglist {
			// 	if got != arglist[i] {
			if gotArgs != tt.wantArgs {
				t.Errorf("AppArgs(%v) gotArgs = %v, want %v", tt.name, gotArgs, tt.wantArgs)

			}
		})
	}
}
