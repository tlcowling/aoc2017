package main

import (
	"testing"
)

func TestOne(t *testing.T)  {
	cases := map[string]int{
		"1122": 3,
		"1111": 4,
		"1234": 0,
		"91212129": 9,
	}

	for input, output := range cases {
		ans := one(input)
		if ans != output {
			t.Error(input, "should give", output, "but got", ans)
		}
	}
}

func TestOne2(t *testing.T)  {
	cases := map[string]int{
		"1212": 6,
		"1221": 0,
		"123425": 4,
		"123123": 12,
		"12131415": 4,
	}

	for input, output := range cases {
		ans := one2(input)
		if ans != output {
			t.Error(input, "should give", output, "but got", ans)
		}
	}
}
