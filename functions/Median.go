package cc

import "sort"

func Median(arr []float64) float64 {
	sortedArr := make([]float64, len(arr))
	copy(sortedArr, arr)

	sort.Float64s(sortedArr)

	if len(sortedArr)%2 == 0 {
		return (sortedArr[len(sortedArr)/2-1] + sortedArr[len(sortedArr)/2]) / 2
	}
	return sortedArr[len(sortedArr)/2]
}