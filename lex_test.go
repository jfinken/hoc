package main

import "testing"

// Tables for the tests
var eqnData = []struct {
	in  string
	out rune
}{
	{"", eof},
	{"+", '+'},
	{"/", '/'},
	{"-", '-'},
	{"^", '^'},
	{"$$", '$'},
	{"2 * 3", '2'},
	{"a = 45.67", 'a'},
	{"j = 12", 'j'},
	{"XyZ = 12", 'X'},
	{"-2", '-'},
}
var charData = []struct {
	in    string
	out   rune
	sidx  int
	width int
}{
	{"", eof, 0, 0},
	{"+", '+', 1, 1},
	{"/", '/', 1, 1},
	{"-", '-', 1, 1},
	{"^", '^', 1, 1},
	{"$$", '$', 1, 1},
	{"2", '2', 1, 1},
	{"a", 'a', 1, 1},
	{"X", 'X', 1, 1},
	{"-2", '-', 1, 1},
	{"3.14", '3', 1, 1},
	{"⌘", '⌘', 3, 3},
	{"本", '本', 3, 3},
}

func TestRead(t *testing.T) {
	for _, td := range charData {
		lxr := &HocLex{src: td.in}
		gotc := lxr.read()

		if gotc != td.out {
			t.Errorf("lexer.read err for %s: Got %q, Exp: %q\n", td.in, gotc, td.out)
		}
		if (lxr.sidx != lxr.width) || lxr.width != td.width {
			t.Errorf("lexer.read err for %s: lxr.width: %d, lxr.sidx: %d, exp: %d\n",
				td.in, lxr.width, lxr.sidx, td.width)
		}
	}
}
func TestBackup(t *testing.T) {
	for _, td := range charData {
		lxr := &HocLex{src: td.in}

		_ = lxr.read()
		lxr.backup()

		if lxr.sidx != (td.sidx - lxr.width) {
			t.Errorf("lexer.backup err for %s: lxr.width: %d, lxr.sidx: %d, exp: %d\n",
				td.in, lxr.width, lxr.sidx, (td.sidx - lxr.width))
		}
	}
}
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
