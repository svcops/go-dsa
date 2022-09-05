// Package avl
// @Description: avl 树的辅助数据结构和方法

package avl

//
//  avlNode
//  @Description: avl 树的节点
//
type avlNode struct {
	k      int
	v      string
	left   *avlNode
	right  *avlNode
	height int
}
