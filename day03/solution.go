package day03

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func Solve(inputFile string) {
	input := utils.ReadFileToString("day03/" + inputFile + ".txt")
	doozers := strings.Split(input, "do()")
	muller := 0
	for i := 0; i < len(doozers); i++ {
		doozer := doozers[i]
		donter := strings.Split(doozer, "don't()")
		enabled := donter[0]
		postMulones := strings.Split(enabled, "mul(")
		postMulones = postMulones[1:]
		for x := 0; x < len(postMulones); x++ {
			postMulone := postMulones[x]
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
	}
	fmt.Println("muller sum", muller)
}
