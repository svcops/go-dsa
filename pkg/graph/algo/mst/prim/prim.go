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
