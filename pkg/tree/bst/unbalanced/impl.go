package unbalanced

import (
	"fmt"
	"go-dsa/pkg/tree/bst"
)

const empty = ""

type unbalancedBST struct {
	size  *int
	root  *bstNode
	debug bool
}

// NewUnbalancedBST 这个地方可以优化掉
func NewUnbalancedBST() *unbalancedBST {
	size := 0
	return &unbalancedBST{
		size:  &size,
		root:  nil,
		debug: false,
	}
}

// IsEmpty 是否为空
func (tree *unbalancedBST) IsEmpty() bool {
	return *tree.size == 0
}

func (tree *unbalancedBST) Size() int {
	return *tree.size
}

func (tree *unbalancedBST) Get(k int) (string, error) {
	find := tree.root.get(k)
	if find == nil {
		return empty, bst.TreeError(fmt.Sprintf("找不到 k = %d 的值", k))
	} else {
		return find.v, nil
	}
}

func (tree *unbalancedBST) Contains(k int) bool {
	find := tree.root.get(k)
	return find != nil
}

func (tree *unbalancedBST) Max() (int, string, error) {
	max := tree.root.max()
	if max == nil {
		return 0, empty, bst.TreeError("Not Found Max")
	} else {
		return max.k, max.v, nil
	}
}

func (tree *unbalancedBST) Min() (int, string, error) {
	min := tree.root.min()
	if min == nil {
		return 0, empty, bst.TreeError("Not Found Min")
	} else {
		return min.k, min.v, nil
	}
}

// Add 新增节点
func (tree *unbalancedBST) Add(k int, v string) {
	tree.root = tree.root.add(k, v, tree.size)
}

// Delete 删除节点
func (tree *unbalancedBST) Delete(k int) (bool, error) {
	if !tree.Contains(k) {
		return false, bst.TreeError(fmt.Sprintf("不存在%d", k))
	}
	tree.root = tree.root.delete(k, tree.size, tree.debug)
	return true, nil
}

// IsBST 是否是二分搜索树
func (tree *unbalancedBST) IsBST() bool {
	var list []int
	tree.InOrder(func(i int, s string) {
		list = append(list, i)
	})
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}
	return true
}

func (tree *unbalancedBST) IsBalanced() bool {
	return false
}

func (tree *unbalancedBST) SetDebug(debug bool) {
	tree.debug = debug
}
