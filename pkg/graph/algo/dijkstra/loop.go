// Package dijkstra
// @Description: 循环的方式实现

package dijkstra

import (
	"container/heap"
	"fmt"
)

func (d *dijkstra) CalculateByL(source string) {
	if !d.g.Contains(source) {
		fmt.Printf("源点 ( %s ) 不存在，不计算\n", source)
		return
	}

	d.init(source)

	marked, whereFrom, distToSource, dtsHp := d.marked, d.whereFrom, d.distToSource, d.dstHp
	// 初始化 源 source
	// 标记source式已经访问了的 , 标记source 到 source 的距离为 0.0 即最小
	marked[source], distToSource[source] = true, 0.0

	heap.Push(dtsHp, distanceToSrc{
		vertex:   source,
		distance: 0,
	})

	for len(*dtsHp) > 0 {
		minVd, _ := heap.Pop(dtsHp).(distanceToSrc)

		curV := minVd.vertex
		// 对所有的邻边进行遍历
		// curV -> wight -> to 组成一条边
		for to, weight := range d.g.Adj(curV) {
			// fmt.Println(v, distance)
			toVisited := marked[to]
			if !toVisited {
				_, ok := whereFrom[to]

				if !ok {
					shorter := distToSource[curV] + weight
					whereFrom[to], distToSource[to] = curV, shorter

					heap.Push(dtsHp, distanceToSrc{
						vertex:   to,
						distance: shorter,
					})

				} else {
					maybeShorter := distToSource[curV] + weight
					if maybeShorter < distToSource[to] {
						whereFrom[to], distToSource[to] = curV, maybeShorter

						heap.Push(dtsHp, distanceToSrc{
							vertex:   to,
							distance: maybeShorter,
						})
					}
				}
			}
		}

		// 遍历完成后，标记一下，下次访问到就跳过
		marked[curV] = true
	}
}
