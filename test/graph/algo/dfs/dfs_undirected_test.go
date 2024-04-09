// Package dfs
// @Description: 无向图测试

package dfs

import (
	"fmt"
	"go-ads/pkg/graph/algo/dfs"
	"go-ads/pkg/graph/sparse"
	"testing"
)

const (
	a, b, c, d, e, f = "A", "B", "C", "D", "E", "F"
)

// TestFindRoutes1
//
//	@Description: 两条路径测试
//	@param t
func TestUndirectedFindRoutes1(t *testing.T) {
	g := sparse.CreateSparseGraph(false, true)

	g.Connect(a, b, 1)
	g.Connect(b, c, 1)
	g.Connect(b, d, 1)
	g.Connect(c, e, 1)
	g.Connect(d, e, 1)

	g.Connect(c, f, 1)

	//            F
	//          /
	//         C
	//       /   \
	// A - B      E
	//       \   /
	//         D

	algo := dfs.CreateDfsAlgo(g)

	src, dest := a, e

	algo.PrintRoutesByV(src, dest)

	fmt.Println()
	algo.PrintRoutesByE(src, dest)
}

// TestFindRoutes2
//
//	@Description: 多条路径测试
//	@param t
func TestUndirectedFindRoutes2(t *testing.T) {
	g := sparse.CreateSparseGraph(false, true)
	// g := dense.CreateDenseGraph(10, false, true)

	g.Connect(a, b, 1.0)
	g.Connect(a, c, 1.0)
	g.Connect(a, d, 1.0)

	g.Connect(b, e, 1.0)
	g.Connect(c, e, 1.0)
	g.Connect(d, e, 1.0)

	g.Connect(b, c, 1.0)
	g.Connect(c, d, 1.0)

	//      B
	//   /  |  \
	// A - C - E
	//  \  |  /
	//     D

	algo := dfs.CreateDfsAlgo(g)

	src, dest := a, e

	algo.PrintRoutesByV(src, dest)
	fmt.Println()
	algo.PrintRoutesByE(src, dest)
}
