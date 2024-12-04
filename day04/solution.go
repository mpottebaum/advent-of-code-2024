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
		forwarderSplit := strings.Split(row, "XMAS")
		forwarder := len(forwarderSplit) - 1
		backwarderSplit := strings.Split(row, "SAMX")
		backwarder := len(backwarderSplit) - 1
		walkerSplit := strings.Split(row, "")
		vertDowners := 0
		vertUppers := 0
		downRighters := 0
		downLefters := 0
		upRighters := 0
		upLefters := 0
		for x := 0; x < len(walkerSplit); x++ {
			letter := row[x]
			if string(letter) == "X" {
				if i < len(rows)-3 {
					nextRow := rows[i+1]
					nextNextRow := rows[i+2]
					nextNextNextRow := rows[i+3]
					if string(nextRow[x]) == "M" && string(nextNextRow[x]) == "A" && string(nextNextNextRow[x]) == "S" {
						vertDowners += 1
					}
					if x < len(row)-3 && string(nextRow[x+1]) == "M" && string(nextNextRow[x+2]) == "A" && string(nextNextNextRow[x+3]) == "S" {
						downRighters += 1
					}
					if x > 2 && string(nextRow[x-1]) == "M" && string(nextNextRow[x-2]) == "A" && string(nextNextNextRow[x-3]) == "S" {
						downLefters += 1
					}
				}
				if i > 2 {
					prevRow := rows[i-1]
					prevPrevRow := rows[i-2]
					prevPrevPrevRow := rows[i-3]
					if string(prevRow[x]) == "M" && string(prevPrevRow[x]) == "A" && string(prevPrevPrevRow[x]) == "S" {
						vertUppers += 1
					}
					if x < len(row)-3 && string(prevRow[x+1]) == "M" && string(prevPrevRow[x+2]) == "A" && string(prevPrevPrevRow[x+3]) == "S" {
						upRighters += 1
					}
					if x > 2 && string(prevRow[x-1]) == "M" && string(prevPrevRow[x-2]) == "A" && string(prevPrevPrevRow[x-3]) == "S" {
						upLefters += 1
					}
				}
			}
		}
		xmass += forwarder + backwarder + vertDowners + vertUppers + downRighters + downLefters + upRighters + upLefters
	}
	fmt.Println("xmass", xmass)
}
