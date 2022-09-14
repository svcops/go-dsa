package sparse

import (
	"fmt"
	"go-ads/pkg/graph"
)

//
//  implGraph
//  @Description: 稀疏图的实现
//
type implGraph struct {
	direct    bool
	g         map[string]map[string]float64
	cover     bool
	allVertex graph.Set
}

//
// CreateSparseGraph
//  @Description:
//  @param direct
//  @param cover 添加了重复的边，是否选择覆盖
//  @return *implGraph
//
func CreateSparseGraph(direct, cover bool) *implGraph {
	return &implGraph{
		direct:    direct,
		g:         make(map[string]map[string]float64),
		cover:     cover,
		allVertex: graph.CreateSet(),
	}
}

func (sp *implGraph) IsDirect() bool {
	return sp.direct
}

func (sp *implGraph) EdgesNum() int {
	// TODO implement me
	size := 0
	for _, es := range sp.g {
		size += len(es)
	}
	return size
}

func (sp *implGraph) VerticesNum() int {
	return len(sp.g)
}

func (sp *implGraph) Vertices() graph.Set {
	return sp.allVertex
}

func (sp *implGraph) Contains(v string) bool {
	return sp.allVertex.Contains(v)
}

func (sp *implGraph) Connect(from, to string, weight float64) {
	if from == to {
		// 忽略自环边
		return
	}

	g := sp.g
	// 连接两点
	fromEdges := g[from]

	if fromEdges == nil {
		fromEdges = make(map[string]float64)
		g[from] = fromEdges
	}

	// 两种情况需要更新边 1. cover==true 覆盖策略为true 2. from->to 的边还没有初始化,即没有包含
	if _, contains := fromEdges[to]; sp.cover || !contains {
		// 直接覆盖原有的边,利用map特性
		fromEdges[to] = weight
	}

	// 无向图
	if !sp.direct {
		toEdges := g[to]
		if toEdges == nil {
			toEdges = make(map[string]float64)
			g[to] = toEdges
		}
		if _, contains := toEdges[from]; sp.cover || !contains {
			toEdges[from] = weight
		}
	}

	sp.allVertex.Add(from)
	sp.allVertex.Add(to)
}

// Adj 邻边迭代器，可以优化
func (sp *implGraph) Adj(v string) map[string]float64 {
	edges, contains := sp.g[v]
	if contains {
		return edges
	}
	// 不存在边，返回空map
	return make(map[string]float64)
}

func (sp *implGraph) HasEdge(from, to string) bool {
	_, has := sp.GetEdgeWeight(from, to)
	return has
}

func (sp *implGraph) GetEdgeWeight(from, to string) (float64, bool) {
	if !sp.Contains(from) || !sp.Contains(to) || from == to {
		// 忽略自环边
		return -1, false
	}
	adj := sp.Adj(from)
	for v, w := range adj {
		if v == to {
			return w, true
		}
	}
	return -1, false
}

func (sp *implGraph) Show() {
	fmt.Printf("SparseGraph info : direct = %t ; cover = %t ; vertexSize = %d ; edgeSize = %d) \n",
		sp.IsDirect(), sp.cover, sp.VerticesNum(), sp.EdgesNum())
	g := sp.g
	for from, edges := range g {
		fmt.Printf("[ from = %s ] : ", from)
		for to, weight := range edges {
			fmt.Printf("[ to = %s weight = %g ]   ", to, weight)
		}
		fmt.Println()
	}
	fmt.Println()
}
