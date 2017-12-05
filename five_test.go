package main

import "testing"

func TestFive(t *testing.T) {
	test := []int{0, 3, 0, 1, -3}
	actual := iterate(test)
	expected := 5
	if actual != expected {
		t.Error(actual, "should be", expected, "for", test)
	}
}

func TestFiveTwo(t *testing.T) {
	test := []int{0, 3, 0, 1, -3}
	actual := iteratePart2(test)
	expected := 10
	if actual != expected {
		t.Error(actual, "should be", expected, "for", test)
	}
}
