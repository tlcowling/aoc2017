package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func five() int {
	instructions := []int{}
	file, err := os.Open("inputs/5")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileline := scanner.Text()
		num, _ := strconv.Atoi(fileline)
		instructions = append(instructions, num)
	}
	return iterate(instructions)
}

func fivePartTwo() int {
	instructions := []int{}
	file, err := os.Open("inputs/5")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileline := scanner.Text()
		num, _ := strconv.Atoi(fileline)
		instructions = append(instructions, num)
	}
	return iteratePart2(instructions)
}

func iterate(instructions []int) int {
	count := 0
	position := 0
	for {
		if position < 0 || position >= len(instructions) {
			return count
		}

		diff := instructions[position]
		instructions[position]++
		position = position + diff

		count++
	}
	return count
}

func iteratePart2(instructions []int) int {
	count := 0
	position := 0
	for {
		if position < 0 || position >= len(instructions) {
			return count
		}

		diff := instructions[position]
		if diff >= 3 {
			instructions[position]--
		} else {
			instructions[position]++
		}

		position = position + diff

		count++
	}
	return count
}
