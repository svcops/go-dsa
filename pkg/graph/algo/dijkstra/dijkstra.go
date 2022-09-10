// Package algo
// @Description: 单源最短路径

package dijkstra

import (
	"container/heap"
	"go-ads/pkg/graph"
)

type dijkstra struct {
	g            graph.Graph
	marked       map[string]bool
	whereFrom    map[string]string
	distToSource map[string]float64
	dstHp        *dtsMinHeap
	source       *string
}

func CreateDijkstraAlgo(g graph.Graph) *dijkstra {
	return &dijkstra{
		g:            g,
		marked:       nil,
		whereFrom:    nil,
		distToSource: nil,
		dstHp:        nil,
		source:       nil,
	}
}

//
// init
//  @Description:  初始化辅助的数据结构
//  @receiver d
//
func (d *dijkstra) init(source string) {
	// 标记源头
	marked := make(map[string]bool)
	for v := range d.g.Vertices() {
		marked[v] = false
	}

	// 标记最短距离的点从哪里来的，动态更新
	whereFrom := make(map[string]string)

	// 标记到源的最短距离，动态更新
	distToSource := make(map[string]float64)

	// 初始化辅助的容器
	dstHp := &dtsMinHeap{}
	heap.Init(dstHp) // go 堆的用法

	d.marked, d.whereFrom, d.distToSource, d.dstHp, d.source = marked, whereFrom, distToSource, dstHp, &source
}
