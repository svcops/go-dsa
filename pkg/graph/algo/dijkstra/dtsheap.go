package dijkstra

//
//  distanceToSrc
//  @Description: 临时的数据结构，临时记录点到源的最小距离
//  @Description: 放到最小堆里，贪心的思想
//
type distanceToSrc struct {
	vertex   string
	distance float64
}

// type dtsMinHeap []distanceToSrc
// 定义一个类型
type dtsMinHeap []distanceToSrc

// Len 绑定len方法,返回长度
func (h dtsMinHeap) Len() int {
	return len(h)
}

// Less 绑定less方法
// 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
func (h dtsMinHeap) Less(i, j int) bool {
	// return h[i] < h[j]
	return h[i].distance < h[j].distance
}

// Swap 绑定swap方法，交换两个元素位置
func (h dtsMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Pop 绑定pop方法，从最后拿出一个元素并返回
func (h *dtsMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Push 绑定push方法，插入新元素
func (h *dtsMinHeap) Push(x any) {
	*h = append(*h, x.(distanceToSrc))
}
