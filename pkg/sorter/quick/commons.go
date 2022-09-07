package quick

import "go-ads/pkg/sorter"

// sortOneWay [l ... r] 闭区间
func quickSort[T sorter.Iterm](arr []T, l, r int,
	partition func([]T, int, int) int) {
	if l >= r {
		return
	}
	mid := partition(arr, l, r)
	quickSort(arr, l, mid-1, partition)
	quickSort(arr, mid+1, r, partition)
}
