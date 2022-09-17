package mst

import "go-ads/pkg/graph"

type MST interface {
	Calculate() []graph.Edge
}
