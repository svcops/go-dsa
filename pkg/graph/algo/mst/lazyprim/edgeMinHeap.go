package lazyprim

import (
	"go-dsa/pkg/graph"
)

type edgeMinHeap []graph.Edge

func (eh edgeMinHeap) Len() int {
	return len(eh)
}

func (eh edgeMinHeap) Less(i, j int) bool {
	return eh[i].Weight < eh[j].Weight
}

func (eh edgeMinHeap) Swap(i, j int) {
	eh[i], eh[j] = eh[j], eh[i]
}

func (eh *edgeMinHeap) Push(x any) {
	*eh = append(*eh, x.(graph.Edge))
}

func (eh *edgeMinHeap) Pop() any {
	old := *eh
	n := len(old)
	x := old[n-1]
	*eh = old[0 : n-1]
	return x
}
