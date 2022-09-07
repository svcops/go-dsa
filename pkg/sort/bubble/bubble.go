package bubble

import "go-ads/pkg/sort"

func Sort[T sort.Iterm](arr []T) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				sort.Swap(arr, i, j)
			}
		}
	}
}
