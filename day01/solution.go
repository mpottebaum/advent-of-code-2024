package day01

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func Solve(inputFile string) {
	input := utils.ReadFileToString("day01/" + inputFile + ".txt")
	rows := strings.Split(input, "\n")
	var listA, listB []int
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		splitRow := strings.Split(row, "   ")
		if len(splitRow) == 2 {
			listAInt, err := utils.ParseInt(splitRow[0])
			if err != nil {
				fmt.Println("parse error", err)
			}
			listBInt, err := utils.ParseInt(splitRow[1])
			if err != nil {
				fmt.Println("parse error", err)
			}
			listA = append(listA, listAInt)
			listB = append(listB, listBInt)
		}
	}
	var bMap = map[int]int{}
	for i := 0; i < len(listB); i++ {
		num := listB[i]
		if val, exists := bMap[num]; exists {
			bMap[num] = val + 1
		} else {
			bMap[num] = 1
		}
	}
	similarityScore := 0
	for i := 0; i < len(listA); i++ {
		num := listA[i]
		if appearances, exists := bMap[num]; exists {
			similarityScore += appearances * num
		}
	}
	fmt.Println("similarity score", similarityScore)
}
