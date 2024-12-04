package day04

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func Solve(inputFile string) {
	input := utils.ReadFileToString("day04/" + inputFile + ".txt")
	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	xmass := 0
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		walkerSplit := strings.Split(row, "")
		for x := 0; x < len(walkerSplit); x++ {
			letter := row[x]
			if x < len(row)-2 && i < len(rows)-2 {
				nextRow := rows[i+1]
				nextNextRow := rows[i+2]
				if string(letter) == "M" {
					if string(row[x+2]) == "S" && string(nextRow[x+1]) == "A" && string(nextNextRow[x]) == "M" && string(nextNextRow[x+2]) == "S" {
						xmass += 1
					}
					if string(row[x+2]) == "M" && string(nextRow[x+1]) == "A" && string(nextNextRow[x]) == "S" && string(nextNextRow[x+2]) == "S" {
						xmass += 1
					}
				}
				if string(letter) == "S" {
					if string(row[x+2]) == "S" && string(nextRow[x+1]) == "A" && string(nextNextRow[x]) == "M" && string(nextNextRow[x+2]) == "M" {
						xmass += 1
					}
					if string(row[x+2]) == "M" && string(nextRow[x+1]) == "A" && string(nextNextRow[x]) == "S" && string(nextNextRow[x+2]) == "M" {
						xmass += 1
					}
				}
			}
		}
	}
	fmt.Println("xmass", xmass)
}
