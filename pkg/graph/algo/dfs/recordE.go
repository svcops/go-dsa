// Package dfs
// @Description: 深度遍历所有的路径，记录下边
package dfs

import (
	"fmt"
	"go-ads/pkg/graph"
	"golang.org/x/exp/maps"
)

func (d *Dfs) PrintRoutesByE(src, dest string) {
	routes := d.FindRoutesByE(src, dest)

	fmt.Printf("从 %s 到 %s 共有 %d 条路径\n", src, dest, len(routes))

	for _, route := range routes {
		routeWeight := 0.0

		fmt.Print("路径为: ", src)
		for _, edge := range route {
			routeWeight += edge.Weight
			fmt.Printf(" -(%g)-> %s", edge.Weight, edge.To)
		}
		fmt.Println()
		fmt.Println("路径的权重为:", routeWeight)
		fmt.Println()
	}

}

// FindRoutesByE
//
//	@Description: 返回所有路径
//	@receiver d
//	@param src
//	@param dest
//	@return [][]*Edge
func (d *Dfs) FindRoutesByE(src, dest string) [][]*graph.Edge {
	d.dest = dest

	visited := make(map[string]bool)
	for v := range d.g.Vertices() {
		visited[v] = false
	}

	srcRoutes := edgeRoutes{
		from: src,
		rts:  make([][]*graph.Edge, 0),
	}

	d.dfsRecordE(src, visited, &srcRoutes)
	return srcRoutes.rts
}

// dfsRecordE
//
//	@Description: dfs 遍历寻找所有路径
//	@receiver d
//	@param v
//	@param visited
//	@param vRoutes 节点v的所有子路径集合
func (d *Dfs) dfsRecordE(v string, visited map[string]bool, vRoutes *edgeRoutes) {
	visited[v] = true
	for other, weight := range d.g.Adj(v) {
		if other == d.dest {
			// 找到了最小的子路径 e.g. C-E
			findRoute := []*graph.Edge{{From: v, To: other, Weight: weight}}
			vRoutes.rts = append(vRoutes.rts, findRoute)

			visited[other] = true
		} else {
			if hasVisited, _ := visited[other]; !hasVisited {
				subVisited := make(map[string]bool)
				maps.Copy(subVisited, visited)

				subRoutes := edgeRoutes{
					from: other,
					rts:  make([][]*graph.Edge, 0),
				}

				d.dfsRecordE(other, subVisited, &subRoutes)

				// 递归后的回溯
				if len(subRoutes.rts) > 0 {
					// 有子路径，合并每一条子路径
					for _, subRoute := range subRoutes.rts {
						// 路径合并
						mergeRoute := make([]*graph.Edge, 1+len(subRoute))
						mergeRoute[0] = &graph.Edge{
							From:   v,
							To:     other,
							Weight: weight,
						}

						for i, edge := range subRoute {
							mergeRoute[i+1] = edge
						}

						vRoutes.rts = append(vRoutes.rts, mergeRoute)
					}
				}

				// ??? release memory
				subVisited = nil
				subRoutes.rts = nil
			}
		}

	}
}

type edgeRoutes struct {
	from string
	rts  [][]*graph.Edge
}
