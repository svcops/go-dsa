package fixed

import (
	"fmt"
	"go-ads/pkg/tree/maxheap"
)

type heap struct {
	size     int
	capacity int
	values   []float64
}

// CreateFixedLengthHeap 创建固定大小的堆
func CreateFixedLengthHeap(capacity int) *heap {
	return &heap{
		size:     0,
		capacity: capacity,
		values:   make([]float64, capacity+1),
	}
}

func (h *heap) Size() int {
	return h.size
}

func (h *heap) IsEmpty() bool {
	return h.size == 0
}

func (h *heap) Add(value float64) (bool, error) {
	if h.size == h.capacity {
		return false, maxheap.HeapError(fmt.Sprintf("OutBoundError capacity is %d", h.size))
	}
	h.size++
	h.values[h.size] = value
	h.shiftUpByLoop(h.size)
	return true, nil
}

func (h *heap) shiftUp(cur int) {
	if cur == 1 {
		return
	}
	parent := cur / 2
	if h.values[cur] > h.values[parent] {
		swap(h.values, cur, parent)
		h.shiftUp(parent)
	}
}

func (h *heap) shiftUpByLoop(cur int) {
	i, pi := cur, cur/2
	for i >= 2 && h.values[i] > h.values[pi] {
		swap(h.values, i, pi)
		i, pi = pi, pi/2
	}
}

func (h *heap) Max() (float64, error) {
	if h.Size() == 0 {
		return 0.0, maxheap.HeapError("Size is zero")
	}
	return h.values[1], nil
}

func (h *heap) ExtractMax() (float64, error) {
	if h.size == 0 {
		return 0.0, maxheap.HeapError("Heap is empty")
	}

	if h.size == 1 {
		h.size--
		return h.values[1], nil
	}
	max := h.values[1]
	// 第一个和最后一个交换
	swap(h.values, 1, h.size)
	// 先减少size
	h.size--
	h.shiftDownLoop(1)
	return max, nil
}

func (h *heap) shiftDown(cur int) {
	left, right := cur*2, cur*2+1
	if left <= h.size || right <= h.size {
		// if left == h.size {
		// 	// 没有右子树,也是最后一次比较
		// 	if h.values[left] > h.values[cur] {
		// 		swap(h.values, cur, left)
		// 	}
		// 	return
		// } else {
		// 	// 有右子树,不确定是不是最后一次
		// 	var subMax int
		// 	if h.values[left] > h.values[right] {
		// 		subMax = left
		// 	} else {
		// 		subMax = right
		// 	}
		// 	if h.values[subMax] > h.values[cur] {
		// 		swap(h.values, cur, subMax)
		// 		h.shiftDown(subMax)
		// 	}
		// }
		var choose int
		if left == h.size {
			choose = left
		} else {
			if h.values[left] > h.values[right] {
				choose = left
			} else {
				choose = right
			}
		}

		if h.values[choose] > h.values[cur] {
			swap(h.values, cur, choose)
			h.shiftDown(choose)
		}
	}
}

func (h *heap) shiftDownLoop(cur int) {
	p, left, right := cur, cur*2, cur*2+1
	for left <= h.size || right <= h.size {
		if left == h.size {
			if h.values[left] > h.values[p] {
				swap(h.values, p, left)
			}
			break
		} else {
			var subMax int
			if h.values[left] > h.values[right] {
				subMax = left
			} else {
				subMax = right
			}
			if h.values[subMax] > h.values[p] {
				swap(h.values, p, subMax)
				p, left, right = subMax, subMax*2, subMax*2+1
			}
		}
	}
}

func swap(arr []float64, x, y int) {
	arr[x], arr[y] = arr[y], arr[x]
}
