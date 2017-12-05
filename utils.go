package main

func DivisorSum(line SpreadsheetLine) int {
	sum := 0
	for i := 0; i < len(line); i++ {
		for j := 0; j < len(line); j++ {
			if i == j {
				continue
			}
			if line[i]%line[j] == 0 {
				sum += line[i] / line[j]
			}
		}
	}
	return sum
}

func MinMax(line SpreadsheetLine) (int, int) {
	if len(line) == 0 {
		return 0, 0
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

func AsciiNumForChar(input uint8) uint8 {
	return input - '0'
}
