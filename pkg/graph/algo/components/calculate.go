package components

import "fmt"

func (tr *components) init() {
	visited := make(map[string]bool)
	vs := tr.g.GetAllVertex()
	for v, _ := range vs {
		visited[v] = false
	}

	uf := make(map[string]int)

	tr.groupNum, tr.visited, tr.uf = 0, visited, uf
}

//
// Calculate
//  @Description: 遍历
//  @receiver tr
//
func (tr *components) Calculate() {
	if tr.g.IsDirect() {
		fmt.Println("有向图暂时不做遍历")
		return
	}
	tr.init()

	// 可以优化
	for v := range tr.g.GetAllVertex() {
		b, _ := tr.visited[v]
		// 没有被访问过的节点才能参与深度遍历
		if !b {
			tr.groupNum = tr.groupNum + 1
			// 一次深度遍历可以找出所有的在一个联通分量里的数据
			tr.deep(v)
		}
	}

	return
}

//
// deep
//  @Description: 深度遍历
//  @receiver tr
//  @param v
//
func (tr *components) deep(v string) {
	tr.visited[v] = true

	tr.uf[v] = tr.groupNum

	adj := tr.g.Adj(v)
	for to, _ := range adj {
		if !tr.visited[to] {
			tr.deep(to)
		}
	}
}
