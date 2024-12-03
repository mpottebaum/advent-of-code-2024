package day03

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func Solve(inputFile string) {
	fileName := "input"
	if len(inputFile) > 0 {
		fileName = inputFile
	}
	input := utils.ReadFileToString("day03/" + fileName + ".txt")
	postMulones := strings.Split(input, "mul(")
	postMulones = postMulones[1:]
	muller := 0
	for i := 0; i < len(postMulones); i++ {
		postMulone := postMulones[i]
		closingTime := strings.Split(postMulone, ")")
		pairStr := closingTime[0]
		mulps := strings.Split(pairStr, ",")
		if len(mulps) == 2 {
			mulpa := mulps[0]
			mulpb := mulps[1]
			inta, erra := utils.ParseInt(mulpa)
			intb, errb := utils.ParseInt(mulpb)
			if erra == nil && errb == nil {
				muller += inta * intb
			}
		}
	}
	fmt.Println("muller sum", muller)
}
