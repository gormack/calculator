package main

import (
	"testing"
)

func TestDiv(t *testing.T) {
	tables := []TestProviderFloat{
		{12, 10, 1.2},
		{12, 6, 2.0},
		{12, 4, 3.0},
		{12, 3, 4.0},
	}
	for _, table := range tables {
		result := Div(table.A, table.B)
		if result != table.Expected {
			t.Errorf(
				"Div(%d, %d). Got: `%f`. Want: `%f`.",
				table.A,
				table.B,
				result,
				table.Expected,
			)
		}
	}
}
