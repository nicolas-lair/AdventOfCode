package day4

import (
	"AoC22/utils"
	"log"
	"strconv"
	"strings"
)

type Interval struct {
	lower int
	upper int
}

func (i *Interval) contains(other Interval) bool {
	return i.lower <= other.lower && i.upper >= other.upper
}

func Part1() int {
	scanner := utils.ReadDataLineByLine("day4/input")
	nOverlap := 0
	for scanner.Scan() {
		line := scanner.Text()
		assignments := strings.Split(line, ",")
		ass1 := readAssignment(assignments[0])
		ass2 := readAssignment(assignments[1])
		overlap := ass1.contains(ass2) || ass2.contains(ass1)
		if overlap {
			nOverlap += 1
		}
	}
	return nOverlap
}

func readAssignment(s string) Interval {
	assignment := strings.Split(s, "-")
	lower, err := strconv.Atoi(assignment[0])
	if err != nil {
		log.Fatal("Error reading lower assignment bound")
	}
	upper, err := strconv.Atoi(assignment[1])
	if err != nil {
		log.Fatal("Error reading upper assignment bound")
	}
	return Interval{lower: lower, upper: upper}
}
