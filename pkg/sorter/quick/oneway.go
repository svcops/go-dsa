package quick

import (
	"go-dsa/pkg/sorter"
)

func SortOneWay[T sorter.Iterm](arr []T) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	// [l ... r]
	sortOneWay(arr, 0, len(arr)-1)
}

// sortOneWay [l ... r] 闭区间
func sortOneWay[T sorter.Iterm](arr []T, l, r int) {
	if l >= r {
		return
	}
	mid := partitionOneWay(arr, l, r)
	sortOneWay(arr, l, mid-1)
	sortOneWay(arr, mid+1, r)
}

// partitionOneWay 对arr[l....r]  闭区间的部分进行 partition 操作
func partitionOneWay[T sorter.Iterm](arr []T, l, r int) int {
	v := arr[l]
	// 对 [l+1 ... R] 的数据进行比较
	// 初始化一个左边的索引,默认l
	indexOfLeftMin := l
	for i := l + 1; i <= r; i++ {
		if arr[i] >= v {
			continue
		} else {
			// 要放到左边
			sorter.Swap(arr, indexOfLeftMin+1, i)
			indexOfLeftMin++
		}
	}
	sorter.Swap(arr, l, indexOfLeftMin)
	return indexOfLeftMin
}
