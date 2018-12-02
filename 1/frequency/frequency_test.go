package frequency

import (
	"testing"
)

func TestParseString(t *testing.T) {
	type r struct {
		delta int
		op    rune
	}

	tables := []struct {
		input     string
		separator string
		changes   []r
	}{
		{"+1 +1 +1", " ", []r{{1, '+'}, {1, '+'}, {1, '+'}}},
		{"+1 +1 -2", " ", []r{{1, '+'}, {1, '+'}, {2, '-'}}},
		{"-1 -2 -3", " ", []r{{1, '-'}, {2, '-'}, {3, '-'}}},
		{"+1 -2 +3 +1", " ", []r{{1, '+'}, {2, '-'}, {3, '+'}, {1, '+'}}},
	}

	for ti, table := range tables {
		log := ParseString(table.input, table.separator)

		for i, c := range log.Changes() {
			expected := table.changes[i]

			if expected.delta != c.delta {
				t.Errorf("(%d) delta was invalid at index %d. Expecting %d but had %d", ti, i, expected.delta, c.delta)
			}

			if expected.op != c.operator {
				t.Errorf("(%d) operator was invalid at index %d. Expecting %q but had %q", ti, i, expected.op, c.operator)
			}
		}
	}
}

func TestApply(t *testing.T) {
	tables := []struct {
		initial  int
		operator rune
		delta    int
		expected int
	}{
		{1, '+', 0, 1},
		{1, '-', 2, -1},
		{-1, '+', 3, 2},
		{2, '+', 1, 3},
		{1, '+', 1, 2},
		{2, '+', 1, 3},
		{2, '-', 2, 0},
		{-1, '-', 2, -3},
		{-3, '-', 3, -6},
	}

	for ti, table := range tables {
		f := FrequencyChange{
			delta:    table.delta,
			operator: table.operator,
		}

		actual := f.Apply(table.initial)

		if table.expected != actual {
			t.Errorf("(%d) Apply returned %d %c %d = %d, was expecting %d", ti, table.initial, table.operator, table.delta, actual, table.expected)
		}
	}
}
