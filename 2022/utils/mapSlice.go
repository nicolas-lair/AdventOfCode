package utils

func SliceMap[T any, M any](mySlice []T, myFunc func(T) M) []M {
	result := make([]M, len(mySlice))
	for i, x := range mySlice {
		result[i] = myFunc(x)
	}
	return result
}
