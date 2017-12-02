package main

import "testing"

func TestTwo(t *testing.T)  {
	test := Spreadsheet{
		lines: []SpreadsheetLine{
			SpreadsheetLine{5,1,9,5},
			SpreadsheetLine{7,5,3},
			SpreadsheetLine{2,4,6,8},
		},
	}

	if test.checksum() != 18 {
		t.Error("Checksum should be 18 but was", test.checksum()  )
	}
}