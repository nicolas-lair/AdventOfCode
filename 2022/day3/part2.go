package day3

import (
	"AoC22/utils"
)

func Part2(groupSize int) int {
	scanner := utils.ReadDataLineByLine("day3/input")
	priorityValue := 0

	var newLine []rune
	var commonItem utils.Set

reader:
	for {
		for i := 0; i < groupSize; i++ {
			endOfFile := !scanner.Scan()
			if endOfFile {
				break reader
			}
			newLine = []rune(scanner.Text())
			if i == 0 {
				commonItem = utils.PopulateSet(newLine)
			} else {
				commonItem = *commonItem.Intersection(utils.PopulateSet(newLine))
			}
		}
		elem := utils.SetAsSlice[rune](&commonItem)
		priorityValue += evalLetter(elem[0])
	}
	return priorityValue
}
