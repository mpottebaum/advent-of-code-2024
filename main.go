package main

import (
	"aoc/day01"
	"os"
)

func main() {
	if meArgs := os.Args; len(meArgs) >= 2 {
		var inputFile string
		if len(meArgs) >= 3 {
			inputFile = meArgs[2]
		}
		switch day := meArgs[1]; day {
		case "1":
			day01.Solve(inputFile)
		}
	}
}
