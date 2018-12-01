package second

import (
	"fmt"

	"github.com/malafreniere/aoe2018/1/frequency"
)

func Execute(log *frequency.FrequencyChangeLog) {
	frequencies := make(map[int]bool)
	freq := 0

	for {
		for _, change := range log.Changes() {
			freq = change.Apply(freq)

			_, ok := frequencies[freq]

			if ok {
				fmt.Printf("Puzzle 2: %d", freq)
				return
			}

			frequencies[freq] = true
		}
	}
}
