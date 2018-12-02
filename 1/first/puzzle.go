package first

import (
	"github.com/malafreniere/aoc2018/1/frequency"
)

func Execute(log *frequency.FrequencyChangeLog) int {
	freq := 0

	for _, change := range log.Changes() {
		freq = change.Apply(freq)
	}

	return freq
}
