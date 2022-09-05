package bst

import (
	"go-ads/tree/bst"
	"go-ads/tree/bst/simple"
	"testing"
)

func Test_ADD(t *testing.T) {
	testAdd(t, prepareUnbalanceBST(true))
}

func Test_GET(t *testing.T) {
	testGet(t, prepareUnbalanceBST(true))
}

func Test_IsBST(t *testing.T) {
	testIsBST(t, prepareUnbalanceBST(false))
}

func Test_Bfs(t *testing.T) {
	testBfs(t, prepareUnbalanceBST(false))
}

func Test_Dfs(t *testing.T) {
	testDfs(t, prepareUnbalanceBST(false))
}

func Test_PreOrder(t *testing.T) {
	testPreOrder(t, prepareUnbalanceBST(false))
}

func Test_InOrder(t *testing.T) {
	testInOrder(t, prepareUnbalanceBST(false))
}

func Test_PostOrder(t *testing.T) {
	testPostOrder(t, prepareUnbalanceBST(false))
}

func Test_Delete(t *testing.T) {
	testDelete(t, prepareUnbalanceBST(true))
}

func Test_Delete2(t *testing.T) {
	testDelete2(t, prepareUnbalanceBST(true))
}

func prepareUnbalanceBST(empty bool) prepareFunc {
	if empty {
		// 准备一个空的不平衡的二分搜索树
		return func() bst.BST {
			return simple.NewSimpleBST()
		}
	} else {
		// 准备一个包含数据的不平衡的二分搜索树
		return func() bst.BST {
			tree := simple.NewSimpleBST()
			tree.Add(2, "BBB")
			tree.Add(1, "AAA")
			tree.Add(3, "CCC")
			return tree
		}
	}
}
