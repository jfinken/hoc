package main

import "testing"

func TestIsIdentChar(t *testing.T) {
	var charTestData = []struct {
		in  rune // careful, casting to rune
		out bool
	}{
		{'_', true},
		{'a', true},
		{'j', true},
		{'⌘', false},
		{6, false},
		{6 - 'a', true},
		{255, true},
	}
	for _, tt := range charTestData {
		got := isIdentChar(tt.in)

		if got != tt.out {
			t.Errorf("isIdentChar err for %q: Got %t, Exp: %t\n", tt.in, got, tt.out)
		}
	}
}

func TestIsSpace(t *testing.T) {
	var charTestData = []struct {
		in  rune // careful, casting to rune
		out bool
	}{
		{'_', false},
		{'a', false},
		{' ', true},
		{'\t', true},
		{'\n', false}, // handled by the grammar
		{255, false},
	}
	for _, tt := range charTestData {
		got := isSpace(tt.in)

		if got != tt.out {
			t.Errorf("isSpace err for %q: Got %t, Exp: %t\n", tt.in, got, tt.out)
		}
	}
}
