package main

import "testing"

func TestTwo(t *testing.T) {
	test := Spreadsheet{
		lines: []SpreadsheetLine{
			SpreadsheetLine{5, 1, 9, 5},
			SpreadsheetLine{7, 5, 3},
			SpreadsheetLine{2, 4, 6, 8},
		},
	}

	ans := test.checksum()
	expected := 18
	if ans != expected {
		t.Error("Checksum should be", expected, "but was", ans)
	}
}

func TestLine(t *testing.T) {
	cases := map[*SpreadsheetLine]int{
		&SpreadsheetLine{5, 9, 2, 8}: 4,
		&SpreadsheetLine{9, 4, 7, 3}: 3,
		&SpreadsheetLine{3, 8, 6, 5}: 2,
	}

	for testInput, testOutput := range cases {
		actual := DivisorSum(*testInput)
		if actual != testOutput {
			t.Error(testInput, "should be", testOutput, "but was", actual)
		}
	}
}

func TestTwo2(t *testing.T) {
	test := Spreadsheet{
		lines: []SpreadsheetLine{
			SpreadsheetLine{5, 9, 2, 8},
			SpreadsheetLine{9, 4, 7, 3},
			SpreadsheetLine{3, 8, 6, 5},
		},
	}

	ans := test.divisorsum()
	expected := 9
	if ans != expected {
		t.Error("evenDivisorSum should be", expected, "but was", ans)
	}
}
