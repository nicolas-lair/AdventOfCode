package utils

import "log"

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
