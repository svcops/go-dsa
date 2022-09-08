// Package dijkstra
// @Description: 递归的方式实现

package dijkstra

import (
	"container/heap"
	"fmt"
)

func (d *dijkstra) CalculateByR(source string) {
	if !d.g.Contains(source) {
		fmt.Printf("源点 ( %s ) 不存在，不计算\n", source)
		return
	}

	d.init(source)

	d.marked[source], d.distToSource[source] = true, 0.0
	// 无所谓的定义
	d.calculateByR(source)
}

func (d *dijkstra) calculateByR(someV string) {
	marked, whereFrom, distToSource, dstHp := d.marked, d.whereFrom, d.distToSource, d.dstHp

	// 所有邻边
	adj := d.g.Adj(someV)

	// 1. 对邻边进行遍历
	for v, weight := range adj {

		// 5. 某个点遍历完了是要标记一下，以后不要再遍历了
		hasMarked := marked[v]
		if hasMarked {
			continue
		}

		// 2. hasCal: 判断领边的另外一个点到 source 的距离是否计算过
		old, hasCal := distToSource[v]

		if hasCal {
			// 3. 计算过
			if distToSource[someV]+weight < old {
				// 重新赋值了距离，如果要找到源的路径，还需要回溯 | 回溯
				distToSource[v], whereFrom[v] = distToSource[someV]+weight, someV
			}

		} else {
			// 3. 没有计算过
			tmpMin := distToSource[someV] + weight
			distToSource[v], whereFrom[v] = tmpMin, someV
		}

		// 6. 被遍历的点，要放到一起，为了找出到源点距离最短的点，进入下一轮循环
		// dstHp.save(v,d.distToSource[v])
		heap.Push(dstHp, distanceToSrc{
			vertex:   v,
			distance: distToSource[v],
		})
	}

	// 4. 某个点遍历完了之后，就不需要在遍历了
	marked[someV] = true

	// 8.容器中没有内容，递归终止
	if len(*dstHp) == 0 {
		return
	}

	// 7. 找出到源点距离最短的点，进入下一轮循环
	// minOfRemaining := someContainer.getMin()
	// calculateByR(d,minOfRemaining)
	minOfRemaining := heap.Pop(dstHp).(distanceToSrc)

	d.calculateByR(minOfRemaining.vertex)
}
