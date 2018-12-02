package second

import (
	"testing"

	"github.com/malafreniere/aoe2018/1/frequency"
)

func TestExecute(t *testing.T) {
	tables := []struct {
		changes  string
		expected int
	}{
		{"+1 -2 +3 +1", 2},
		{"+1 -1", 0},
		{"+3 +3 +4 -2 -4", 10},
		{"-6 +3 +8 +5 -6", 5},
		{"+7 +7 -2 -7 -4", 14},
	}

	for ti, table := range tables {
		log := frequency.ParseString(table.changes, " ")

		actual := Execute(log)

		if actual != table.expected {
			t.Errorf("(%d) Execute returned %d, was expecting %d", ti, actual, table.expected)
		}
	}
}
