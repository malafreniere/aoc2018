package first

import "github.com/malafreniere/aoe2018/2/inventory"

func Execute(inv *inventory.Inventory) int {
	twoLetters := 0
	threeLetters := 0

	for _, b := range inv.Boxes() {
		letters := countLetters(b.Id)
		hasTwo := false
		hasThree := false
		for _, c := range letters {
			if c == 2 {
				hasTwo = true
			} else if c == 3 {
				hasThree = true
			}
		}

		if hasTwo {
			twoLetters++
		}
		if hasThree {
			threeLetters++
		}
	}

	return twoLetters * threeLetters
}

func countLetters(id string) map[rune]int {
	chars := make(map[rune]int)

	for _, c := range id {
		chars[c] = chars[c] + 1
	}

	return chars
}
