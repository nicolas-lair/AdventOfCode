package utils

import "log"

func MaxInt(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxInts(x []int) int {
	if len(x) == 0 {
		log.Fatal("Empty Slice")
	}
	max := x[0]
	for _, item := range x[1:] {
		max = MaxInt(max, item)
	}
	return max
}

func SumInts(x []int) int {
	if len(x) == 0 {
		log.Fatal("Empty Slice")
	}
	if len(x) != 1 {
		for _, e := range x {
			return e + SumInts(x[1:])
		}
	}
	return x[0]
}
