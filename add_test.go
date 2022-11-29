package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tables := []TestProvider{
		{1, 1, 2},
		{1, 2, 3},
		{1, 3, 4},
		{1, -5, -4},
	}
	for _, table := range tables {
		result := Add(table.A, table.B)
		if result != table.Expected {
			t.Errorf(
				"Add(%d, %d). Got: `%d`. Want: `%d`.",
				table.A,
				table.B,
				result,
				table.Expected,
			)
		}
	}
}
