package day01

import (
	"aoc/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

func Solve(inputFile string) {
	fileName := "input"
	if len(inputFile) > 0 {
		fileName = inputFile
	}
	input := utils.ReadFileToString("day01/" + fileName + ".txt")
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
	sort.Ints(listA)
	sort.Ints(listB)
	var sum int
	for i := 0; i < len(listA); i++ {
		intA := listA[i]
		intB := listB[i]
		distance := math.Abs(float64(intA - intB))
		sum += int(distance)
	}
	fmt.Println("sum of distances", sum)
}
