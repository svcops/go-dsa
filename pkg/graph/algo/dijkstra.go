// Package algo
// @Description: 单源最短路径

package algo

import (
	"container/heap"
	"go-ads/pkg/graph"
)

type dijkstra struct {
	g            graph.Graph
	source       string
	marked       map[string]bool
	whereFrom    map[string]string
	distToSource map[string]float64
}

func CreateDijkstraAlgo(g graph.Graph, source string) *dijkstra {
	// 初始化辅助的数据结构
	marked := make(map[string]bool)
	for v, _ := range g.GetAllVertex() {
		marked[v] = false
	}

	// 标记最短距离的点从哪里来的，动态更新
	whereFrom := make(map[string]string)

	// 标记到源的最短距离，动态更新
	distToSource := make(map[string]float64)

	return &dijkstra{
		g:            g,
		source:       source,
		marked:       marked,
		whereFrom:    whereFrom,
		distToSource: distToSource,
	}

}

func (d *dijkstra) Calculate() {
	// 一、 初始化 源 source
	// 标记source式已经访问了的
	d.marked[d.source] = true
	// 标记source 到 source 的距离为 0.0 即最小
	d.distToSource[d.source] = 0.0

	// 最小堆，提升性能
	vdheap := &vdMinHeap{}
	heap.Init(vdheap)
	heap.Push(vdheap, vDistance{
		vertex:   d.source,
		distance: 0,
	})

	for len(*vdheap) > 0 {

		minVd, _ := heap.Pop(vdheap).(vDistance)

		curV := minVd.vertex
		// 从最小堆里取出来，标记一下
		d.marked[curV] = true

		// 对所有的邻边进行遍历
		// curV -> wight -> to 组成一条边
		for to, weight := range d.g.Adj(curV) {
			// fmt.Println(v, distance)
			toVisited := d.marked[to]
			if !toVisited {
				_, ok := d.whereFrom[to]

				if !ok {
					shorter := d.distToSource[curV] + weight

					d.whereFrom[to] = curV
					d.distToSource[to] = shorter

					heap.Push(vdheap, vDistance{
						vertex:   to,
						distance: shorter,
					})
				} else {
					maybeShorter := d.distToSource[curV] + weight
					if maybeShorter < d.distToSource[to] {
						d.whereFrom[to] = curV
						d.distToSource[to] = maybeShorter

						heap.Push(vdheap, vDistance{
							vertex:   to,
							distance: maybeShorter,
						})
					}
				}
			}
		}
	}
}
