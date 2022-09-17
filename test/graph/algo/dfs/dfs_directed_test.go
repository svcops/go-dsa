// Package dfs
// @Description: 有向图测试

package dfs

import (
	"fmt"
	"go-ads/pkg/graph/algo/dfs"
	"go-ads/pkg/graph/sparse"
	"testing"
)

//
// TestFindRoutes1
//  @Description: 两条路径测试
//  @param t
//
func TestDirectedFindRoute(t *testing.T) {
	g := sparse.CreateSparseGraph(true, true)

	g.Connect(a, b, 1)
	g.Connect(b, e, 2)

	g.Connect(a, c, 3)
	g.Connect(c, e, 4)

	g.Connect(a, d, 5)
	g.Connect(d, e, 6)

	algo := dfs.CreateDfsAlgo(g)

	src, dest := a, e

	algo.PrintRoutesByV(src, dest)

	fmt.Println()
	algo.PrintRoutesByE(src, dest)
}
