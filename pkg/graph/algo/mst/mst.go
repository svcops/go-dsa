package mst

import "go-dsa/pkg/graph"

type MST interface {
	Calculate() []graph.Edge
}
