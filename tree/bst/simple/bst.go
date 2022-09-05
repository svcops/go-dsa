package simple

import (
	"errors"
	"fmt"
	"go-ads/tree/bst"
)

const empty = ""

type BST struct {
	size *int
	root *bstNode
}

// 是否为空
func (bst *BST) IsEmpty() bool {
	return *bst.size == 0
}

func (bst *BST) Size() int {
	return *bst.size
}

func (bst *BST) Add(k int, v string) {
	bst.root.add(k, v, bst.size)
}

// add  *int 主要是更新size
func (node *bstNode) add(k int, v string, size *int) *bstNode {
	if node == nil {
		*size++
		return &bstNode{
			k: k,
			v: v,
		}
	}
	if k == node.k {
		node.v = v
	} else if k < node.k {
		node.left = node.left.add(k, v, size)
	} else {
		node.right = node.right.add(k, v, size)
	}
	return node
}

func (bst *BST) Get(k int) (string, error) {
	return bst.root.get(k)
}

func (node *bstNode) get(k int) (string, error) {
	if node == nil {
		return empty, errors.New(fmt.Sprintf("找不到 k = %d 的值", k))
	}

	if k == node.k {
		return node.v, nil
	} else if k < node.k {
		return node.left.get(k)
	} else {
		return node.right.get(k)
	}
}

func (bst *BST) Delete(k int) {

}

func (bst *BST) IsBST() bool {
	return true
}

func (bst *BST) Dfs(ac bst.Action) {

}

func (bst *BST) Bfs(ac bst.Action) {

}

func (bst *BST) PreOrder(ac bst.Action) {

}

func (bst *BST) InOrder(ac bst.Action) {

}

func (bst *BST) PostOrder(ac bst.Action) {

}
