package cycles

import (
	"go-dsa/pkg/graph/algo/cycles"
	"go-dsa/pkg/graph/sparse"
	"testing"
)

func Test_Directed_Cycle_Analyzer(t *testing.T) {

	/*
			  ↗ B ↘
			A   ↓   C
		      ↖ D ↙
	*/
	g := sparse.CreateSparseGraph(true, false)
	g.Connect("A", "B", 1)
	g.Connect("B", "C", 1)
	g.Connect("C", "D", 1)
	g.Connect("D", "A", 1)
	g.Connect("B", "D", 1)

	result := cycles.CreateCycleAnalyzer(g, true).FindCycles()

	result.PrintCycles()
}

func Test_InDirected_Cycle_Analyzer(t *testing.T) {

	/*
			  ↗ B ↘
			A   ↓   C
		      ↖ D ↙
	*/
	g := sparse.CreateSparseGraph(false, false)
	g.Connect("A", "B", 1)
	g.Connect("B", "C", 1)
	g.Connect("C", "D", 1)
	g.Connect("D", "A", 1)
	g.Connect("B", "D", 1)

	result := cycles.CreateCycleAnalyzer(g, true).FindCycles()

	result.PrintCycles()
}
