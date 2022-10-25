package citsaae

import "testing"

func TestArrayStringsAreEqual(t *testing.T) {
	if arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}) == false {
		t.Errorf("expected true, got false")
	}
}
