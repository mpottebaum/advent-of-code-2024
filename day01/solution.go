package day01

import "fmt"

func Solve(inputFile string) {
	fileName := "input"
	if len(inputFile) > 0 {
		fileName = inputFile
	}
	fmt.Println("yo world", fileName)
}
