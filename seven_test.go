package main

import (
	"log"
	"testing"
)

func TestSeven(t *testing.T) {
	p := parseTowerProgramFromLine("bxlur (38) -> ktlj, gibl")

	if p.name != "bxlur" {
		t.Error("name should be blxur")
	}
	if p.weight != 38 {
		t.Error("weight should be 38")
	}
	log.Println(p.dependents)
}
