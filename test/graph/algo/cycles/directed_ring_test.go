package cycles

import (
	"go-dsa/pkg/graph/algo/cycles"
	"go-dsa/pkg/graph/sparse"
	"testing"
)

// Test_Ring1
//
//	@Description: 有向图有环测试
//	@param t
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

// Test_Ring2
//
//	@Description: 有向图无环测试
//	@param t
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
