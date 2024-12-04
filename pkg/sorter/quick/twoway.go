package quick

import (
	"go-dsa/pkg/sorter"
)

func SortTwoWay[T sorter.Iterm](arr []T) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	// [l ... r]
	sortTwoWay(arr, 0, len(arr)-1)
}

// sortTwoWay 双路快排
func sortTwoWay[T sorter.Iterm](arr []T, l, r int) {
	if l >= r {
		// 说明不能在分了
		return
	}
	mid := partitionTwoWay(arr, l, r)
	sortTwoWay(arr, l, mid-1)
	sortTwoWay(arr, mid+1, r)
}

func partitionTwoWay[T sorter.Iterm](arr []T, l, r int) int {
	v := arr[l]
	i, j := l+1, r
	for true {
		// i 不越界
		for i <= r && arr[i] <= v {
			i++
		}
		for j >= l+1 && arr[j] >= v {
			j--
		}
		if i > j {
			break
		}
		// 交换完成，并进入下一轮循环
		sorter.Swap(arr, i, j)
		i++
		j--
	}
	sorter.Swap(arr, l, j)
	return j
}
