package main

import (
	"testing"
)

func TestMul(t *testing.T) {
	tables := []TestProvider{
		{2, 2, 4},
		{2, 3, 6},
		{2, 4, 8},
		{2, 5, 10},
	}
	for _, table := range tables {
		result := Mul(table.A, table.B)
		if result != table.Expected {
			t.Errorf(
				"Mul(%d, %d). Got: `%d`. Want: `%d`.",
				table.A,
				table.B,
				result,
				table.Expected,
			)
		}
	}
}
