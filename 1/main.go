package main

import (
	"fmt"

	"github.com/malafreniere/aoe2018/1/first"
	"github.com/malafreniere/aoe2018/1/frequency"
	"github.com/malafreniere/aoe2018/1/second"
)

func main() {
	const fileName = "input.txt"

	log := frequency.ParseFile(fileName)

	fmt.Printf("Puzzle #1: %d\n", first.Execute(log))
	fmt.Printf("Puzzle #2: %d", second.Execute(log))
}
