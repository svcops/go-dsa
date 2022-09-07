package bubble

import "go-ads/pkg/sorter"

func Sort[T sorter.Iterm](arr []T) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				sorter.Swap(arr, i, j)
			}
		}
	}
}
