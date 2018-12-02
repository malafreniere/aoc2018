package first

import (
	"testing"

	"github.com/malafreniere/aoe2018/2/inventory"
)

func TestExecute(t *testing.T) {
	tables := []struct {
		ids      string
		sep      string
		expected int
	}{
		{"abcdef bababc abbcde abcccd aabcdd abcdee ababab", " ", 12},
	}

	for ti, table := range tables {
		inv := inventory.ParseString(table.ids, table.sep)
		checksum := Execute(inv)

		if checksum != table.expected {
			t.Errorf("(%d) Execute returned checksum %d, was expecting %d", ti, checksum, table.expected)
		}
	}
}
