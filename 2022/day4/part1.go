package day4

import (
	"AoC22/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Day4(part int) int {
	if !(part == 1 || part == 2) {
		log.Fatal("Wrong Arguments")
	}
	conditionFunc := getConditionFunc(part)

	scanner := utils.ReadDataLineByLine("2022/day4/input")
	nOverlap := 0
	for scanner.Scan() {
		line := scanner.Text()
		assignments := strings.Split(line, ",")
		ass1 := readAssignment(assignments[0])
		ass2 := readAssignment(assignments[1])
		overlap := conditionFunc(ass1, ass2)
		if overlap {
			nOverlap += 1
		} else {
			fmt.Println(line, overlap)
		}
	}
	return nOverlap
}

func getConditionFunc(part int) func(utils.Interval, utils.Interval) bool {
	if part == 1 {
		return part1Condition
	} else {
		return part2Condition
	}
}

func part1Condition(ass1 utils.Interval, ass2 utils.Interval) bool {
	return ass1.Contains(ass2) || ass2.Contains(ass1)
}

func part2Condition(ass1 utils.Interval, ass2 utils.Interval) bool {
	return ass1.Overlaps(ass2)
}

func readAssignment(s string) utils.Interval {
	assignment := strings.Split(s, "-")
	lower, err := strconv.Atoi(assignment[0])
	if err != nil {
		log.Fatal("Error reading lower assignment bound")
	}
	upper, err := strconv.Atoi(assignment[1])
	if err != nil {
		log.Fatal("Error reading upper assignment bound")
	}
	return utils.Interval{Lower: lower, Upper: upper}
}
