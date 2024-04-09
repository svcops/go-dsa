package ring

import (
	"go-ads/pkg/graph/algo/ring"
	"go-ads/pkg/graph/sparse"
	"testing"
)

// Test_Ring3
//
//	@Description: 无向图有环测试
//	@param t
func Test_Ring3(t *testing.T) {
	g := sparse.CreateSparseGraph(false, false)
	g.Connect("A", "B", 1)
	g.Connect("B", "C", 1)
	g.Connect("C", "D", 1)

	g.Connect("D", "B", 1)
	// A - B - C -  D
	//      \       |
	//        - - -

	// g.Show()

	algo := ring.CreateRingAlgo(g)

	algo.SetDebug(true)

	algo.Calculate()

	if algo.HasRing() {
		t.Logf("Calculate Success (HasRing = %v)", algo.HasRing())
	} else {
		t.Errorf("Calculate Failure (HasRing = %v)", algo.HasRing())
	}

}

// Test_Ring4
//
//	@Description: 无向图无环测试
//	@param t
func Test_Ring4(t *testing.T) {
	g := sparse.CreateSparseGraph(false, false)
	g.Connect("A", "B", 1)
	g.Connect("A", "C", 1)
	g.Connect("A", "D", 1)
	g.Connect("A", "E", 1)
	//     B
	//     |
	// C - A - D
	//     |
	//     E

	algo := ring.CreateRingAlgo(g)

	algo.SetDebug(true)

	algo.Calculate()
	if !algo.HasRing() {
		t.Logf("Calculate Success (HasRing = %v)", algo.HasRing())
	} else {
		t.Errorf("Calculate Failure (HasRing = %v)", algo.HasRing())
	}
}
