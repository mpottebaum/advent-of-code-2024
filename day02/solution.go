package day02

import (
	"aoc/utils"
	"fmt"
	"math"
	"strings"
)

func Solve(inputFile string) {
	fileName := "input"
	if len(inputFile) > 0 {
		fileName = inputFile
	}
	input := utils.ReadFileToString("day02/" + fileName + ".txt")
	rows := strings.Split(input, "\n")
	numSafe := 0
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		numStrs := strings.Split(row, " ")
		var nums []int
		for x := 0; x < len(numStrs); x++ {
			numStr := numStrs[x]
			num, err := utils.ParseInt(numStr)
			if err != nil {
				fmt.Println("parse err", err)
			}
			nums = append(nums, num)
		}
		var isIncreasing bool
		for x := 0; x < len(nums); x++ {
			if len(nums) < 2 {
				break
			}
			if x == len(nums)-1 {
				numSafe += 1
			} else {
				numA := nums[x]
				numB := nums[x+1]
				diff := numB - numA
				if diff == 0 {
					break
				}
				if math.Abs(float64(diff)) > 3 {
					break
				}
				if x == 0 {
					isIncreasing = diff > 0
				} else if isIncreasing && diff < 0 {
					break
				} else if !isIncreasing && diff > 0 {
					break
				}
			}
		}
	}
	fmt.Println("number of safe reports", numSafe)
}
