package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func seven() int {
	dat, err := ioutil.ReadFile("inputs/7")
	if err != nil {
		log.Fatalln("Cant read input 7", err)
	}
	tps := parseTowerPrograms(strings.Split(string(dat), "\n"))
	log.Println(tps)
	return 0
}

type TowerProgram struct {
	name       string
	weight     int
	dependents []string
}

func parseTowerPrograms(lines []string) []TowerProgram {
	tps := make([]TowerProgram, len(lines))
	for _, line := range lines {
		tps = append(tps, parseTowerProgramFromLine(line))
	}
	return tps
}

func parseTowerProgramFromLine(line string) TowerProgram {
	tokens := strings.Split(line, "->")
	nameAndWeight := strings.Split(tokens[0], " ")
	dependents := []string{}
	if len(tokens) > 1 {
		dependents = strings.Split(tokens[1], " , ")
	}

	if len(nameAndWeight) > 1 {
		log.Println(nameAndWeight)
		weightToken := nameAndWeight[1]
		weightToken = strings.Trim(weightToken, "(")
		weightToken = strings.Trim(weightToken, ")")
		weight, err := strconv.Atoi(weightToken)
		if err != nil {
			log.Fatalln("Unable to parse int", weightToken)
			return TowerProgram{}
		}
		return TowerProgram{
			name:       nameAndWeight[0],
			weight:     weight,
			dependents: dependents,
		}
	}
	log.Fatalln("Unable to parse line")
	return TowerProgram{}
}
