package algo

import (
	"go-ads/pkg/graph/algo/ring"
	"go-ads/pkg/graph/sparse"
	"testing"
)

func Test_Ring1(t *testing.T) {
	g := sparse.CreateSparseGraph(true, false)
	g.Connect("A", "B", 1)
	g.Connect("B", "C", 1)
	g.Connect("C", "D", 1)
	// A -> B -> C -> D

	algo := ring.CreateRingAlgo(g)
	algo.Calculate()
	t.Log("第一次计算是否有环", algo.HasRing())

	algo.SetDebug(true)
	g.Connect("D", "B", 1)
	// A -> B -> C ->  D
	//      ↑         ↙
	//        ↖ <-  ↙

	g.Show()
	t.Log("连接 D -> B ")
	algo.Calculate()
	t.Log("第二次计算是否有环", algo.HasRing())
}

func Test_Ring2(t *testing.T) {
	g := sparse.CreateSparseGraph(true, false)
	g.Connect("A", "B", 1)
	g.Connect("B", "C", 1)
	g.Connect("C", "D", 1)

	// 算法的正确性的测试
	g.Connect("E", "C", 1)
	// A -> B -> C -> D
	//         ↗
	//        E
	algo := ring.CreateRingAlgo(g)
	algo.SetDebug(true)
	algo.Calculate()
	t.Log("计算是否有环", algo.HasRing())

}

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

func Test_Ring4(t *testing.T) {
	g := sparse.CreateSparseGraph(false, false)
	g.Connect("A", "B", 1)
	g.Connect("A", "C", 1)
	g.Connect("A", "D", 1)
	g.Connect("A", "E", 1)

	algo := ring.CreateRingAlgo(g)

	algo.SetDebug(true)

	algo.Calculate()
	if !algo.HasRing() {
		t.Logf("Calculate Success (HasRing = %v)", algo.HasRing())
	} else {
		t.Errorf("Calculate Failure (HasRing = %v)", algo.HasRing())
	}
}
