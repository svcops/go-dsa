package components

import (
	"fmt"
	"go-ads/pkg/graph"
)

//
//  components
//  @Description: groupNum 分组的编号
//
type components struct {
	g        graph.Graph
	groupNum int
	visited  map[string]bool
	uf       map[string]int
}

//
// CreateComponentsAlgo
//  @Description:j 针对无向图的遍历
//  @param g
//  @return *algo
//
func CreateComponentsAlgo(g graph.Graph) *components {
	return &components{
		g: g,
	}
}

func (tr *components) ComponentSize() int {
	return tr.groupNum
}

//
// HasPath
//  @Description: 两点之间是否有路径
//  @receiver tr
//  @param from
//  @param to
//  @return bool
//
func (tr *components) HasPath(from, to string) bool {
	fromGroupNum, fromOk := tr.uf[from]
	toGroupNum, toOk := tr.uf[to]

	// 原本点就不存在
	if !fromOk || !toOk {
		fmt.Printf("from(%s) or to(%s) does not exist\n", from, to)
		return false
	}
	return fromGroupNum == toGroupNum
}
