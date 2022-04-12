package defaults

import "testing"

func TestIPAddr_String(t *testing.T) {
	tests := []struct {
		name string
		i    IPAddr
		want string
	}{
		{"127.0.0.1", IPAddr{127, 0, 0, 1}, "127.0.0.1"},
		{"8.8.8.8", IPAddr{8, 8, 8, 8}, "8.8.8.8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("IPAddr.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
