package first

import (
	"testing"

	"github.com/malafreniere/aoe2018/1/frequency"
)

func TestExecute(t *testing.T) {
	tables := []struct {
		changes  string
		expected int
	}{
		{"+1 -2 +3 +1", 3},
		{"+1 +1 +1", 3},
		{"+1 +1 -2", 0},
		{"-1 -2 -3", -6},
	}

	for ti, table := range tables {
		log := frequency.ParseString(table.changes, " ")

		actual := Execute(log)

		if actual != table.expected {
			t.Errorf("(%d) Execute returned %d, was expecting %d", ti, actual, table.expected)
		}
	}
}
