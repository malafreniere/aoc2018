package second

import (
	"github.com/malafreniere/aoc2018/1/frequency"
)

func Execute(log *frequency.FrequencyChangeLog) int {
	freq := 0

	frequencies := make(map[int]bool)
	frequencies[freq] = true

	for {
		for _, change := range log.Changes() {
			freq = change.Apply(freq)

			_, ok := frequencies[freq]

			if ok {
				return freq
			}

			frequencies[freq] = true
		}
	}
}
