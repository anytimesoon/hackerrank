package Mini_MaxSum

import (
	"sort"
)

func MiniMaxSum(arr []int32) []int64 {
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	mid := int64(arr[1]) + int64(arr[2]) + int64(arr[3])

	return []int64{mid + int64(arr[0]), mid + int64(arr[4])}
}
