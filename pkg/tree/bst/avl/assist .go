// Package avl
// @Description: avl 树的辅助数据结构和方法

package avl

import (
	"go-dsa/pkg/tree/bst"
)

// avlNode
// @Description: avl 树的节点
type avlNode struct {
	k      int
	v      string
	left   *avlNode
	right  *avlNode
	height int
}

// ***** ***** ***** ***** ***** ***** ***** ***** ***** *****

// add  *int 主要是更新size
func (node *avlNode) add(k int, v string, size *int, debug bool) *avlNode {
	if node == nil {
		*size++
		// 新增的节点，默认高度为1
		return &avlNode{
			k:      k,
			v:      v,
			height: 1,
		}
	}
	if k == node.k {
		node.v = v
	} else if k < node.k {
		node.left = node.left.add(k, v, size, debug)
	} else {
		node.right = node.right.add(k, v, size, debug)
	}
	// 递归生成完成节点的回溯

	// 添加完节点，计算节点的高度
	node.updateHeight()

	// 完成节点的自平衡
	return node.selfBalance(debug)
}

func (node *avlNode) delete(k int, size *int, debug bool) *avlNode {
	if node == nil {
		return nil
	}
	var rtNode *avlNode
	if k < node.k {
		node.left = node.left.delete(k, size, debug)
		rtNode = node
	} else if k > node.k {
		node.right = node.right.delete(k, size, debug)
		rtNode = node
	} else {
		if node.left == nil && node.right == nil {
			*size--
			return nil
		} else if node.left == nil {
			// 找到要删除的节点
			*size--
			rtNode = node.right
		} else if node.right == nil {
			*size--
			rtNode = node.left
		} else {
			// 找右子树的最小值
			rightMin := node.right.min()

			// 1. 先删除
			// 为什么要先删除， 如果先赋值 node.left ,那么删除就走不到上面三个if的判断，上面的三个的if是真正删除
			rightMin.right = node.right.delete(rightMin.k, size, debug)

			// 2. 后给节点 当前节点的左节点给 rightMin
			rightMin.left = node.left

			// release memory for gc?
			node.left = nil
			rtNode = rightMin
		}
	}
	if rtNode == nil {
		return nil
	}
	rtNode.updateHeight()
	return rtNode.selfBalance(debug)
}

// ***** ***** ***** ***** ***** ***** ***** ***** ***** *****

func (node *avlNode) preOrder(ac bst.Action) {
	if node == nil {
		return
	}
	ac(node.k, node.v)
	node.left.preOrder(ac)
	node.right.preOrder(ac)
}

func (node *avlNode) inOrder(ac bst.Action) {
	if node == nil {
		return
	}
	node.left.inOrder(ac)
	ac(node.k, node.v)
	node.right.inOrder(ac)
}

func (node *avlNode) postOrder(ac bst.Action) {
	if node == nil {
		return
	}
	node.left.postOrder(ac)
	node.right.postOrder(ac)
	ac(node.k, node.v)
}
