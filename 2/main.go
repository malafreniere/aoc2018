package main

import (
	"fmt"

	"github.com/malafreniere/aoe2018/2/first"
	"github.com/malafreniere/aoe2018/2/inventory"
	"github.com/malafreniere/aoe2018/2/second"
)

func main() {
	const fileName = "input.txt"

	log := inventory.ParseFile(fileName)

	fmt.Printf("Puzzle #1: %d\n", first.Execute(log))
	fmt.Printf("Puzzle #2: %v\n", second.Execute(log))
}
