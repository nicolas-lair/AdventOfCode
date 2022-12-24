package main

import (
	"AoC22/day1"
	"AoC22/day2"
	"AoC22/day3"
	"AoC22/day4"
	"AoC22/day5"
	"fmt"
)

func main() {
	// Day 1
	fmt.Printf("Day 1, part 1 : %v\n", day1.Day1(1))
	fmt.Printf("Day 1, part 2 : %v\n", day1.Day1(3))

	// Day 2
	fmt.Printf("Day 2, part 1 : %v\n", day2.Part1())
	fmt.Printf("Day 2, part 2 : %v\n", day2.Part2())

	// Day 3
	fmt.Printf("Day 3, part 1 : %v\n", day3.Part1())
	fmt.Printf("Day 3, part 2 : %v\n", day3.Part2(3))

	// Day 4
	fmt.Printf("Day 4, part 1 : %v\n", day4.Day4(1))
	fmt.Printf("Day 4, part 2 : %v\n", day4.Day4(2))

	// Day 5
	fmt.Printf("Day 5, part 1 : %v\n", day5.Day5(9000))
	fmt.Printf("Day 5, part 2 : %v\n", day5.Day5(9001))

}
