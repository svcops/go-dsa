package avl

import (
	"go-ads/pkg/tree/bst"
	"go-ads/pkg/tree/bst/avl"
	template "go-ads/test/tree/bst"
	"testing"
)

func creator() bst.BST {
	return avl.CreateAvlTree()
}

func Test_LL(t *testing.T) {
	tree := creator()
	tree.SetDebug(true)

	tree.Add(3, "3")
	tree.Add(2, "2")
	tree.Add(1, "1")

	t.Log(tree)
}

func Test_RR(t *testing.T) {
	tree := creator()
	tree.SetDebug(true)

	tree.Add(1, "1")
	tree.Add(2, "2")
	tree.Add(3, "3")

	t.Log(tree)
}

func Test_LR(t *testing.T) {
	tree := creator()
	tree.SetDebug(true)

	tree.Add(3, "3")
	tree.Add(1, "1")
	tree.Add(2, "2")

	t.Log(tree)
}

func Test_RL(t *testing.T) {
	tree := creator()
	tree.SetDebug(true)

	tree.Add(1, "1")
	tree.Add(3, "3")
	tree.Add(2, "2")

	t.Log(tree)
}

func Test_Add(t *testing.T) {
	template.TestAdd(t, creator)
}

func Test_Get(t *testing.T) {
	template.TestGet(t, creator)
}

func Test_IsBST(t *testing.T) {
	template.TestIsBST(t, creator)
}

func Test_Bfs(t *testing.T) {
	template.TestBfs(t, creator)
}

func Test_Dfs(t *testing.T) {
	template.TestDfs(t, creator)
}

func Test_PreOrder(t *testing.T) {
	template.TestPreOrder(t, creator)
}

func Test_InOrder(t *testing.T) {
	template.TestInOrder(t, creator)
}

func Test_PostOrder(t *testing.T) {
	template.TestPostOrder(t, creator)
}

func Test_Delete(t *testing.T) {
	template.TestDelete(t, creator)
}

func Test_RandomDelete(t *testing.T) {
	size := 10000 * 10
	max := size * 100
	template.TestRandomDelete(t, creator, size, max)
}

func Test_IsBalanced(t *testing.T) {
	size := 10000 * 10
	tree := creator()
	template.RandomDataInit(tree, size, size)
	t.Log("IsBalanced =", tree.IsBalanced())
}
