package main

import (
	"math"
)

func three() float64 {
	input := 361527

	X := float64(0)
	Y := float64(0)
	steps := 0
	directionIndex := 0

	for steps < input-1 {
		length := (directionIndex / 2) + 1
		for i := 0; i < length; i++ {
			if steps == input-1 {
				break
			}
			steps++
			direction := directionIndex % 4
			switch direction {
			case 0:
				X++
			case 1:
				Y++
			case 2:
				X--
			default:
				Y--
			}
		}
		directionIndex++
	}
	return math.Abs(X) + math.Abs(Y)
}

func threePartTwo() float64 {
	// on other computer!
	return 0
}
