// Package dense
// @Description: 稠密图的实现，非标准实现
package dense

import (
	"fmt"
	"go-ads/pkg/graph"
)

//
//  implGraph
//  @Description: vertexCap 节点容量
//
type implGraph struct {
	direct    bool
	g         [][]float64
	cover     bool
	allVertex graph.Set
	edgesNum  int
	vertexCap int
	increment int
	strToI    map[string]int
	iToStr    []string
}

// CreateDenseGraph
//  @Description:
//  @param vertexNum 顶点个数，和稀疏图不同
//  @param direct
//  @param cover  添加了重复的边，是否选择覆盖
//  @return *implGraph
//
func CreateDenseGraph(vertexCap int, direct, cover bool) *implGraph {
	// 初始化二维矩阵
	g := make([][]float64, vertexCap)
	for i := 0; i < vertexCap; i++ {
		g[i] = make([]float64, vertexCap)
	}

	return &implGraph{
		direct:    direct,
		g:         g,
		cover:     cover,
		allVertex: graph.CreateSet(),
		edgesNum:  0,
		vertexCap: vertexCap,
		increment: 0,
		strToI:    make(map[string]int),
		iToStr:    []string{},
	}
}

func (de *implGraph) IsDirect() bool {
	return de.direct
}

func (de *implGraph) EdgesNum() int {
	return de.edgesNum
}

//
// VerticesNum
//  @Description: 返回初始化了的节点的个数
//  @receiver de
//  @return int
//
func (de *implGraph) VerticesNum() int {
	// return de.vertexCap
	// return len(de.strToI)
	return de.increment
}

func (de *implGraph) Vertices() graph.Set {
	return de.allVertex
}

func (de *implGraph) Contains(v string) bool {
	_, ok := de.strToI[v]
	return ok
}

func (de *implGraph) Connect(from, to string, weight float64) {
	if from == to {
		// 忽略自环边
		return
	}

	i, j, get := de.getRelation(from, to)
	if !get {
		// 获取下标索引失败
		return
	}

	ijW := de.g[i][j]
	if ijW == 0 {
		de.g[i][j] = weight
	} else if de.cover {
		// 需要覆盖
		de.g[i][j] = weight
	}

	if !de.cover {
		de.edgesNum++
	}

	// 无向图
	if !de.direct {
		jiW := de.g[j][i]
		if jiW == 0 {
			de.g[j][i] = weight
		} else if de.cover {
			de.g[j][i] = weight
		}
		if !de.cover {
			de.edgesNum++
		}
	}

	de.allVertex.Add(from)
	de.allVertex.Add(to)
}

// 创建string点 和数组下标的映射
func (de *implGraph) getRelation(from, to string) (int, int, bool) {
	i, hasFrom := de.strToI[from]
	j, hasTo := de.strToI[to]

	rLen := len(de.strToI)
	if hasFrom && hasTo {
		return i, j, true
	} else if rLen == de.vertexCap && (!hasFrom || !hasTo) {
		// 已经满了，还有一个点不在关系列表里面
		return -1, -1, false
	} else if hasFrom && !hasTo {
		// map没有满，存在 from 不存在 to
		// 维护关系
		de.strToI[to] = de.increment
		de.iToStr = append(de.iToStr, to)

		de.increment++
		return i, de.strToI[to], true
	} else if hasTo && !hasFrom {
		// map没有满，存在 to 不存在 from,
		// 维护关系
		de.strToI[from] = de.increment
		de.iToStr = append(de.iToStr, from)

		de.increment++
		return de.strToI[from], j, true
	} else if rLen == de.vertexCap-1 {
		// 只剩下一个槽位了！！！ from和to都不存在
		return -1, -1, false
	} else {
		de.strToI[from] = de.increment
		de.iToStr = append(de.iToStr, from)
		de.increment++
		// 还有很多槽位
		de.strToI[to] = de.increment
		de.iToStr = append(de.iToStr, to)
		de.increment++

		return de.strToI[from], de.strToI[to], true
	}
}

func (de *implGraph) Adj(v string) map[string]float64 {
	if !de.Contains(v) {
		return make(map[string]float64)
	}
	adj := make(map[string]float64)

	// vs := de.g[de.strToI[v]]
	// for j, weight := range vs {
	// 	if weight != 0 {
	// 		// 有边
	// 		adj[de.iToStr[j]] = weight
	// 	}
	// }

	i := de.strToI[v]
	for j := 0; j < de.VerticesNum(); j++ {
		weight := de.g[i][j]
		if weight != 0.0 {
			adj[de.iToStr[j]] = weight
		}
	}
	return adj
}

func (de *implGraph) HasEdge(from, to string) bool {
	if !de.Contains(from) || !de.Contains(to) {
		return false
	}
	return de.g[de.strToI[from]][de.strToI[to]] != 0
}

func (de *implGraph) GetEdgeWeight(from, to string) (float64, bool) {
	if !de.Contains(from) || !de.Contains(to) {
		return 0, false
	}
	weight := de.g[de.strToI[from]][de.strToI[to]]

	if weight == 0 {
		return 0, false
	}
	return weight, true
}

func (de *implGraph) Show() {
	fmt.Printf("DenseGraph info : direct = %t ; cover = %t ; vertexsNum = %d ; edgesNum = %d) \n",
		de.IsDirect(), de.cover, de.VerticesNum(), de.EdgesNum())

	fmt.Print("\t\t")

	for i := 0; i < de.vertexCap; i++ {
		if i < len(de.iToStr) {
			fmt.Printf("[%d:%s]\t", i, de.iToStr[i])
		} else {
			fmt.Printf("[%d: ]\t", i)
		}
	}

	fmt.Println()
	for i := 0; i < de.vertexCap; i++ {
		if i < len(de.iToStr) {
			fmt.Printf("[%d:%s]\t", i, de.iToStr[i])
		} else {
			fmt.Printf("[%d: ]\t", i)
		}
		for j := 0; j < de.vertexCap; j++ {
			fmt.Printf("%g\t\t", de.g[i][j])
		}
		fmt.Println()
	}
}
