package lazyprim

import (
	"container/heap"
	"go-ads/pkg/graph"
)

type lazyPrim struct {
	g      graph.Graph
	marked map[string]bool
	emh    *edgeMinHeap
}

func CreateLazyPrimAlgo(g graph.Graph) *lazyPrim {
	return &lazyPrim{
		g:      g,
		marked: nil,
	}
}

func (lp *lazyPrim) Calculate() []graph.Edge {
	if lp.g.IsDirect() {
		// 一般情况下，只对联通的无向图做最小生成树的计算
		return make([]graph.Edge, 0)
	}

	one, ok := lp.g.Vertices().GetRandomOne()
	if !ok {
		// 没有顶点
		return make([]graph.Edge, 0)
	}

	lp.init()
	lp.visit(one)
	minSpanTree, index := make([]graph.Edge, len(lp.g.Vertices())-1), 0

	// 最小堆中的元素不为空
	for len(*lp.emh) > 0 {
		minEdge := heap.Pop(lp.emh).(graph.Edge)
		// 如果两个的访问标记是一样的，说明是切分的同一边，这条边需要了
		if lp.marked[minEdge.From] == lp.marked[minEdge.To] {
			continue
		}
		minSpanTree[index] = minEdge
		index++

		if !lp.marked[minEdge.From] {
			lp.visit(minEdge.From)
		} else {
			lp.visit(minEdge.To)
		}
	}

	return minSpanTree
}

func (lp *lazyPrim) visit(from string) {
	lp.marked[from] = true
	for to, weight := range lp.g.Adj(from) {
		// 判断还没有被切分
		if !lp.marked[to] {
			heap.Push(lp.emh, graph.Edge{From: from, To: to, Weight: weight})
		}
	}
}

// init
//
//	@Description: 初始化
//	@receiver lp
func (lp *lazyPrim) init() {
	visited := make(map[string]bool)
	for v := range lp.g.Vertices() {
		visited[v] = false
	}
	lp.marked = visited

	// 初始化堆
	emh := &edgeMinHeap{}
	heap.Init(emh)
	lp.emh = emh
}
