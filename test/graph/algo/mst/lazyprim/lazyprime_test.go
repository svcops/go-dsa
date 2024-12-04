package lazyprim

import (
	"go-dsa/pkg/graph/algo/mst/lazyprim"
	"go-dsa/pkg/graph/sparse"
	"testing"
)

func TestMST(t *testing.T) {
	g := sparse.CreateSparseGraph(false, true)

	g.Connect("A", "B", 1)
	g.Connect("B", "C", 2)
	g.Connect("C", "D", 3)
	g.Connect("D", "A", 4)

	algo := lazyprim.CreateLazyPrimAlgo(g)

	mst := algo.Calculate()

	for _, edge := range mst {
		t.Log(edge.ToString())
	}

}
