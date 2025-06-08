package sort

import (
	"go-dsa/pkg/sorter"
	"go-dsa/pkg/sorter/insert"
	"go-dsa/pkg/sorter/quick"
	"go-dsa/pkg/sorter/select"
	"testing"
)

const size = 100000

const simple = true

func Test_Select(t *testing.T) {
	arr := generateArr(size, simple, t)

	bubble.Sort(arr)

	// 验证
	sorted := sorter.IsSorted(arr)

	if !sorted {
		t.Error("冒泡排序顺序错误")
	}
}

func Test_Insert_1(t *testing.T) {
	arr := generateArr(size, simple, t)

	insert.Sort1(arr)
	sorted := sorter.IsSorted(arr)

	if !sorted {
		t.Error("插入排序方式一排序错误")
	}
}

func Test_Insert_2(t *testing.T) {
	arr := generateArr(size, simple, t)

	insert.Sort2(arr)
	sorted := sorter.IsSorted(arr)
	if !sorted {
		t.Error("插入排序方式二排序错误")
	}
}

func Test_Quick_OneWay(t *testing.T) {
	arr := generateArr(size, simple, t)

	quick.SortOneWay(arr)

	sorted := sorter.IsSorted(arr)
	if !sorted {
		t.Error("单路快排顺序错误")
	}
}

func Test_Quick_TwoWay(t *testing.T) {
	arr := generateArr(size, simple, t)
	cp := arrCopy(arr)

	quick.SortTwoWay(arr)

	// 验证排序性
	sorted := sorter.IsSorted(arr)
	if !sorted {
		t.Error("双路快排顺序错误")
	}
	// 验证正确性
	equals := sorter.Equals(cp, arr)
	if !equals {
		t.Error("双路快排元素错误")
	}
}

func arrCopy(src []int) []int {
	cp := make([]int, len(src))
	copy(cp, src)
	return cp
}
