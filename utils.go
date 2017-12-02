package main

func MinMax(line SpreadsheetLine) (int, int) {
	if len(line) == 0 {
		return 0,0
	}
	max := line[0]
	min := line[0]

	for _, value := range line {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}