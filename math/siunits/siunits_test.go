package siunits

import "testing"

var si = SI

func Test_siMap_Symbol(t *testing.T) {
	var siTests = []struct {
		name    string
		key     int
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"3", 3, "k", false},
		{"9", 9, "G", false},
		{"-6", -6, "µ", false},
		{"0", 0, "", false},
		{"-7", -7, "", false},
	}

	for _, tt := range siTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := si.Symbol(tt.key); got != tt.want {
				t.Errorf("siMap.Symbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_siMap_Name(t *testing.T) {
	var siTests = []struct {
		name    string
		key     int
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"3", 3, "kilo", false},
		{"9", 9, "giga", false},
		{"-6", -6, "micro", false},
		{"0", 0, "", false},
		{"-7", -7, "", false},
	}

	for _, tt := range siTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := si.Name(tt.key); got != tt.want {
				t.Errorf("siMap.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
