package main

import "testing"

func TestPadZeros(t *testing.T) {
	s1 := "1a3"
	s2 := "1a2b3c"

	s1padded := padZeros(s1, 4)
	if len(s1padded) != 4 {
		t.Errorf("Expected length to be 4, but got %v", len(s1padded))
	}
	if s1padded[0] != '0' {
		t.Errorf("Expected first symbol to be 0, but got %v", string(s1padded[0]))
	}
	if s1padded[1:len(s1padded)] != s1 {
		t.Errorf("Expected end symbols to be %v, but got %v", s1, s1padded[1:len(s1padded)])
	}

	s2pad := padZeros(s2, 4)
	if s2 != s2pad {
		t.Errorf("Expected string not to change, but got %v", s2pad)
	}
}
