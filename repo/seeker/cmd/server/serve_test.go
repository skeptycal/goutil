package main

import (
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	const s, sep, want = "long_string", "string", 5
	got := strings.Index(s, sep)
	if got != want {
		t.Errorf("Index(%q, %q) = %d, want %d", s, sep, got, want)
	}
}
