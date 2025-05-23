// Package cycles
// @Description: 判断无向图是否有环
package cycles

import "go-dsa/pkg/graph"

type cycles struct {
	g       graph.Graph
	hasRing bool
	debug   bool
}

func CreateCyclesAlgo(g graph.Graph) *cycles {
	return &cycles{
		g:     g,
		debug: false,
	}
}

func (r *cycles) Calculate() {
	r.hasRing = false

	direct := r.g.IsDirect()
	if direct {
		dr := &directedRing{
			cycles: r,
		}
		dr.calculate()
	} else {
		dr := &undirectedRing{
			cycles: r,
		}
		dr.calculate()
	}
}

// HasRing
//
//	@Description: 是否有环
//	@receiver r
//	@return bool
func (r *cycles) HasRing() bool {
	return r.hasRing
}

func (r *cycles) SetDebug(debug bool) {
	r.debug = debug
}
