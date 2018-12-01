package first

import (
	"fmt"

	"github.com/malafreniere/aoe2018/1/frequency"
)

func Execute(log *frequency.FrequencyChangeLog) {
	freq := 0

	for _, change := range log.Changes() {
		freq = change.Apply(freq)
	}

	fmt.Printf("Puzzle 1: %d\n", freq)
}
