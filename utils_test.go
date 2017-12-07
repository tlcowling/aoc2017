package main

import (
	"testing"
)

func TestArrayEquality(t *testing.T) {
	type equalArrayTestInput struct {
		A []int
		B []int
	}

	tests := map[*equalArrayTestInput]bool{
		{A: []int{}, B: []int{}}:               true,
		{A: []int{}, B: []int{1, 2, 3}}:        false,
		{A: []int{3, 4, 5}, B: []int{}}:        false,
		{A: []int{1, 2, 3}, B: []int{1, 2, 3}}: true,
	}

	for input, expected := range tests {
		actual := EqualArrays(input.A, input.B)
		if expected != actual {
			t.Error("Expected", expected, "but got", actual, "for", input)
		}
	}
}

// func TestArrayAddition(t *testing.T) {
// 	type equalArrayTestInput struct {
// 		A []int
// 		B []int
// 	}

// 	tests := map[*equalArrayTestInput][]int{
// 		{A: []int{3, 4, 5}, B: []int{}}:        []int{},
// 		{A: []int{1, 2, 3}, B: []int{1, 2, 3}}: []int{2, 4, 6},
// 	}

// 	for input, expected := range tests {
// 		actual, _ := AddArrays(input.A, input.B)
// 		if EqualArrays(actual, expected) {
// 			t.Error("Expected", expected, "but got", actual, "for", input)
// 		}
// 	}
// }
