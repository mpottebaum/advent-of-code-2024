package day06

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func stringifyOppDir(row int, col int, dir string) string {
	rowStr := strconv.Itoa(row)
	colStr := strconv.Itoa(col)
	return "r" + rowStr + "c" + colStr + dir
}

func isLoopy(startRow int, startCol int, oppRow int, oppCol int, rows []string) bool {
	oppDirSet := map[string]bool{}
	guardRow := startRow
	guardCol := startCol
	isWalking := true
	direction := "^"
	for isWalking {
		switch direction {
		case "^":
			for guardRow -= 1; guardRow >= 0; guardRow-- {
				char := rows[guardRow][guardCol]
				if string(char) == "#" || (guardRow == oppRow && guardCol == oppCol) {
					stringerBell := stringifyOppDir(guardRow, guardCol, direction)
					if exists := oppDirSet[stringerBell]; exists {
						return true
					} else {
						oppDirSet[stringerBell] = true
					}
					direction = ">"
					guardRow += 1
					break
				} else {
					if guardRow == 0 {
						isWalking = false
					}
				}
			}
		case "v":
			for guardRow += 1; guardRow < len(rows); guardRow++ {
				char := rows[guardRow][guardCol]
				if string(char) == "#" || (guardRow == oppRow && guardCol == oppCol) {
					stringerBell := stringifyOppDir(guardRow, guardCol, direction)
					if exists := oppDirSet[stringerBell]; exists {
						return true
					} else {
						oppDirSet[stringerBell] = true
					}
					direction = "<"
					guardRow -= 1
					break
				} else {
					if guardRow == len(rows)-1 {
						isWalking = false
					}
				}
			}
		case ">":
			row := rows[guardRow]
			for guardCol += 1; guardCol < len(row); guardCol++ {
				char := rows[guardRow][guardCol]
				if string(char) == "#" || (guardRow == oppRow && guardCol == oppCol) {
					stringerBell := stringifyOppDir(guardRow, guardCol, direction)
					if exists := oppDirSet[stringerBell]; exists {
						return true
					} else {
						oppDirSet[stringerBell] = true
					}
					direction = "v"
					guardCol -= 1
					break
				} else {
					if guardCol == len(row)-1 {
						isWalking = false
					}
				}
			}
		case "<":
			for guardCol -= 1; guardCol >= 0; guardCol-- {
				char := rows[guardRow][guardCol]
				if string(char) == "#" || (guardRow == oppRow && guardCol == oppCol) {
					stringerBell := stringifyOppDir(guardRow, guardCol, direction)
					if exists := oppDirSet[stringerBell]; exists {
						return true
					} else {
						oppDirSet[stringerBell] = true
					}
					direction = "^"
					guardCol += 1
					break
				} else {
					if guardCol == 0 {
						isWalking = false
					}
				}
			}
		}
	}
	return false
}

func Solve(inputFile string) {
	input := utils.ReadFileToString("day06/" + inputFile + ".txt")
	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	var startRow, startCol int
getGuard:
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		for x := 0; x < len(row); x++ {
			col := row[x]
			if string(col) == "^" {
				startRow = i
				startCol = x
				break getGuard
			}
		}
	}
	numbLoopies := 0
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		for x := 0; x < len(row); x++ {
			char := row[x]
			if string(char) != "#" {
				if loopy := isLoopy(startRow, startCol, i, x, rows); loopy {
					numbLoopies += 1
				}
			}
		}
	}

	fmt.Println("numb loopies", numbLoopies)
}
