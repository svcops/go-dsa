package graph

import (
	"go-ads/pkg/graph"
	"go-ads/pkg/graph/sparse"
	"testing"
)

type e struct {
	from, to string
	weight   float64
}

func Test_Show(t *testing.T) {

	tests := []struct {
		direct bool
		cover  bool
		data   []e
	}{
		{
			// 有向图测试数据
			direct: true,
			cover:  false,
			data: []e{
				{"A", "B", 3},
				{"B", "C", 4},
				{"C", "A", 5},
			},
		},
		{
			// 无向图测试数据
			direct: false,
			cover:  false,
			data: []e{
				{"A", "B", 3},
				{"B", "C", 4},
				{"C", "A", 5},
			},
		},
		{
			// 有向图覆盖测试
			direct: true,
			cover:  true,
			data: []e{
				{"A", "B", 3},
				{"B", "C", 4},
				{"C", "A", 5},
				{"C", "A", 9},
			},
		},
		{
			// 无向图向图覆盖测试
			direct: false,
			cover:  true,
			data: []e{
				{"A", "B", 3},
				{"B", "C", 4},
				{"C", "A", 5},
				{"C", "A", 9},
			},
		},
	}
	var g graph.Graph

	for _, test := range tests {
		// t.Logf("show graph; direct = %v ; cover = %v", test.direct, test.cover)
		g = sparse.CreateSparseGraph(test.direct, test.cover)

		for _, d := range test.data {
			g.Connect(d.from, d.to, d.weight)
		}
		g.Show()
	}
}

func TestGetAllVertex(t *testing.T) {
	g := sparse.CreateSparseGraph(true, false)
	g.Connect("a", "b", 1)

	if g.Vertices().Size() != 2 {
		t.Error("g.Vertices() error !!!")
	}
}
