package main

import (
	"testing"
)

func TestSub(t *testing.T) {
	tables := []TestProvider{
		{3, 1, 2},
		{3, 2, 1},
		{3, 3, 0},
		{3, 5, -2},
	}
	for _, table := range tables {
		result := Sub(table.A, table.B)
		if result != table.Expected {
			t.Errorf(
				"Sub(%d, %d). Got: `%d`. Want: `%d`.",
				table.A,
				table.B,
				result,
				table.Expected,
			)
		}
	}
}
