package main

import (
	"aoc/day01"
	"aoc/day02"
	"aoc/day03"
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
		case "2":
			day02.Solve(inputFile)
		case "3":
			day03.Solve(inputFile)
		}
	}
}
