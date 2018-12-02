package frequency

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const NewLineSeparator = "\r\n"

type FrequencyChange struct {
	operator rune
	delta    int
}

type FrequencyChangeLog struct {
	changes []FrequencyChange
}

func (f FrequencyChange) Apply(frequency int) int {
	switch f.operator {
	case '+':
		return frequency + f.delta
	case '-':
		return frequency - f.delta
	}

	log.Fatal(fmt.Errorf("unknown operator %v", f.operator))

	return -1
}

func (l *FrequencyChangeLog) Changes() []FrequencyChange {
	return l.changes[:]
}

func ParseFile(file string) *FrequencyChangeLog {
	f, e := os.Open(file)

	if e != nil {
		log.Fatal(e)
	}

	bytes, e := ioutil.ReadAll(f)

	if e != nil {
		log.Fatal(e)
	}

	content := fmt.Sprintf("%s", bytes)

	return ParseString(content, NewLineSeparator)
}

func ParseString(value string, separator string) *FrequencyChangeLog {
	log := &FrequencyChangeLog{changes: make([]FrequencyChange, 0, 100)}

	for _, line := range strings.Split(value, separator) {
		freq := parseFrequency(line)
		log.changes = append(log.changes, freq)
	}

	return log
}

func parseFrequency(value string) FrequencyChange {
	delta, err := strconv.Atoi(value[1:])

	if err != nil {
		log.Fatal(err)
	}

	return FrequencyChange{
		operator: rune(value[0]),
		delta:    delta,
	}
}
