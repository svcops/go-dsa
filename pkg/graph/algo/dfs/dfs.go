// Package dfs
// @Description: 深度遍历
// @Description: 求所有路径

package dfs

import (
	"go-ads/pkg/graph"
)

type Dfs struct {
	g    graph.Graph
	dest string
}

func CreateDfsAlgo(g graph.Graph) *Dfs {
	return &Dfs{
		g: g,
	}
}
