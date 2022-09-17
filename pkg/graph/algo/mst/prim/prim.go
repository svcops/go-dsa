// Package prim
// @Description: 优化的思路是用 IndexMinHeap
// @Description: 因为我实现的图是用 map[string]map[string]float64 这种数据结构，所以用不了IndexMinHeap

package prim

import (
	"go-ads/pkg/graph"
)

type prim struct {
	g graph.Graph
}

func CreatePrimAlgo(g graph.Graph) *prim {
	return &prim{
		g: g,
	}
}

func (p *prim) Calculate() []graph.Edge {
	return nil
}
