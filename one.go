package main


func one(input string) int {
	sum := 0

	for i := 0; i < len(input); i++ {
		this := input[i] - '0'
		next := input[(i+1) % len(input)] - '0'

		if this == next {
			sum += int(this)
		}
	}

	return sum
}

func one2(input string) int {
	sum := 0

	for i := 0; i < len(input); i++ {
		this := input[i] - '0'
		next := input[(i+len(input)/2) % len(input)] - '0'

		if this == next {
			sum += int(this)
		}
	}

	return sum
}