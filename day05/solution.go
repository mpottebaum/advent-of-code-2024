package day05

import (
	"aoc/utils"
	"fmt"
	"math"
	"strings"
)

func isCorrect(rules []string, update []string) (bool, [][]string) {
	isCorrectlyOrdered := true
	relevantRules := [][]string{}
	for x := 0; x < len(rules); x++ {
		rule := strings.Split(rules[x], "|")
		first := rule[0]
		second := rule[1]
		matchesFirst := false
		matchesSecond := false
		for y := 0; y < len(update); y++ {
			num := update[y]
			if num == first && !matchesSecond {
				matchesFirst = true
			} else if num == first && matchesSecond {
				isCorrectlyOrdered = false
				relevantRules = append(relevantRules, rule)
				break
			} else if num == second && !matchesFirst {
				matchesSecond = true
			} else if num == second && matchesFirst {
				relevantRules = append(relevantRules, rule)
				break
			}
		}
	}
	return isCorrectlyOrdered, relevantRules
}

func shuffler(rules [][]string, update []string) []string {
	for x := 0; x < len(rules); x++ {
		rule := rules[x]
		bro := rule[0]
		ker := rule[1]
		var broIndie, kerIndie int
		for y := 0; y < len(update); y++ {
			num := update[y]
			if num == bro {
				broIndie = y
			} else if num == ker {
				kerIndie = y
			}
		}
		if broIndie > kerIndie {
			shouldaBeenFirst := update[broIndie]
			shouldaBeenSecond := update[kerIndie]
			update[kerIndie] = shouldaBeenFirst
			update[broIndie] = shouldaBeenSecond
		}
	}
	return update
}

func Solve(inputFile string) {
	input := utils.ReadFileToString("day05/" + inputFile + ".txt")
	splitter := strings.Split(input, "\n\n")
	rulesStr := splitter[0]
	rules := strings.Split(rulesStr, "\n")
	updatesStr := splitter[1]
	updates := strings.Split(updatesStr, "\n")
	updates = updates[:len(updates)-1]
	middler := 0
	for i := 0; i < len(updates); i++ {
		update := strings.Split(updates[i], ",")
		isCorrectlyOrdered, relevantRules := isCorrect(rules, update)
		if !isCorrectlyOrdered {
			isChangeCorrect := false
			for !isChangeCorrect {
				update = shuffler(relevantRules, update)
				isChangeCorrect, _ = isCorrect(rules, update)
			}
			middleIndie := int(math.Floor(float64(len(update) / 2)))
			middleManStr := update[middleIndie]
			if middleMan, err := utils.ParseInt(middleManStr); err == nil {
				middler += middleMan
			}
		}
	}
	fmt.Println("middler", middler)
}
