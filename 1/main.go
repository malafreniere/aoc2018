package main

import (
	"github.com/malafreniere/aoe2018/1/first"
	"github.com/malafreniere/aoe2018/1/frequency"
	"github.com/malafreniere/aoe2018/1/second"
)

func main() {
	const fileName = "input.txt"

	log := frequency.ParseFile(fileName)

	first.Execute(log)
	second.Execute(log)
}
