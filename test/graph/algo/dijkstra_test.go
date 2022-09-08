package algo

import (
	"go-ads/pkg/graph/algo/dijkstra"
	"go-ads/pkg/graph/sparse"
	"testing"
)

func TestDijkstraL(t *testing.T) {
	g := sparse.CreateSparseGraph(true, false)

	g.Connect("A", "B", 2)
	g.Connect("B", "C", 3)
	g.Connect("A", "C", 6)

	d := dijkstra.CreateDijkstraAlgo(g)
	d.CalculateByL("A")

	d.PrintRoute("C")
}

func TestDijkstraR(t *testing.T) {
	g := sparse.CreateSparseGraph(true, false)

	g.Connect("A", "B", 2)
	g.Connect("B", "C", 3)
	g.Connect("A", "C", 6)

	d := dijkstra.CreateDijkstraAlgo(g)
	d.CalculateByR("A")

	d.PrintRoute("C")
}

func TestDijkstra(t *testing.T) {
	g := sparse.CreateSparseGraph(false, false)

	g.Connect("A", "B", 4)
	g.Connect("A", "D", 2)
	g.Connect("B", "D", 1)
	g.Connect("B", "C", 4)
	g.Connect("C", "D", 1)
	g.Connect("C", "E", 3)
	g.Connect("D", "E", 7)

	d := dijkstra.CreateDijkstraAlgo(g)
	d.CalculateByR("A")

	dests := []string{"A", "B", "C", "D", "E"}

	for _, dest := range dests {
		d.PrintRoute(dest)
	}

}
