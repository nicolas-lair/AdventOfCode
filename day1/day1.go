package day1

import (
	"AoC22/utils"
	"log"
	"sort"
	"strconv"
)

func Day1(nMax int) int {
	scanner := utils.ReadDataLineByLine("data/input.txt")

	maxCalorieList := make([]int, nMax+1)
	currentCalorieCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			maxCalorieList[0] = currentCalorieCount
			sort.Ints(maxCalorieList)
			currentCalorieCount = 0
		} else {
			calorie, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			currentCalorieCount += calorie
		}

	}
	return utils.SumInts(maxCalorieList[1 : nMax+1])
}
