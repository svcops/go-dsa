// Package insert
// @Description: 插入排序

package insert

import (
	"go-ads/pkg/sorter"
)

// Sort1 插入排序的第一种实现
func Sort1[T sorter.Iterm](arr []T) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	// i从第二个元素开始，贪心的思想
	for i := 1; i < len(arr); i++ {
		// 寻找元素arr[i]合适的插入位置
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				sorter.Swap(arr, j, j-1)
			}
		}
	}
}

// Sort2 插入排序的优化
func Sort2[T sorter.Iterm](arr []T) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	// 交换操作编程赋值操作
	for i := 1; i < len(arr); i++ {
		e := arr[i]
		var j int
		for j = i; j >= 1; j-- {
			if arr[j-1] >= e {
				arr[j] = arr[j-1]
			} else {
				break
			}
		}
		arr[j] = e
	}
}

/*
insert value is 3

1 2 4 5 [3]

1 2 4 5 5

3 < 4

1 2 4 4 5

3 >= 2

1 2 [4] 4 5  -> 1 2 [3] 4 5

*/
