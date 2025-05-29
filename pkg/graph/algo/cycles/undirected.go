// Package cycles
// @Description: 无向图判断环
package cycles

import (
	"log"
)

type undirectedRing struct {
	*cycles
	visited map[string]bool
}

func (ur *undirectedRing) calculate() {
	ur.visited = make(map[string]bool)
	for v := range ur.g.Vertexes() {
		ur.visited[v] = false
	}
	for v := range ur.g.Vertexes() {
		if ur.hasRing {
			// 如果找到了环，就不需要再寻找了
			break
		}
		if !ur.visited[v] {
			ur.dfsDirected(v, v)
		}
	}
}

// dfsDirected
//
//	@Description:
//	@receiver urk
//	@param v
//	@param from 深度遍历的时候，记录从哪个节点(or边)过来的
//	@return bool
func (ur *undirectedRing) dfsDirected(v string, from string) {
	ur.visited[v] = true
	if ur.debug {
		log.Printf("开始遍历节点 %s ; 记录上一层节点 %v", v, from)
	}
	for to := range ur.g.Adj(v) {
		if ur.hasRing {
			// 回溯的时候，如果有环了，就不需要再判断了
			break
		}
		if v == from {
			if ur.debug {
				log.Printf("节点 %s 第一次进入 dfs", v)
			}
			// ur.dfsDirected(to, &from{v})
			ur.dfsDirected(to, v)
		} else {
			if from == to {
				// 排除类似  A -> B
				//              B -> A
				continue
			} else {
				if ur.visited[to] {
					ur.hasRing = true
					if ur.debug {
						log.Printf("在节点 %s 找到了环", to)
					}
				} else {
					// 这里一定要返回，递归的逐级返回问题
					ur.dfsDirected(to, v)
				}
			}
		}

	}

	if ur.debug {
		log.Printf("节点 %s 遍历完，回溯上一层节点 %s ", v, from)
	}
}
