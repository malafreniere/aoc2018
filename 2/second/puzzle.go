package second

import (
	"strings"

	"github.com/malafreniere/aoe2018/2/inventory"
)

func Execute(inv *inventory.Inventory) string {

	for i, b := range inv.Boxes() {
		for _, b2 := range inv.SubBoxes(i + 1) {
			if charDiffCount(b.Id, b2.Id) == 1 {
				return removeDiff(b.Id, b2.Id)
			}
		}
	}

	panic("found none with 1 diff")
}

func charDiffCount(id1, id2 string) int {
	diff := 0
	for i, c := range id1 {
		if c != rune(id2[i]) {
			diff++
		}
	}

	return diff
}

func removeDiff(id1, id2 string) string {
	result := strings.Builder{}
	for i, c := range id1 {
		if c == rune(id2[i]) {
			result.WriteRune(c)
		}
	}
	return result.String()
}
