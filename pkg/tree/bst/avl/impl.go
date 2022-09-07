// Package avl
// @Description: avl 树的实现

package avl

import (
	"fmt"
	"go-ads/pkg/tree/bst"
)

//
//  avlTree
//  @Description: debug 辅助打印中间过程
//
type avlTree struct {
	size  *int
	root  *avlNode
	debug bool
}

const (
	empty = ""
)

// CreateAvlTree 创建一个avl树
func CreateAvlTree() *avlTree {
	size := 0
	return &avlTree{
		size:  &size,
		root:  nil,
		debug: false,
	}
}

func (tree *avlTree) IsEmpty() bool {
	return *tree.size == 0
}

func (tree *avlTree) Size() int {
	return *tree.size
}

func (tree *avlTree) Get(k int) (string, error) {
	find := tree.root.get(k)
	if find == nil {
		return empty, bst.TreeError(fmt.Sprintf("找不到 k = %d 的值", k))
	} else {
		return find.v, nil
	}
}

// Contains 是否包含元素
func (tree *avlTree) Contains(k int) bool {
	find := tree.root.get(k)
	return find != nil
}

func (tree *avlTree) Max() (int, string, error) {
	max := tree.root.max()
	if max == nil {
		return 0, empty, bst.TreeError("Not Found Max")
	} else {
		return max.k, max.v, nil
	}
}

func (tree *avlTree) Min() (int, string, error) {
	min := tree.root.min()
	if min == nil {
		return 0, empty, bst.TreeError("Not Found Min")
	} else {
		return min.k, min.v, nil
	}
}

func (tree *avlTree) Add(k int, v string) {
	tree.root = tree.root.add(k, v, tree.size, tree.debug)
}

func (tree *avlTree) Delete(k int) (bool, error) {
	if !tree.Contains(k) {
		return false, bst.TreeError(fmt.Sprintf("不存在%d", k))
	}
	tree.root = tree.root.delete(k, tree.size, tree.debug)
	return true, nil
}

func (tree *avlTree) IsBST() bool {
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

func (tree *avlTree) IsBalanced() bool {
	return isBalanced(tree.root)
}

func isBalanced(node *avlNode) bool {
	if node == nil {
		return true
	}
	bf := getBalanceFactor(node)
	if bf >= 2 || bf <= -2 {
		return false
	}
	return isBalanced(node.left) && isBalanced(node.right)
}

func (tree *avlTree) SetDebug(debug bool) {
	tree.debug = debug
}
