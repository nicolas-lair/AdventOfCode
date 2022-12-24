package day3

import (
	"AoC22/utils"
	"unicode"
)

func Part1() int {
	scanner := utils.ReadDataLineByLine("day3/input")
	priorityValue := 0
	for scanner.Scan() {
		line := scanner.Text()
		comp1 := utils.PopulateSet([]rune(line[:len(line)/2]))
		comp2 := utils.PopulateSet([]rune(line[len(line)/2:]))
		elem := utils.SetAsSlice[rune](comp1.Intersection(comp2))
		priorityValue += evalLetter(elem[0])
	}
	return priorityValue
}

func evalLetter(s rune) int {
	var value int
	if unicode.IsUpper(s) {
		value = int(s-'A') + 27
	} else {
		value = int(s-'a') + 1
	}
	return value
}
