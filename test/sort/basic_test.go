package sort

import (
	"go-ads/pkg/sorter"
	"testing"
)

func Test_Len(t *testing.T) {
	s := []int{1, 2, 3}
	t.Log(s)
	// sorter.Swap(s, 0, 1)
	t.Log(sorter.Len(s))
}

func Test_Swap(t *testing.T) {
	s := []int{1, 2, 3}
	t.Log(s)
	sorter.Swap(s, 0, 1)
	t.Log(s)
}
