package dijkastra

import (
	"go-dsa/pkg/graph/algo/dijkstra"
	"go-dsa/pkg/graph/sparse"
	"testing"
)

const (
	directed = true
	// undirected = false
	// cover      = true
	uncover = false
)

func TestDijkstraL(t *testing.T) {
	g := sparse.CreateSparseGraph(directed, uncover)

	g.Connect("A", "B", 2)
	g.Connect("B", "C", 3)
	g.Connect("A", "C", 6)

	d := dijkstra.CreateDijkstraAlgo(g)
	d.CalculateByL("A")

	d.PrintShortestRoute("C")
}

func TestDijkstraR(t *testing.T) {
	g := sparse.CreateSparseGraph(directed, uncover)

	g.Connect("A", "B", 2)
	g.Connect("B", "C", 3)
	g.Connect("A", "C", 6)

	d := dijkstra.CreateDijkstraAlgo(g)
	d.CalculateByR("A")

	d.PrintShortestRoute("C")
}

func TestHasShortestPath(t *testing.T) {
	g := sparse.CreateSparseGraph(directed, uncover)

	g.Connect("A", "B", 2)
	g.Connect("B", "C", 3)
	g.Connect("A", "C", 6)
	g.Connect("D", "C", 1)

	d := dijkstra.CreateDijkstraAlgo(g)
	d.CalculateByL("A")

	t.Log("has shortest root from ", d.HasShortestRoute("D"))
}

func TestDijkstra(t *testing.T) {
	g := sparse.CreateSparseGraph(directed, uncover)

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
		d.PrintShortestRoute(dest)
	}

}
