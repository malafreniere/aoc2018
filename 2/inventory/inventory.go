package inventory

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Inventory struct {
	boxes []Box
}

type Box struct {
	Id string
}

const NewLineSeparator = "\r\n"

func (i *Inventory) Boxes() []Box {
	return i.boxes[:]
}

func (i *Inventory) SubBoxes(start int) []Box {
	return i.boxes[start:]
}

func ParseFile(fileName string) *Inventory {
	f, e := os.Open(fileName)

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

func ParseString(value string, separator string) *Inventory {
	inv := &Inventory{boxes: make([]Box, 0, 100)}

	for _, line := range strings.Split(value, separator) {
		inv.boxes = append(inv.boxes, Box{
			Id: line,
		})
	}

	return inv
}
