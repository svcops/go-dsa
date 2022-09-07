package sort

import (
	"go-ads/pkg/sort"
	"testing"
)

func Test_Len(t *testing.T) {
	s := []int{1, 2, 3}
	t.Log(s)
	// sort.Swap(s, 0, 1)
	t.Log(sort.Len(s))
}

func Test_Swap(t *testing.T) {
	s := []int{1, 2, 3}
	t.Log(s)
	sort.Swap(s, 0, 1)
	t.Log(s)
}
