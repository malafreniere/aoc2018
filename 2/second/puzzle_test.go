package second

import (
	"testing"

	"github.com/malafreniere/aoe2018/2/inventory"
)

func TestExecute(t *testing.T) {
	tables := []struct {
		ids      string
		sep      string
		expected string
	}{
		{"abcde fghij klmno	pqrst fguij	axcye wvxyz", " ", "fgij"},
	}

	for ti, table := range tables {
		inv := inventory.ParseString(table.ids, table.sep)
		diff := Execute(inv)

		if diff != table.expected {
			t.Errorf("(%d) Execute returned %s, was expecting %s", ti, diff, table.expected)
		}
	}
}
