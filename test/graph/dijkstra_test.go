package graph

import (
	"fmt"
	"go-ads/pkg/graph/algo"
	"go-ads/pkg/graph/sparse"
	"testing"
)

func TestDijkstra(t *testing.T) {
	g := sparse.CreateSparseGraph(true, false)

	g.Connect("A", "B", 2)
	g.Connect("B", "C", 3)
	g.Connect("A", "C", 6)

	di := algo.CreateDijkstraAlgo(g, "A")

	di.Calculate()

	fmt.Println()

}
