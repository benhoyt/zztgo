package main

import (
	"testing"
)

func TestUpCase(t *testing.T) {
	tests := []struct {
		in, out byte
	}{
		{0, 0},
		{' ', ' '},
		{'0', '0'},
		{'@', '@'},
		{'A', 'A'},
		{'Z', 'Z'},
		{'`', '`'},
		{'a', 'A'},
		{'z', 'Z'},
		{'{', '{'},
	}
	for _, test := range tests {
		out := UpCase(test.in)
		if out != test.out {
			t.Errorf("UpCase(%q): expected %q, got %q",
				test.in, test.out, out)
		}
	}
}

func TestUpCaseString(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{"", ""},
		{"   ", "   "},
		{" foo ", " FOO "},
		{"Foo BaR", "FOO BAR"},
		{"ABC", "ABC"},
	}
	for _, test := range tests {
		out := UpCaseString(test.in)
		if out != test.out {
			t.Errorf("UpCaseString(%q): expected %q, got %q",
				test.in, test.out, out)
		}
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		in           string
		index, count int16
		out          string
	}{
		{"ABCDEF", 1, 6, "ABCDEF"},
		{"ABCDEF", 0, 6, "ABCDEF"},
		{"ABCDEF", 1, 10, "ABCDEF"},
		{"ABCDEF", 2, 4, "BCDE"},
		{"ABCDEF", 2, 0, ""},
		{"ABCDEF", 3, 1, "C"},
		{"foobar", 2, 5, "oobar"},
	}
	for _, test := range tests {
		out := Copy(test.in, test.index, test.count)
		if out != test.out {
			t.Errorf("Copy(%q, %d, %d): expected %q, got %q",
				test.in, test.index, test.count, test.out, out)
		}
	}
}

func TestPos(t *testing.T) {
	tests := []struct {
		b   byte
		s   string
		pos int16
	}{
		{'.', "xyz", 0},
		{'x', "xyz", 1},
		{'y', "xyz", 2},
		{'z', "xyz", 3},
	}
	for _, test := range tests {
		pos := Pos(test.b, test.s)
		if pos != test.pos {
			t.Errorf("Pos(%q, %q): expected %d, got %d",
				test.b, test.s, test.pos, pos)
		}
	}
}

func TestVal(t *testing.T) {
	tests := []struct {
		s    string
		val  int16
		code int16
	}{
		// Successes
		{"0", 0, 0},
		{"1234", 1234, 0},
		{"  1234", 1234, 0},
		{"-1234", -1234, 0},
		{"+1234", 1234, 0},
		{" -1234", -1234, 0},
		{" +1234", 1234, 0},
		{"32767", 32767, 0},
		{"32768", -32768, 0},
		{"-32768", -32768, 0},
		// Errors
		{"", 0, 1},
		{"-", 0, 2},
		{"123.", 0, 4},
		{"1z", 0, 2},
		{"0x123", 0, 2},
		{" -32768q", 0, 8},
	}
	for _, test := range tests {
		var code int16
		val := Val(test.s, &code)
		if val != test.val || code != test.code {
			t.Errorf("Val(%q): expected val %d code %d, got val %d code %d",
				test.s, test.val, test.code, val, code)
		}
	}
}

func TestStr(t *testing.T) {
	tests := []struct {
		n int16
		s string
	}{
		{0, "0"},
		{123, "123"},
		{-123, "-123"},
	}
	for _, test := range tests {
		s := Str(test.n)
		if s != test.s {
			t.Errorf("Str(%d): expected %q, got %q",
				test.n, test.s, s)
		}
	}

}

func TestStrWidth(t *testing.T) {
	tests := []struct {
		n, width int16
		s        string
	}{
		{0, 5, "    0"},
		{123, 5, "  123"},
		{-123, 5, " -123"},
		{123, 2, "123"},
		{-123, 2, "-123"},
	}
	for _, test := range tests {
		s := StrWidth(test.n, test.width)
		if s != test.s {
			t.Errorf("StrWidth(%d, %d): expected %q, got %q",
				test.n, test.width, test.s, s)
		}
	}

}

func TestDelete(t *testing.T) {
	tests := []struct {
		in           string
		index, count int16
		out          string
	}{
		{"abcdef", 1, 1, "bcdef"},
		{"abcdef", 3, 2, "abef"},
		{"abcdef", 6, 1, "abcde"},
	}
	for _, test := range tests {
		out := Delete(test.in, test.index, test.count)
		if out != test.out {
			t.Errorf("Delete(%q, %d, %d): expected %q, got %q",
				test.in, test.index, test.count, test.out, out)
		}
	}
}

func TestReplace(t *testing.T) {
	tests := []struct {
		s     string
		index int16
		b     byte
		out   string
	}{
		{"foo:bar", 4, '.', "foo.bar"},
		{"foo", 1, 'z', "zoo"},
		{"foo", 3, 'x', "fox"},
	}
	for _, test := range tests {
		out := Replace(test.s, test.index, test.b)
		if out != test.out {
			t.Errorf("Replace(%q, %d, %q): expected %q, got %q",
				test.s, test.index, test.b, test.out, out)
		}
	}
}
