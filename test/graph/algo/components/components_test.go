package components

import (
	"go-dsa/pkg/graph/algo/components"
	"go-dsa/pkg/graph/sparse"
	"testing"
)

func TestTraverse(t *testing.T) {

	g := sparse.CreateSparseGraph(false, true)

	g.Connect("A", "B", 1)
	g.Connect("C", "D", 1)

	g.Show()

	compAlgo := components.CreateComponentsAlgo(g)

	compAlgo.Calculate()

	t.Log("component size", compAlgo.ComponentSize())

	t.Log("has path ", compAlgo.HasPath("A", "B"))
	t.Log("has path ", compAlgo.HasPath("A", "C"))

	t.Log("has path ", compAlgo.HasPath("AA", "BB"))
	t.Log("has path ", compAlgo.HasPath("AA", "CC"))
}
