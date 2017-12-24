package main

import (
	"testing"
)

func TestDiscard(t *testing.T) {
	s := [][]int{{0, 2}, {2, 2}, {2, 3}, {3, 4}, {3, 5}, {0, 1}, {10, 1}, {9, 10}}
	p := []int{2, 3}

	sn := discard(s, p)

	if len(sn) != len(s)-1 {
		t.Errorf("Expected length to be smaller by one")
	}

	if sn[2][0] != 3 && sn[2][1] != 4 {
		t.Errorf("Expected third element to be removed")
	}
}
