package main

import "testing"

func TestCycle(t *testing.T) {
	bank := []int{0, 2, 7, 0}
	expected := 5
	actual := cycleCount(bank)
	if actual != expected {
		t.Error("expected", expected, "but got", actual, "for", bank)
	}
}
