// Package unbalanced
// @Description: 二分搜索树的辅助数据结构和方法

package unbalanced

import (
	"go-ads/pkg/tree/bst"
	"log"
)

type bstNode struct {
	k     int
	v     string
	left  *bstNode
	right *bstNode
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

// ***** ***** ***** ***** ***** ***** ***** ***** ***** *****

func (node *bstNode) delete(k int, size *int, debug bool) *bstNode {
	if node == nil {
		return nil
	}
	if k < node.k {
		node.left = node.left.delete(k, size, debug)
		return node
	} else if k > node.k {
		node.right = node.right.delete(k, size, debug)
		return node
	} else {
		if node.left == nil && node.right == nil {
			*size--
			return nil
		}
		// 找到要删除的节点
		if node.left == nil {
			*size--
			return node.right
		}
		if node.right == nil {
			*size--
			return node.left
		}
		if debug {
			log.Println("删除节点寻找找右子树的最小值")
		}

		// 找右子树的最小值
		rightMin := node.right.min()

		// 1. 先删除
		// 为什么要先删除， 如果先赋值 node.left ,那么删除就走不到上面三个if的判断，上面的三个的if是真正删除
		rightMin.right = node.right.delete(rightMin.k, size, debug)

		// 2. 后给节点 当前节点的左节点给 rightMin
		rightMin.left = node.left

		// release memory?
		node.left = nil
		return rightMin
	}
}

// ***** ***** ***** ***** ***** ***** ***** ***** ***** *****

func (node *bstNode) preOrder(ac bst.Action) {
	if node == nil {
		return
	}
	ac(node.k, node.v)
	node.left.preOrder(ac)
	node.right.preOrder(ac)
}

func (node *bstNode) inOrder(ac bst.Action) {
	if node == nil {
		return
	}
	node.left.inOrder(ac)
	ac(node.k, node.v)
	node.right.inOrder(ac)
}

func (node *bstNode) postOrder(ac bst.Action) {
	if node == nil {
		return
	}
	node.left.postOrder(ac)
	node.right.postOrder(ac)
	ac(node.k, node.v)
}
