package simple

import (
	"fmt"
	"go-ads/tree/bst"
)

const empty = ""

type unbalancedBST struct {
	size *int
	root *bstNode
}

// NewSimpleBST 这个地方可以优化掉
func NewSimpleBST() *unbalancedBST {
	size := 0
	return &unbalancedBST{
		size: &size,
	}
}

// 是否为空
func (tree *unbalancedBST) IsEmpty() bool {
	return *tree.size == 0
}

func (tree *unbalancedBST) Size() int {
	return *tree.size
}

func (tree *unbalancedBST) Add(k int, v string) {
	tree.root = tree.root.add(k, v, tree.size)
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

// Delete 删除节点
func (tree *unbalancedBST) Delete(k int) (bool, error) {
	if !tree.Contains(k) {
		return false, bst.TreeError(fmt.Sprintf("不存在%d", k))
	}
	tree.root = tree.root.delete(k, tree.size)
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

// Dfs 深度遍历
func (tree *unbalancedBST) Dfs(ac bst.Action) {
	// 默认中序遍历
	tree.root.inOrder(ac)
}

// Bfs 广度遍历
func (tree *unbalancedBST) Bfs(ac bst.Action) {
	if tree.IsEmpty() {
		return
	}
	queue := newQueue()
	queue.add(tree.root)
	for !queue.isEmpty() {
		node := queue.poll()
		ac(node.k, node.v)
		if node.left != nil {
			queue.add(node.left)
		}
		if node.right != nil {
			queue.add(node.right)
		}
	}
}

// PreOrder 先序遍历
func (tree *unbalancedBST) PreOrder(ac bst.Action) {
	tree.root.preOrder(ac)
}

// InOrder 中序遍历
func (tree *unbalancedBST) InOrder(ac bst.Action) {
	tree.root.inOrder(ac)
}

// PostOrder 后序遍历
func (tree *unbalancedBST) PostOrder(ac bst.Action) {
	tree.root.postOrder(ac)
}

func (tree *unbalancedBST) IsBalanced() bool {
	return false
}
