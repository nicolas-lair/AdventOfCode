package day6

import (
	"AoC22/utils"
)

func Part1() int {
	return FindSolution(4)
}

func Part2() int {
	return FindSolution(14)
}

func FindSolution(markerLength int) int {
	fileContent := utils.ReadEntireFile("2022/day6/input")
	var lastThreeLetters []string
	var startPacket int

	for i, letter := range fileContent {
		if i < markerLength {
			lastThreeLetters = append(lastThreeLetters, string(letter))
		} else {
			lastThreeLetters = append(lastThreeLetters[1:], string(letter))
			if allDifferent(lastThreeLetters) {
				startPacket = i + 1
				break
			}
		}
	}
	return startPacket
}

func allDifferent[T any](reader []T) bool {
	s := new(utils.Set)
	s.Add(reader[0])
	for _, k := range reader[1:] {
		if s.Has(k) {
			return false
		} else {
			s.Add(k)
		}
	}
	return true
}
