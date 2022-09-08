package graph

import (
	"container/heap"
	"fmt"
	intMinHeap "go-ads/pkg/graph/intheap"
	"testing"
)

func TestIntHeap(t *testing.T) {
	h := &intMinHeap.IntMinHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0} // 创建slice
	heap.Init(h)                                              // 初始化heap
	fmt.Println(*h)
	fmt.Println(heap.Pop(h)) // 调用pop

	push(h)

	fmt.Println(*h)
	for len(*h) > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()
}

func push(h *intMinHeap.IntMinHeap) {
	heap.Push(h, 6) // 调用push
}
