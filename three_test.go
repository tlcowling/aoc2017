package main

import "testing"

func TestGenerate(t *testing.T) {

	// 1,9,25,49  8 16 24.  8n+8
	// 1,5,17,37  4 12 20.  4n+8
	// 1,7,21,43  6 14 22.  6n+8
	// 1,3,13,31  2 10 18.  2n+8
}

// func TestSquare1(t *testing.T) {
// 	testCases := map[int][2]int{
// 		1:  [2]int{0, 0},
// 		2:  [2]int{0, 1},
// 		3:  [2]int{1, 1},
// 		4:  [2]int{1, 0},
// 		5:  [2]int{-1, 1},
// 		6:  [2]int{-1, 0},
// 		7:  [2]int{-1, -1},
// 		8:  [2]int{0, -1},
// 		9:  [2]int{1, -1},
// 		10: [2]int{2, -1},
// 		11: [2]int{2, 0},
// 		12: [2]int{2, 1},
// 		13: [2]int{2, 2},
// 		14: [2]int{1, 2},
// 		15: [2]int{0, 2},
// 		16: [2]int{-1, 2},
// 		17: [2]int{-2, 2},
// 		18: [2]int{-2, 1},
// 		19: [2]int{-2, 0},
// 		20: [2]int{-2, -1},
// 		21: [2]int{-2, -2},
// 		22: [2]int{-1, -2},
// 		23: [2]int{0, -2},
// 		24: [2]int{1, -2},
// 		25: [2]int{-2, 2},
// 	}
// 	ig := InfinityGrid{}
// 	for input, expected := range testCases {
// 		x := ig.StepsFromOrigin(input)
// 		if x != expected[0] {
// 			t.Error(input, "Got ", x, "should be", expected[0])
// 		}
// 	}
// }
