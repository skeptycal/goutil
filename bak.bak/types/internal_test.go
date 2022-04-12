package types

import "testing"

func TestTerminal(t *testing.T) {
	t.Parallel()
	if IsTerminal != isTerminal() {
		t.Errorf("IsTerminal = %v, want %v", IsTerminal, isTerminal())
	}
}

func Test_isTerminal(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		want bool
	}{
		{"IsTerminal", IsTerminal},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTerminal(); got != tt.want {
				t.Errorf("isTerminal() = %v, want %v", got, tt.want)
			}
		})
	}
}
