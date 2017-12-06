package main

import "log"

func cycleCount(bank []int) int {
	count := 0
	max := 0
	seenBefore := [][]int{}
	log.Println(bank)
	lastLargestIndex := 0

	for i := 0; i < len(bank); i++ {
		if bank[i] > max {
			max, lastLargestIndex = bank[i], i
		}
	}

	bank[lastLargestIndex] = bank[l2

	log.Println("amount:", amountToAdd)
	log.Println("bank now:", bank)

	for i := lastLargestIndex; i < lastLargestIndex+len(bank); i++ {
		j := i % len(bank)
		bank[j] = bank[j] + amountToAdd
	}
	log.Println(bank)

	newState := []int{}
	copy(newState, bank)

	seenBefore = append(seenBefore, newState)
	log.Println(seenBefore)

	return count
}
