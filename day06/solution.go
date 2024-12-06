package day06

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func stringifyPos(row int, col int) string {
	rowStr := strconv.Itoa(row)
	colStr := strconv.Itoa(col)
	return "r" + rowStr + "c" + colStr
}

func Solve(inputFile string) {
	input := utils.ReadFileToString("day06/" + inputFile + ".txt")
	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	var guardRow, guardCol int
getGuard:
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		for x := 0; x < len(row); x++ {
			col := row[x]
			if string(col) == "^" {
				guardRow = i
				guardCol = x
				break getGuard
			}
		}
	}
	startStr := stringifyPos(guardRow, guardCol)
	posSet := map[string]bool{
		startStr: true,
	}
	isWalking := true
	direction := "^"
	for isWalking {
		switch direction {
		case "^":
			for guardRow -= 1; guardRow >= 0; guardRow-- {
				char := rows[guardRow][guardCol]
				if string(char) == "#" {
					direction = ">"
					guardRow += 1
					break
				} else {
					stringerBell := stringifyPos(guardRow, guardCol)
					posSet[stringerBell] = true
					if guardRow == 0 {
						isWalking = false
					}
				}
			}
		case "v":
			for guardRow += 1; guardRow < len(rows); guardRow++ {
				char := rows[guardRow][guardCol]
				if string(char) == "#" {
					direction = "<"
					guardRow -= 1
					break
				} else {
					stringerBell := stringifyPos(guardRow, guardCol)
					posSet[stringerBell] = true
					if guardRow == len(rows)-1 {
						isWalking = false
					}
				}
			}
		case ">":
			row := rows[guardRow]
			for guardCol += 1; guardCol < len(row); guardCol++ {
				char := rows[guardRow][guardCol]
				if string(char) == "#" {
					direction = "v"
					guardCol -= 1
					break
				} else {
					stringerBell := stringifyPos(guardRow, guardCol)
					posSet[stringerBell] = true
					if guardCol == len(row)-1 {
						isWalking = false
					}
				}
			}
		case "<":
			for guardCol -= 1; guardCol >= 0; guardCol-- {
				char := rows[guardRow][guardCol]
				if string(char) == "#" {
					direction = "^"
					guardCol += 1
					break
				} else {
					stringerBell := stringifyPos(guardRow, guardCol)
					posSet[stringerBell] = true
					if guardCol == 0 {
						isWalking = false
					}
				}
			}
		}
	}
	numbPosers := len(posSet)
	fmt.Println("numb posers", numbPosers)
}
