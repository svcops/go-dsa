// Package dfs
// @Description: 深度遍历所有路径，记录下顶点

package dfs

import (
	"fmt"
	"golang.org/x/exp/maps"
)

func (d *Dfs) PrintRoutesByV(src, dest string) {
	routes := d.FindRoutesByV(src, dest)

	fmt.Printf("从 %s 到 %s 共有 %d 条路径\n", src, dest, len(routes))
	for _, route := range routes {
		fmt.Print(src)
		for _, v := range route {
			fmt.Printf(" -> %s", v)
		}
		fmt.Println()
	}
}

// FindRoutesByV
//
//	@Description: 求两点的所有路径
//	@receiver d
func (d *Dfs) FindRoutesByV(src, dest string) [][]string {

	d.dest = dest

	// 初始化节点是否被访问的map
	visited := make(map[string]bool)
	for v := range d.g.Vertexes() {
		visited[v] = false
	}

	route := vertexRoutes{
		from: src,
		rts:  make([][]string, 0),
	}

	d.dfsRecordV(src, visited, &route)
	return route.rts
}

func (d *Dfs) dfsRecordV(v string, visited map[string]bool, vRoutes *vertexRoutes) {
	visited[v] = true

	for other := range d.g.Adj(v) {
		// fmt.Println("weight is",w)
		if other == d.dest {
			// 遍历 C 的时候找到了 E,形成一条路径
			findRoute := []string{other}

			vRoutes.rts = append(vRoutes.rts, findRoute)
			// 已经到达终点，标记终点不需要访问
			visited[other] = true
		} else {
			hasVisited, _ := visited[other]
			if !hasVisited {
				subVisited := make(map[string]bool)
				maps.Copy(subVisited, visited)

				subRoutes := vertexRoutes{
					from: other,
					rts:  make([][]string, 0),
				}

				d.dfsRecordV(other, subVisited, &subRoutes)

				if len(subRoutes.rts) > 0 {
					// 寻找到了子路径
					for _, contains := range subRoutes.rts {
						mergeRoute := make([]string, 1+len(contains))
						// 新增节点B
						mergeRoute[0] = other
						// 合并之前的路径
						// mergeRoute = append(mergeRoute, contains...)
						for i, mv := range contains {
							mergeRoute[i+1] = mv
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

// vertexRoutes
// @Description: 顶点的路径
type vertexRoutes struct {
	from string
	rts  [][]string
}
