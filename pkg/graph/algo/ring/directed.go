// Package ring
// @Description: 有向图判断环
package ring

import (
	"log"
)

type directedRing struct {
	*ring
	// 求环的时候整个图的状态
	visited map[string]bool
}

const (
	// 当访问节点的邻边的时候
	//  - 如果另外一个节点还没有被访问，继续深度遍历下去
	//  - 如果另外一个节点的所有邻边已经被访问过了，说明是第二次进入这个节点，说明有环

	// 一个节点在深度遍历的过程中的三种状态
	// dfsInit: 可以进行深度遍历
	// dfsStart: 开始深度遍历
	// dfsEnd: 所有的邻边都已经遍历完成
	dfsInit  = 0
	dfsStart = 1
	dfsEnd   = 2
)

func (dr *directedRing) calculate() {
	dr.visited = make(map[string]bool)
	for v := range dr.g.Vertices() {
		dr.visited[v] = false
	}
	// A -> B -> C ->  D
	//      ↑         ↙
	//        ↖ <-  ↙

	// A -> B -> C -> D
	//         ↗
	//        E

	for v := range dr.g.Vertices() {

		if dr.hasRing {
			// 如果有环，就不要再遍历了
			break
		}

		if !dr.visited[v] {
			dfsStatus := make(map[string]int)
			for tmpV := range dr.g.Vertices() {
				dfsStatus[tmpV] = dfsInit
			}
			dr.directedDfs(v, dfsStatus)
		}
	}
}

// 有向图的深度遍历
func (dr *directedRing) directedDfs(from string, dfsStatus map[string]int) {
	if dr.debug {
		log.Printf("开始访问节点 %s", from)
	}
	// 第一种状态，全局看，已经被访问了，下一次的for循环就不会遍历到了
	dr.visited[from] = true
	dfsStatus[from] = dfsStart

	for to := range dr.g.Adj(from) {
		if dfsStatus[to] == dfsInit {
			dr.directedDfs(to, dfsStatus)
		} else if dfsStatus[to] == dfsStart {
			// 第二次进入 from 节点
			dr.hasRing = true
			if dr.debug {
				log.Printf("在节点 %s 找到了环", to)
			}
		}
	}

	if dr.debug {
		log.Printf("节点 %s 所有的邻边都被访问了", from)
	}

	dfsStatus[from] = dfsEnd
}
